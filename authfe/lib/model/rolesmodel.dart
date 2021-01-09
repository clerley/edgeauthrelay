/*
MIT License

Copyright (c) 2020 Clerley Silveira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

import 'dart:convert';
import 'dart:developer';
import 'dart:io';

import 'package:authfe/model/permissionmodel.dart';
import 'package:authfe/model/settingsmodel.dart';
import 'package:authfe/model/usermodel.dart';
import 'package:flutter/cupertino.dart';
import 'package:http/http.dart';
import 'package:http/io_client.dart';

class RolesProvider extends ChangeNotifier {
  static final RolesProvider _theInstance = RolesProvider._internal();
  ListRolesResponse _cachedRoles;

  RolesProvider._internal();

  factory RolesProvider() {
    return _theInstance;
  }

  Future<RoleResponse> insertRole(Role role) async {
    return roleUpdate(role, "insert");
  }

  Future<RoleResponse> saveRole(Role role) async {
    return roleUpdate(role, "update");
  }

  Future<RoleResponse> roleUpdate(Role role, String updateType) async {
    RoleResponse resp = RoleResponse();
    resp.status = "Failure";
    GlobalSettings settings = GlobalSettings();
    try {
      UserProvider users = UserProvider();
      if (!users.login.isLoggedIn()) {
        return resp;
      }
      var jn = role.toJson();
      var jsonObject = json.encode(jn);
      var httpHeader = {
        "Content-type": "application/json",
        "Authorization": "bearer ${users.login.sessionToken}"
      };

      final ioc = new HttpClient();
      ioc.badCertificateCallback =
          (X509Certificate cert, String host, int port) => true;
      final http = new IOClient(ioc);

      var rawResp;
      if (updateType == "insert") {
        var fullUrl = settings.url + "/jwt/role";
        rawResp =
            await http.put(fullUrl, headers: httpHeader, body: jsonObject);
      } else {
        var fullUrl = settings.url + "/jwt/role/${role.id}";
        rawResp =
            await http.post(fullUrl, headers: httpHeader, body: jsonObject);
      }
      resp = RoleResponse.fromJson(json.decode(rawResp.body));
      role.id = resp.role.id;
    } catch (e, stacktrace) {
      print(stacktrace.toString());
      print(e.toString());
    }

    return resp;
  }

  Future<ListRolesResponse> listRoles(int startAt, int endAt) async {
    ListRolesResponse resp = ListRolesResponse();
    resp.status = "Failure";
    GlobalSettings settings = GlobalSettings();
    try {
      UserProvider users = UserProvider();
      if (!users.login.isLoggedIn()) {
        return resp;
      }
      var httpHeader = {"Authorization": "bearer ${users.login.sessionToken}"};

      final ioc = new HttpClient();
      ioc.badCertificateCallback =
          (X509Certificate cert, String host, int port) => true;
      final http = new IOClient(ioc);

      Response rawResp;
      var fullUrl = settings.url + "/jwt/role/$startAt/$endAt";
      rawResp = await http.get(fullUrl, headers: httpHeader);
      if (rawResp.statusCode == 200) {
        resp = ListRolesResponse.fromJson(json.decode(rawResp.body));
      }
      this._cachedRoles = resp;

      notifyListeners();
    } catch (e, stacktrace) {
      print(stacktrace.toString());
      print("Error: ${e.toString()}");
    }

    return resp;
  }

  Future<Role> findRoleByID(String id) async {
    Role role;

    if (this._cachedRoles == null) {
      return role;
    }

    for (var i = 0; i < this._cachedRoles.roles.length; i++) {
      Role role = this._cachedRoles.roles[i];
      if (role.id == id) {
        return role;
      }
    }

    return role;
  }

  bool isCached() {
    if (this._cachedRoles != null &&
        this._cachedRoles.roles != null &&
        this._cachedRoles.roles.length > 0) {
      return true;
    }
    return false;
  }

  List<Role> getCached() {
    if (this.isCached()) {
      return this._cachedRoles.roles;
    }
    return [];
  }

  doNotification() {
    notifyListeners();
  }

  List<Role> filterByDescription(String description) {
    List<Role> filteredList = [];

    if (!isCached()) {
      return filteredList;
    }

    for (var i = 0; i < this._cachedRoles.roles.length; i++) {
      description = description.toLowerCase();
      Role role = this._cachedRoles.roles[i];
      if (role.description.toLowerCase().indexOf(description) >= 0) {
        filteredList.add(role);
      }
    }

    return filteredList;
  }
}

class RoleResponse {
  String status;
  Role role;

  RoleResponse();

  toJson() {
    Map<String, dynamic> jsonMap = {};
    jsonMap['status'] = status;
    jsonMap['role'] = role.toJson();
    return jsonMap;
  }

  RoleResponse.fromJson(Map<String, dynamic> jsonMap) {
    this.status = jsonMap['status'];
    this.role = Role.fromJson(jsonMap['role']);
  }
}

class ListRolesResponse {
  String status;
  List<Role> roles = [];

  ListRolesResponse();

  ListRolesResponse.fromJson(Map<String, dynamic> jsonMap) {
    this.status = jsonMap['status'];
    List<dynamic> rolesList = jsonMap['roles'];
    rolesList.forEach((element) {
      var role = Role.fromJson(element);
      this.roles.add(role);
    });
  }
}

class Role {
  String id;
  String description;
  List<Permission> permissions = [];

  Role() {
    this.id = "-1";
  }

  Role.fromJson(Map<String, dynamic> jsonMap) {
    this.id = jsonMap['id'];
    this.description = jsonMap['description'];
    var jsonEncodedList = jsonMap['permissions'];
    jsonEncodedList.forEach((element) {
      var perm = Permission.fromJson(element);
      this.permissions.add(perm);
    });
  }

  toJson() {
    Map<String, dynamic> jsonMap = Map<String, dynamic>();
    jsonMap['id'] = id;
    jsonMap['description'] = description;
    jsonMap['permissions'] = [];
    for (var i = 0; i < this.permissions.length; i++) {
      var perm = this.permissions[i];
      jsonMap['permissions'].add(perm.toJson());
    }
    return jsonMap;
  }

  bool hasPermission(Permission perm) {
    for (var i = 0; i < permissions.length; i++) {
      var tempPerm = permissions[i];
      if (tempPerm.id == perm.id) {
        return true;
      }
    }

    return false;
  }

  void addPermission(Permission perm) {
    if (!this.hasPermission(perm)) {
      this.permissions.add(perm);
    }
  }

  void removePermission(Permission perm) {
    if (!this.hasPermission(perm)) {
      print(
          "The role does not have the permission: ${perm.description}. It cannot be removed!");
      return;
    }

    int idx = -1;
    for (var i = 0; i < permissions.length; i++) {
      var permission = permissions[i];

      if (permission.id == perm.id) {
        idx = i;
        break;
      }
    }

    if (idx >= 0) {
      this.permissions.removeAt(idx);
      log("Removed the permission with ID:$idx");
    }
  }

  bool isInsertable() {
    return this.id == "-1";
  }
}

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

import 'package:authfe/model/permissionmodel.dart';
import 'package:authfe/model/settings.dart';
import 'package:authfe/model/user.dart';
import 'package:flutter/cupertino.dart';
import 'package:http/http.dart' as http;

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
    GlobalSettings settings = GlobalSettings();
    try {
      UserProvider users = UserProvider();
      if (!users.login.isLoggedIn()) {
        return resp;
      }
      var httpHeader = {
        "Content-type": "application/json",
        "Authentication": "bearer ${users.login.sessionToken}"
      };
      var rawResp;
      var fullUrl = settings.url + "/jwt/role/$startAt/$endAt";
      rawResp = await http.put(fullUrl, headers: httpHeader);
      resp = ListRolesResponse.fromJson(json.decode(rawResp.body));
      this._cachedRoles = resp;
    } catch (e, stacktrace) {
      print(stacktrace.toString());
      print(e.toString());
    }

    return resp;
  }

  Future<Role> findRoleByID(String id) async {
    Role role;

    if (this._cachedRoles == null) {
      return role;
    }

    this._cachedRoles.roles.forEach((element) {
      if (element.id == id) {
        return element;
      }
    });

    return role;
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
    List<Map<String, dynamic>> rolesList = jsonMap['roles'];
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
    var jsonMap = {};
    jsonMap['id'] = id;
    jsonMap['description'] = description;
    jsonMap['permissions'] = [];
    this.permissions.forEach((element) {
      jsonMap['permissions'].add(element.toJson());
    });
  }

  bool hasPermission(Permission perm) {

    for(var i = 0;i< permissions.length; i++ ) {
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
      return;
    }

    Permission tempPerm;
    for (var i = 0; i < permissions.length; i++) {
      var permission = permissions[i];

      if (permission.id == perm.id) {
        tempPerm = perm;
        break;
      }
    }

    if (tempPerm != null) {
      this.permissions.remove(tempPerm);
      log("Removed the permission with ID:${tempPerm.id}");
    }
  }

  bool isInsertable() {
    return this.id == "-1";
  }
}

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

import 'package:authfe/model/settingsmodel.dart';
import 'package:authfe/model/usermodel.dart';
import 'package:flutter/cupertino.dart';
import 'package:http/io_client.dart';

class PermissionProvider extends ChangeNotifier {
  static final PermissionProvider _theInstance =
      PermissionProvider._privateInstance();

  PermissionProvider._privateInstance();

  factory PermissionProvider() {
    return _theInstance;
  }

  ListPermissionResponse _cachedPermissions;

  Permission findPermission(String description, String permission) {
    Permission perm;

    if (_cachedPermissions == null) {
      log('There is currently no cached permission');
      return perm;
    }

    for (int i = 0; i < _cachedPermissions.permissions.length; i++) {
      if (_cachedPermissions.permissions[i].description == description &&
          _cachedPermissions.permissions[i].permission == permission) {
        perm = _cachedPermissions.permissions[i];
        break;
      }
    }

    return perm;
  }

  Permission findPermissionById(String id) {
    Permission perm;

    if (_cachedPermissions == null) {
      log('There is currently no cached permission');
      return perm;
    }

    for (int i = 0; i < _cachedPermissions.permissions.length; i++) {
      if (_cachedPermissions.permissions[i].id == id) {
        perm = _cachedPermissions.permissions[i];
        break;
      }
    }

    return perm;
  }

  Future<ListPermissionResponse> listPermissions(int startAt, int endAt) async {
    GlobalSettings settings = GlobalSettings();
    UserProvider userProvider = UserProvider();
    var listResp = ListPermissionResponse();
    String fullURL = settings.url + "/jwt/permission/$startAt/$endAt";
    final ioc = new HttpClient();
    ioc.badCertificateCallback =
        (X509Certificate cert, String host, int port) => true;
    final http = new IOClient(ioc);

    var response = await http.get(fullURL, headers: {
      "Authorization": "bearer ${userProvider.login.sessionToken}"
    });
    if (response.statusCode == 200) {
      var jsonResp = json.decode(response.body);
      listResp = ListPermissionResponse.fromJson(jsonResp);
    } else {
      listResp.status = "Failure";
    }
    notifyListeners();
    _cachedPermissions = listResp;
    return listResp;
  }

  Future<InsertPermissionResponse> insertPermission(Permission perm) async {
    InsertPermissionResponse resp = InsertPermissionResponse();
    GlobalSettings settings = GlobalSettings();
    UserProvider userProvider = UserProvider();
    String fullURL = settings.url + "/jwt/permission";
    try {
      var jsonObj = json.encode(perm.toJson());

      final ioc = new HttpClient();
      ioc.badCertificateCallback =
          (X509Certificate cert, String host, int port) => true;
      final http = new IOClient(ioc);

      var response = await http.put(fullURL,
          headers: {
            "Content-type": "application/json",
            "Authorization": "bearer ${userProvider.login.sessionToken}",
          },
          body: jsonObj);
      if (response.statusCode == 200) {
        resp = InsertPermissionResponse.fromJson(json.decode(response.body));
        perm.id = resp.id;
      }
    } catch (e, stacktrace) {
      print(stacktrace.toString());
      print(e.toString());
    }
    notifyListeners();
    return resp;
  }

  Future<InsertPermissionResponse> updatePermission(Permission perm) async {
    InsertPermissionResponse resp;
    GlobalSettings settings = GlobalSettings();
    UserProvider userProvider = UserProvider();
    String fullURL = settings.url + "/jwt/permission/${perm.id}";
    try {
      var jsonObj = json.encode(perm.toJson());

      final ioc = new HttpClient();
      ioc.badCertificateCallback =
          (X509Certificate cert, String host, int port) => true;
      final http = new IOClient(ioc);

      var response = await http.post(fullURL,
          headers: {
            "Content-type": "application/json",
            "Authorization": "bearer ${userProvider.login.sessionToken}"
          },
          body: jsonObj);
      if (response.statusCode == 200) {
        resp = InsertPermissionResponse.fromJson(json.decode(response.body));
        perm.id = resp.id;
      }
    } catch (e, stacktrace) {
      print(stacktrace.toString());
      print(e.toString());
    }
    notifyListeners();

    return resp;
  }

  bool isCached() {
    if (this._cachedPermissions == null) {
      return false;
    }

    if (this._cachedPermissions.permissions == null) {
      return false;
    }

    return (this._cachedPermissions.permissions.length > 0);
  }

  List<Permission> getCachedPermissions() {
    List<Permission> permissions = [];

    if (!isCached()) {
      return permissions;
    }

    this._cachedPermissions.permissions.forEach((element) {
      permissions.add(element);
    });

    return permissions;
  }

  doNotification() {
    notifyListeners();
  }
}

class InsertPermissionResponse {
  String status;
  String id;

  InsertPermissionResponse.fromJson(Map<String, dynamic> jsonObj)
      : this.status = jsonObj['status'],
        this.id = jsonObj['id'];

  InsertPermissionResponse() {
    this.status = "Failure";
  }
}

class ListPermissionResponse {
  String status;
  List<Permission> permissions;
  int startAt;
  int endAt;

  ListPermissionResponse.fromJson(Map<String, dynamic> jsonObj) {
    this.status = jsonObj['status'];
    List<dynamic> jsonEncodedList = jsonObj['permissions'];
    if (jsonEncodedList == null) {
      log('There is currently no permission defined in the response');
      return;
    }
    this.permissions = [];
    jsonEncodedList.forEach((element) {
      var perm = Permission.fromJson(element);
      permissions.add(perm);
    });
  }

  ListPermissionResponse() {
    this.status = "Failure";
  }
}

class Permission {
  String id;
  String permission;
  String description;

  Permission();

  Permission.fromJson(Map<String, dynamic> jsonObj)
      : this.id = jsonObj['id'],
        this.permission = jsonObj['permission'],
        this.description = jsonObj['description'];

  Map<String, dynamic> toJson() {
    Map<String, dynamic> jsonObj = Map<String, dynamic>();
    jsonObj['id'] = id;
    jsonObj['permission'] = permission;
    jsonObj['description'] = description;
    return jsonObj;
  }
}

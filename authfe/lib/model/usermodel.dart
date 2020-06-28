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
import 'dart:async';
import 'dart:io';
import 'package:authfe/model/permissionmodel.dart';
import 'package:authfe/model/settingsmodel.dart';
import 'package:http/http.dart' as http;

import 'package:flutter/material.dart';
import 'package:http/http.dart';

class UserProvider extends ChangeNotifier {
  static final UserProvider _theInstance = UserProvider._privateConstructor();
  UsersList _cachedUserList = UsersList();

  UserProvider._privateConstructor();

  factory UserProvider() {
    return _theInstance;
  }

  var login = Login("Failure", "---");

  Future<Login> requestLogin(
      String uniqueID, String username, String password) async {
    try {
      GlobalSettings globalSettings = GlobalSettings();
      var loginRequest = LoginRequest(uniqueID, username, password);
      var fullURL = globalSettings.url + "/jwt/company/login";
      String bodyStr = json.encode(loginRequest.toJson());
      var response = await http.post(fullURL,
          headers: {'Content-type': 'application/json'}, body: bodyStr);
      if (response.statusCode == 200) {
        login = Login.fromJson(json.decode(response.body));
        if (login.status == "Success") {
          login.user.loggedIn = true;
          login.loginRequest = loginRequest;
        } else {
          print("The user was not successfully logged in");
          login.user.loggedIn = false;
        }
      } else {
        print('The response was a failure. Could not connect to the server!');
        login = Login("Failure", "---");
        login.user.loggedIn = false;
      }

      notifyListeners();
    } on SocketException {
      print('The socket threw a SocketException');
    } on ClientException {
      print('The socket threw a ClientException');
    } catch (e, stackTrace) {
      print(
          "An error occurred while processing the login request ${stackTrace.toString()}");
      print(e.toString());
    }
    return login;
  }

  Future<UsersList> listUsers(int startAt, int endAt) async {
    UsersList resp = UsersList();
    resp.status = "Failure";
    GlobalSettings settings = GlobalSettings();
    try {
      if (!login.isLoggedIn()) {
        return resp;
      }
      var httpHeader = {"Authorization": "bearer ${login.sessionToken}"};
      Response rawResp;
      var fullUrl = settings.url + "/jwt/users/$startAt/$endAt";
      rawResp = await http.get(fullUrl, headers: httpHeader);
      if (rawResp.statusCode == 200) {
        resp = UsersList.fromJson(json.decode(rawResp.body));
      }
      this._cachedUserList = resp;

      notifyListeners();
    } catch (e, stacktrace) {
      print(stacktrace.toString());
      print("Error: ${e.toString()}");
    }

    return resp;
  }

  /// insertUser - Inserting the user */
  Future<UserUpdateResponse> insertUser(User user) async {
    UserUpdateResponse resp = UserUpdateResponse();
    resp.status = "Failure";

    try {
      GlobalSettings settings = GlobalSettings();
      if (!login.isLoggedIn()) {
        return resp;
      }
      var httpHeader = {"Authorization": "bearer ${login.sessionToken}"};
      Response rawResp;
      var fullUrl = settings.url + "/jwt/user";
      var encoded = json.encode(user.toJson());
      rawResp = await http.put(fullUrl, headers: httpHeader, body: encoded);
      if (rawResp.statusCode == 200) {
        resp = UserUpdateResponse.fromJson(json.decode(rawResp.body));
      }
      if (resp.status == "Success") {
        user.id = resp.user.id;
      }
      notifyListeners();
    } catch (e, stackTrace) {
      print(stackTrace);
      print(e);
    }

    return resp;
  }

  /// updateUser - Inserting the user */
  Future<UserUpdateResponse> updateUser(User user) async {
    UserUpdateResponse resp = UserUpdateResponse();
    resp.status = "Failure";

    try {
      GlobalSettings settings = GlobalSettings();
      if (!login.isLoggedIn()) {
        return resp;
      }
      var httpHeader = {"Authorization": "bearer ${login.sessionToken}"};
      Response rawResp;
      var fullUrl = settings.url + "/jwt/user/${user.username}";
      var encoded = json.encode(user.toJson());
      rawResp = await http.post(fullUrl, headers: httpHeader, body: encoded);
      if (rawResp.statusCode == 200) {
        resp = UserUpdateResponse.fromJson(json.decode(rawResp.body));
      }

      notifyListeners();
    } catch (e, stackTrace) {
      print(stackTrace);
      print(e);
    }

    return resp;
  }

  List<User> getCachedListOfUsers() {
    if (this._cachedUserList == null) {
      return [];
    }

    if (this._cachedUserList.users == null) {
      return [];
    }

    return this._cachedUserList.users;
  }
}

class UserUpdateResponse {
  String status;
  User user;

  UserUpdateResponse();

  UserUpdateResponse.fromJson(Map<String, dynamic> jsonObj) {
    this.status = jsonObj['status'];
    this.user = User.fromJson(jsonObj['user']);
  }
}

class UsersList {
  String status;
  List<User> users = [];

  UsersList();

  UsersList.fromJson(Map<String, dynamic> jsonObj) {
    this.status = jsonObj['status'];
    List<dynamic> allUsers = jsonObj['users'];
    for (var i = 0; i < allUsers.length; i++) {
      var user = User.fromJson(allUsers[i]);
      this.users.add(user);
    }
  }
}

class LoginRequest {
  String _uniqueID;
  String _username;
  String _password;

  LoginRequest(this._uniqueID, this._username, this._password);

  String get uniqueID => _uniqueID;
  String get password => _password;
  String get username => _username;

  Map<String, dynamic> toJson() {
    Map<String, dynamic> jsonMap = Map<String, dynamic>();
    jsonMap["uniqueID"] = _uniqueID;
    jsonMap["username"] = _username;
    jsonMap["password"] = _password;
    return jsonMap;
  }

  //LoginRequest - Constructor to initialize with a json object
  LoginRequest.fromJson(Map<String, dynamic> json)
      : _uniqueID = json['uniqueID'],
        _username = json['username'],
        _password = json['password'];
}

class Login {
  String status;
  String sessionToken;
  User user;
  LoginRequest loginRequest;

  Login(this.status, this.sessionToken);

  Login.fromJson(Map<String, dynamic> json)
      : status = json['status'],
        sessionToken = json['sessionToken'],
        user = User.fromJson(json);

  bool isLoggedIn() {
    if (user != null && user.loggedIn != null) {
      return user.loggedIn;
    }

    return false;
  }
}

/*
 * We will make user immutable.
 */
class User {
  String id;
  String username;
  String name;
  bool isThing;
  String secret;
  bool loggedIn;
  List<Permission> permissions = [];
  List<String> roles = [];
  String password;
  String confirmPassword;

  User(this.username);

  User.fromJson(Map<String, dynamic> json) {
    if (json['name'] != null) {
      this.name = json['name'];
    }

    name = json['fullName'];
    username = json['userName'];
    secret = json['secret'];
    isThing = json['isThing'];

    if (json['id'] != null) {
      id = json['id'];
    }

    List<dynamic> allPermissions = json['permissions'];
    if (allPermissions != null) {
      for (var i = 0; i < allPermissions.length; i++) {
        var perm = Permission.fromJson(allPermissions[i]);
        this.permissions.add(perm);
      }
    }

    List<dynamic> allRoles = json['roles'];
    if (allRoles != null) {
      for (var i = 0; i < allRoles.length; i++) {
        var roleID = allRoles[i];
        this.roles.add(roleID);
      }
    }
  }

  toJson() {
    Map<String, dynamic> jsonObj = {};
    if (this.confirmPassword != null) {
      jsonObj['confirmPassword'] = this.confirmPassword;
    }

    if (this.id != null && this.id.isNotEmpty) {
      jsonObj['id'] = this.id;
    }

    jsonObj['isThing'] = isThing;

    if (this.name != null && this.name.isNotEmpty) {
      jsonObj['name'] = this.name;
    }

    if (this.secret != null) {
      jsonObj['secret'] = this.secret;
    }

    jsonObj['userName'] = this.username;

    if (this.password != null) {
      jsonObj['password'] = this.password;
    }

    if (this.permissions != null) {
      List<dynamic> jsonPermissions = [];
      for (var i = 0; i < this.permissions.length; i++) {
        var perm = this.permissions[i].toJson();
        jsonPermissions.add(perm);
      }
      jsonObj['permissions'] = jsonPermissions;
    }

    if (this.roles != null) {
      jsonObj['roles'] = this.roles;
    }

    return jsonObj;
  }
}

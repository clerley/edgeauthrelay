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
import 'dart:io';
import 'dart:async';

import 'package:flutter/material.dart';

String host = "192.168.10.94";
int port = 9119;
String path = "/jwt/company/login";
String uri = "http://192.168.10.94:9119/jwt/company/login";



class UserProvider extends ChangeNotifier {
  static final UserProvider _theInstance = UserProvider._privateConstructor();
  
  UserProvider._privateConstructor();

  factory UserProvider() {
    return _theInstance;
  }


  static var login = Login("Failure", "---");

  Future<Login> requestLogin(String uniqueID, String username, String password) async {
    var loginRequest = LoginRequest(uniqueID, username, password);
    HttpClient httpClient = HttpClient();
    HttpClientRequest httpRequest = await httpClient.postUrl(Uri.parse(uri));
    httpRequest.headers.add("content-type", "application/json");
    httpRequest.add(utf8.encode(json.encode(loginRequest.toJson())));
    var httpResponse = await httpRequest.close();
    if(httpResponse.statusCode == HttpStatus.ok) {
        httpResponse.transform(utf8.decoder).listen((content) {
          login = Login.fromJson(json.decode(content));
          UserProvider.login.user = User(username);
          UserProvider.login.sessionToken = login.sessionToken;
          UserProvider.login.status = login.status;
        });

    } else {
        login = Login("Failure", "---");
        UserProvider.login.user.loggedIn = false;
    }

    notifyListeners();
    return login;
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

  Map<String, dynamic> toJson() =>
  {
    "uniqueID" : _uniqueID,
    "username" : _username,
    "password" : _password
  };

  //LoginRequest - Constructor to initialize with a json object
  LoginRequest.fromJson(Map<String, dynamic> json) : 
    _uniqueID = json['uniqueID'],
    _username = json['username'],
    _password = json['password'];
  
}

class Login {
  String status;
  String sessionToken;
  User user;
  LoginRequest loginRequest;

  Login(this.status, this.sessionToken);

  Login.fromJson(Map<String, dynamic> json):
    status = json['status'],
    sessionToken = json['sessionToken'],
    user = User.fromJson(json);
    

}


/*
 * We will make user immutable.
 */
class User {
  String username;
  String name;
  bool isThing;
  String secret;
  bool loggedIn;

  User(this.username); 

  User.fromJson(Map<String, dynamic> json):
    name = json['fullName'],
    username = json['userName'];
}
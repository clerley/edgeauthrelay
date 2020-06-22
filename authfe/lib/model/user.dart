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
import 'package:authfe/model/settings.dart';
import 'package:http/http.dart' as http;

import 'package:flutter/material.dart';
import 'package:http/http.dart';

class UserProvider extends ChangeNotifier {
  static final UserProvider _theInstance = UserProvider._privateConstructor();
  
  UserProvider._privateConstructor();

  factory UserProvider() {
    return _theInstance;
  }


  var login = Login("Failure", "---");

  Future<Login> requestLogin(String uniqueID, String username, String password) async {
    try {
      //var val = await http.get("https://jsonplaceholder.typicode.com/albums/1");
      //print(val.body);
      GlobalSettings globalSettings = GlobalSettings();
      var loginRequest = LoginRequest(uniqueID, username, password);
      var fullURL = globalSettings.url + "/jwt/company/login";
      String bodyStr = json.encode(loginRequest.toJson());
      var response = await http.post(fullURL, 
                    headers: {'Content-type': 'application/json'}, 
                    body: bodyStr);
      if(response.statusCode == 200) {
        login = Login.fromJson(json.decode(response.body));
        if(login.status == "Success") {
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

    } catch (e, stackTrace)  {
      print("An error occurred while processing the login request ${stackTrace.toString()}");
      print(e.toString());
    }
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

  Map<String, dynamic> toJson() {
    Map<String, dynamic> jsonMap = Map<String, dynamic>();
    jsonMap["uniqueID"] = _uniqueID;
    jsonMap["username"] = _username;
    jsonMap["password"] = _password;
    return jsonMap;
  }

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


  bool isLoggedIn() {
    if(user != null && user.loggedIn != null){
      return user.loggedIn;
    }

    return false;
  }
    

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
    username = json['userName'],
    secret = json['secret'],
    isThing = json['isThing'];
}
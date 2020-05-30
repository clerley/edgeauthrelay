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
import 'dart:core';
import 'dart:developer';

import 'package:authfe/model/settings.dart';
import 'package:flutter/cupertino.dart';
import 'package:http/http.dart' as http;

class AddCompanyResponse {
  String status;
  String companyID; //DB ID.

  AddCompanyResponse.fromJson(Map<String, dynamic> jsonObj):
    this.status = jsonObj['status'],
    this.companyID = jsonObj['companyID'];
  
}

class CompanyProvider extends ChangeNotifier {
  static final CompanyProvider _theInstance =
      CompanyProvider._internalConstructor();

  CompanyProvider._internalConstructor();

  factory CompanyProvider() {
    return _theInstance;
  }

  Future<bool> addCompany(Company company) async {
    var result = false; 
    try {
      var url = "/jwt/company";
      var globalSettings = GlobalSettings();
      var fullURL = globalSettings.url + url;
      var jsonMap = company.toJson();
      var encodedJson = json.encode(jsonMap);
      Map<String, String> headers = new Map<String, String>();
      headers["Content-Type"] = "application/json";
      var response = await http.post(fullURL, headers: headers,  body: encodedJson);
      if (response.statusCode == 200) {
        AddCompanyResponse acr = AddCompanyResponse.fromJson(json.decode(response.body));
        if (acr.status == "Success") {
          result = true;
          company.companyID = acr.companyID;
        } else {
          log("The Company was not correctly created!");
          result = false;
        }
      } else {
        log("The following error code was returned: [${response.statusCode}]");
      }
    } catch (e/*, stack*/) {
      //print(stack.toString());
      print(e.toString());
    }
    notifyListeners();
    return result;
  }

}

class PasswordHandler {
  String password;
  String confirmPassword;

  PasswordHandler(): password = "", confirmPassword = "";
}

class Company {
  String companyID = "";
  String uniqueID = "";
  String name = "";
  String address1 = "";
  String address2 = "";
  String city = "";
  String state = "";
  String zip = "";
  bool isLocation = false;
  bool remotelyManaged = false;
  String authRelay = "";
  int jwtDuration = 30;
  int passwordExpiration = 5;
  String passwordUnit = "Minute";
  String groupOwnerID = "";
  List<String> memberOfGroups;
  PasswordHandler passwordHandler;

  Company() {
    this.memberOfGroups = List<String>();
    this.passwordHandler = PasswordHandler();
    jwtDuration = 30;
    passwordExpiration = 5;
    passwordUnit = "Minute";
  }

  Company.fromJson(Map<String, dynamic> jsonObj) 
      : this.uniqueID = jsonObj['uniqueID'],
        this.companyID = jsonObj['companyID'],
        this.name = jsonObj['name'],
        this.address1 = jsonObj['address1'],
        this.address2 = jsonObj['address2'],
        this.authRelay = jsonObj['authRelay'],
        this.city = jsonObj['city'],
        this.jwtDuration = jsonObj['settings']['jwtDuration'],
        this.isLocation = jsonObj['isInLocation'],
        this.passwordExpiration = jsonObj['settings']['passExpiration'],
        this.passwordUnit = jsonObj['settings']['passUnit'],
        this.remotelyManaged = jsonObj['remotelyManaged'],
        this.groupOwnerID = jsonObj['groupOwnerID'],
        this.memberOfGroups = jsonObj['memberOfGroups'],
        this.state = jsonObj['state'],
        this.zip = jsonObj['zip'];

  Map<String, dynamic> toJson() {
    var jsonObj = Map<String, dynamic>();
    jsonObj['uniqueID'] = this.uniqueID;
    jsonObj['name'] = this.name;
    jsonObj['address1'] = this.address1;
    jsonObj['address2'] = this.address2;
    jsonObj['authRelay'] = this.authRelay;
    jsonObj['city'] = this.city;
    jsonObj['settings'] = Map<String, dynamic>();
    jsonObj['settings']['jwtDuration'] = this.jwtDuration;
    jsonObj['settings']['passUnit'] = this.passwordUnit;
    jsonObj['settings']['passExpiration'] = this.passwordExpiration;
    jsonObj['remotelyManaged'] = this.remotelyManaged.toString();
    jsonObj['groupOwnerID'] = this.groupOwnerID;
    jsonObj['memberOfGroups'] = this.memberOfGroups;
    jsonObj['state'] = this.state;
    jsonObj['zip'] = this.zip;
    jsonObj['isInLocation'] = this.isLocation.toString();
    if(this.passwordHandler != null) {
      jsonObj['password'] = this.passwordHandler.password;
      jsonObj['confirmPassword'] = this.passwordHandler.confirmPassword;
    } else {
      log("The value of passwordHandler is not currently populated!");
    }
    return jsonObj;
  }
}

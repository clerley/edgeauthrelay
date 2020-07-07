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

import 'package:authfe/model/settingsmodel.dart';
import 'package:authfe/model/usermodel.dart';
import 'package:flutter/cupertino.dart';
import 'package:http/http.dart' as http;

class AddCompanyResponse {
  String status;
  String companyID; //DB ID.

  AddCompanyResponse.fromJson(Map<String, dynamic> jsonObj)
      : this.status = jsonObj['status'],
        this.companyID = jsonObj['companyID'];
}

class CompanyProvider extends ChangeNotifier {
  static final CompanyProvider _theInstance =
      CompanyProvider._internalConstructor();

  CompanyProvider._internalConstructor();

  factory CompanyProvider() {
    return _theInstance;
  }

  String companyID = "";
  GetCompanyResponse editCompanyResponse;

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
      var response =
          await http.post(fullURL, headers: headers, body: encodedJson);
      if (response.statusCode == 200) {
        AddCompanyResponse acr =
            AddCompanyResponse.fromJson(json.decode(response.body));
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
    } catch (e /*, stack*/) {
      //print(stack.toString());
      print(e.toString());
    }
    notifyListeners();
    return result;
  }

  Future<UpdateCompanyResponse> updateCompany(Company company) async {
    UpdateCompanyResponse ucr = UpdateCompanyResponse();
    ucr.status = "Failure";
    try {
      var url = "/jwt/company/${company.uniqueID}";
      var globalSettings = GlobalSettings();
      UserProvider userProvider = UserProvider();
      var fullURL = globalSettings.url + url;
      var jsonMap = company.toJson();
      var encodedJson = json.encode(jsonMap);
      Map<String, String> headers = new Map<String, String>();
      headers["Content-Type"] = "application/json";
      headers["Authorization"] = "bearer ${userProvider.login.sessionToken}";
      var response =
          await http.post(fullURL, headers: headers, body: encodedJson);
      if (response.statusCode == 200) {
        UpdateCompanyResponse updtCompanyResponse =
            UpdateCompanyResponse.fromJson(json.decode(response.body));
        if (updtCompanyResponse.status == "Success") {
          ucr = updtCompanyResponse;
        }
      } else {
        log("The following error code was returned: [${response.statusCode}]");
      }
    } catch (e /*, stack*/) {
      //print(stack.toString());
      print(e.toString());
    }
    notifyListeners();
    return ucr;
  }

  Future<GetCompanyResponse> getCompany(String uniqueID) async {
    GetCompanyResponse companyResponse = GetCompanyResponse();
    companyResponse.status = "Failure";

    try {
      GlobalSettings settings = GlobalSettings();
      UserProvider userProvider = UserProvider();

      if (!userProvider.login.isLoggedIn()) {
        return null;
      }

      var fullPath = settings.url + "/jwt/company/$uniqueID";
      var httpHeader = {
        "Authorization": "bearer ${userProvider.login.sessionToken}"
      };

      var response = await http.get(fullPath, headers: httpHeader);
      if (response.statusCode != 200) {
        print("The statusCode was not expected: ${response.statusCode}");
        return companyResponse;
      }

      var jsn = json.decode(response.body);
      companyResponse = GetCompanyResponse.fromJson(jsn);
    } catch (e, stackTrace) {
      print(stackTrace);
      print(e);
    }

    return companyResponse;
  }

  Future<GetGroupResponse> getGroupForCompanyID(String companyID) async {
    GetGroupResponse groupResponse = GetGroupResponse();
    groupResponse.status = "Failure";

    try {
      GlobalSettings settings = GlobalSettings();
      UserProvider userProvider = UserProvider();

      if (!userProvider.login.isLoggedIn()) {
        return null;
      }

      var fullPath = settings.url + "/companies/$companyID";
      var httpHeader = {
        "Authorization": "bearer ${userProvider.login.sessionToken}"
      };

      var response = await http.get(fullPath, headers: httpHeader);
      if (response.statusCode != 200) {
        print("The statusCode was not expected: ${response.statusCode}");
        return groupResponse;
      }

      var jsn = json.decode(response.body);
      groupResponse = GetGroupResponse.fromJson(jsn);
    } catch (e, stackTrace) {
      print(stackTrace);
      print(e);
    }

    return groupResponse;
  }

  Future<CompanyRegistrationResponse> enableRegistration(
      Company company) async {
    CompanyRegistrationResponse ucr = CompanyRegistrationResponse();
    ucr.status = "Failure";
    try {
      var url = "/company/registration/${company.companyID}";
      var globalSettings = GlobalSettings();
      UserProvider userProvider = UserProvider();
      var fullURL = globalSettings.url + url;
      debugPrint('Sending a request to URL:[$fullURL]');
      Map<String, String> headers = new Map<String, String>();
      headers["Content-Type"] = "application/json";
      headers["Authorization"] = "bearer ${userProvider.login.sessionToken}";
      var response = await http.post(fullURL, headers: headers);
      if (response.statusCode == 200) {
        CompanyRegistrationResponse regResp =
            CompanyRegistrationResponse.fromJson(json.decode(response.body));
        if (regResp.status == "Success") {
          ucr = regResp;
          company.regisCode = ucr.regisCode;
        }
      } else {
        debugPrint(
            "The following error code was returned: [${response.statusCode}]");
      }
    } catch (e /*, stack*/) {
      //print(stack.toString());
      print(e.toString());
    }
    notifyListeners();
    return ucr;
  }
}

class CompanyRegistrationResponse {
  String status;
  String regisCode;

  CompanyRegistrationResponse();

  CompanyRegistrationResponse.fromJson(Map<String, dynamic> jsonObj) {
    this.status = jsonObj['status'];
    if (jsonObj['regisCode'] != null) {
      this.regisCode = jsonObj['regisCode'];
    } else {
      debugPrint("The response did not contain a valid registration code");
      this.status = "Failure";
    }
  }
}

class UpdateCompanyResponse {
  String status;

  UpdateCompanyResponse.fromJson(Map<String, dynamic> jsonObj) {
    this.status = jsonObj['status'];
  }

  UpdateCompanyResponse();
}

class GetCompanyResponse {
  String status;
  Company company;

  GetCompanyResponse() {
    company = new Company();
  }

  GetCompanyResponse.fromJson(Map<String, dynamic> jsnMap) {
    this.status = jsnMap["status"];
    this.company = Company.fromJson(jsnMap);
  }

  toJson() {
    Map<String, dynamic> map = this.company.toJson();
    map['status'] = this.status;
    return map;
  }
}

class GetGroupResponse {
  String status;
  List<Company> companies = [];

  GetGroupResponse();

  GetGroupResponse.fromJson(Map<String, dynamic> jsonObj) {
    this.status = jsonObj['status'];
    List<dynamic> allCompanies = jsonObj['companies'];
    for (var i = 0; i < allCompanies.length; i++) {
      var c = Company.fromJson(allCompanies[i]);
      companies.add(c);
    }
  }
}

class PasswordHandler {
  String password;
  String confirmPassword;

  PasswordHandler()
      : password = "",
        confirmPassword = "";
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
  String regisCode = "";
  String apiKey = "";

  Company() {
    this.memberOfGroups = List<String>();
    this.passwordHandler = PasswordHandler();
    jwtDuration = 30;
    passwordExpiration = 5;
    passwordUnit = "Minute";
  }

  Company.fromJson(Map<String, dynamic> jsonObj) {
    this.uniqueID = jsonObj['uniqueID'];
    this.companyID = jsonObj['companyID'];
    this.name = jsonObj['name'];
    this.address1 = jsonObj['address1'];
    this.address2 = jsonObj['address2'];
    this.authRelay = jsonObj['authRelay'];
    this.city = jsonObj['city'];
    this.jwtDuration = jsonObj['settings']['jwtDuration'];
    if (jsonObj['isInLocation'] is bool) {
      this.isLocation = jsonObj['isInLocation'];
    } else {
      this.isLocation = jsonObj['isInLocation'] == "true" ? true : false;
    }
    this.passwordExpiration = jsonObj['settings']['passExpiration'];
    this.passwordUnit = jsonObj['settings']['passUnit'];
    if (jsonObj['remotelyManaged'] is bool) {
      this.remotelyManaged = jsonObj['remotelyManaged'];
    } else {
      this.remotelyManaged =
          jsonObj['remotelyManaged'] == "true" ? true : false;
    }
    this.groupOwnerID = jsonObj['groupOwnerID'];
    this.memberOfGroups = jsonObj['memberOfGroups'];
    this.state = jsonObj['state'];
    this.zip = jsonObj['zip'];
    if (jsonObj['regisCode'] != null) {
      this.regisCode = jsonObj['regisCode'];
    }
    if (jsonObj['apiKey'] != null) {
      this.apiKey = jsonObj['apiKey'];
    }
  }

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
    if (this.passwordHandler != null) {
      jsonObj['password'] = this.passwordHandler.password;
      jsonObj['confirmPassword'] = this.passwordHandler.confirmPassword;
    } else {
      log("The value of passwordHandler is not currently populated!");
    }
    if (this.apiKey != null && this.apiKey.isNotEmpty) {
      jsonObj['apiKey'] = this.apiKey;
    }
    return jsonObj;
  }

  bool isInsertable() {
    if (this.companyID == null || this.companyID.length == 0) {
      return true;
    }

    return false;
  }

  String getFullAddress() {
    String fullAddr = "";
    if (this.address1 != null && this.address1.isNotEmpty) {
      fullAddr = this.address1;
    }

    if (this.address2 != null && this.address2.isNotEmpty) {
      fullAddr = fullAddr + "\n" + this.address2;
    }

    if (this.city != null && this.city.isNotEmpty) {
      fullAddr = fullAddr + "\n" + this.city;
    }

    if (this.state != null && this.state.isNotEmpty) {
      fullAddr = fullAddr + " " + this.state;
    }

    if (this.zip != null && this.zip.isNotEmpty) {
      fullAddr = fullAddr + " " + this.zip;
    }

    return fullAddr;
  }
}

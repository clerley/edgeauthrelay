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

class GlobalSettings {
  static final GlobalSettings _theInstance = GlobalSettings._internal();
  String url;
  String companyUniqueID; //This is the ID provided by the user
  String companyID; //This is the database ID.
  bool _loaded = false;

  GlobalSettings._internal();

  factory GlobalSettings() {
    if (!_theInstance._loaded) {
      _theInstance.load();
    }
    return _theInstance;
  }

  Future<void> save() async {
    //Directory appDocDir = await getApplicationDocumentsDirectory();
    //String appDocPath = appDocDir.path;
    String edgeAuth = "config.json";
    File file = File(edgeAuth);
    _Settings settings = _Settings();
    settings.url = url;
    settings.companyID = companyID;
    settings.companyUniqueID = companyUniqueID;
    file.writeAsStringSync(json.encode(settings.toJson()));
    /*await _prefs.setString('url', url);
    await _prefs.setString('companyUniqueID', companyUniqueID);
    await _prefs.setString('companyID', companyID);*/
  }

  Future<void> load() async {
    //Directory appDocDir = await getApplicationDocumentsDirectory();
    //String appDocPath = appDocDir.path;
    String edgeAuth = "config.json";
    File file = File(edgeAuth);
    String encodedJson = file.readAsStringSync();
    var jsonObj = json.decode(encodedJson);
    _Settings settings = _Settings.fromJson(jsonObj);
    this.url = settings.url;
    this.companyID = settings.companyID;
    this.companyUniqueID = settings.companyUniqueID;

    //If this work I will remove the fields from the GlobalSettings object
    ///and just create a reference to settings.

    /*this.url = _prefs.getString('url');
    this.companyUniqueID = _prefs.getString('companyUniqueID');
    this.companyID = _prefs.getString('companyID');
    this._loaded = true;*/
    if (this.url == null || this.url.isEmpty) {
      this.url = "http://127.0.0.1:9119";
    }
  }
}

class _Settings {
  String url = "http://127.0.0.1:9119";
  String companyUniqueID = "";
  String companyID = "";

  toJson() {
    Map<String, dynamic> jsonObj = {};
    jsonObj['url'] = url;
    jsonObj['companyUniqueID'] = companyUniqueID;
    jsonObj['companyID'] = companyID;
    return jsonObj;
  }

  _Settings();

  _Settings.fromJson(Map<String, dynamic> jsonObj) {
    this.url = jsonObj['url'];
    this.companyID = jsonObj['companyID'];
    this.companyUniqueID = jsonObj['companyUniqueID'];
  }
}

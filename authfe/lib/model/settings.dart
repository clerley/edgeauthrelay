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

import 'package:shared_preferences/shared_preferences.dart';

class GlobalSettings {

  static final GlobalSettings _theInstance = GlobalSettings._internal();
  String url;
  bool _loaded = false;

  GlobalSettings._internal();

  factory GlobalSettings() {
    if(!_theInstance._loaded) {
      _theInstance.load();
    }
    return _theInstance;
  }

  Future<void> save() async {
    
    SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.setString('url', url);

  }

  Future<void> load() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    this.url = prefs.getString('url');
    this._loaded = true;
  }

}


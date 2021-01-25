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
import 'package:authfe/i18n/language.dart';
import 'package:authfe/model/settingsmodel.dart';
import 'package:flutter/material.dart';

import '../main.dart';

/*
 * SettingsView ...
 */
class SettingsView extends StatefulWidget {
  @override
  State<StatefulWidget> createState() => _SettingsViewState();
}

class _SettingsViewState extends State<SettingsView> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(getText("title", LANG_ENGLISH)),
      ),
      body: SingleChildScrollView(
        child: _SettingsViewBody(),
      ),
    );
  }
}

class _SettingsViewBody extends StatefulWidget {
  @override
  State<StatefulWidget> createState() => _SettingsViewBodyState();
}

class _SettingsViewBodyState extends State<_SettingsViewBody> {
  TextEditingController _urlController = TextEditingController();

  _SettingsViewBodyState() {
    _urlController.text = GlobalSettings().url;
  }

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Container(
        margin: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 0.0),
        padding: EdgeInsets.all(10.0),
        width: 900.0,
        decoration: BoxDecoration(
            color: Theme.of(context).backgroundColor,
            borderRadius: BorderRadius.all(Radius.circular(10.0))),
        child: Column(
            mainAxisAlignment: MainAxisAlignment.start,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: <Widget>[
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                child: Text(
                  "Settings",
                  style: Theme.of(context).primaryTextTheme.bodyText1,
                ),
              ),
              Container(
                child: Text("URL"),
              ),
              Container(
                child: TextField(
                  style: Theme.of(context).primaryTextTheme.bodyText2,
                  controller: _urlController,
                ),
              ),
              Center(
                child: Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: <Widget>[
                    Container(
                      padding: EdgeInsets.all(5.0),
                      child: OutlinedButton(
                        style: ButtonStyle(
                          foregroundColor:
                              MaterialStateProperty.all<Color>(Colors.white),
                        ),
                        child: Text(
                          getText("save", LANG_ENGLISH),
                          style: Theme.of(context).primaryTextTheme.button,
                        ),
                        onPressed: () async {
                          var gset = GlobalSettings();
                          gset.url = _urlController.text;
                          await gset.save();
                          Navigator.pushReplacement(
                              context,
                              MaterialPageRoute(
                                  builder: (context) => MyHomePage(
                                      title: getText("title", LANG_ENGLISH))));
                        },
                      ),
                    ),
                    Container(
                      padding: EdgeInsets.all(5.0),
                      child: OutlinedButton(
                        style: ButtonStyle(
                          foregroundColor:
                              MaterialStateProperty.all<Color>(Colors.white),
                        ),
                        onPressed: () {
                          Navigator.pushReplacement(
                              context,
                              MaterialPageRoute(
                                  builder: (context) => MyHomePage(
                                      title: getText("title", LANG_ENGLISH))));
                        },
                        child: Text(
                          getText("cancel", LANG_ENGLISH),
                          style: Theme.of(context).primaryTextTheme.button,
                        ),
                      ),
                    ),
                  ],
                ),
              ),
            ]),
      ),
    );
  }
}

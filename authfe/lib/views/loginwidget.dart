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

import 'package:authfe/model/companymodel.dart';
import 'package:authfe/model/usermodel.dart';
import 'package:authfe/views/company.dart';
import 'package:authfe/views/viewhelper.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../i18n/language.dart';
import 'mainmenu.dart';

class LoginWidget extends StatefulWidget {
  @override
  _LoginState createState() => _LoginState(LANG_ENGLISH);
}

class _LoginState extends State<LoginWidget> {
  String _language = LANG_ENGLISH;
  String _userNameText;
  String _passwordText;
  String _appName;
  String _loginText;
  String _uniqueCompanyID;
  String _newCompany;
  TextEditingController _uniqueIDController = TextEditingController();
  TextEditingController _usernameController = TextEditingController();
  TextEditingController _passwordController = TextEditingController();

  _LoginState(this._language) {
    this._userNameText = getText("username", _language);
    this._passwordText = getText("password", _language);
    this._appName = getText("title", _language);
    this._loginText = getText("login", _language);
    this._uniqueCompanyID = getText("uniqueID", _language);
    this._newCompany = getText("newcompany", _language);
  }

  @override
  Widget build(BuildContext context) {
    var pr = DialogHelper()
        .createProgressDialog(getText("please_wait", this._language), context);
    var userProvider = Provider.of<UserProvider>(context);
    return Center(
        child: Container(
      margin: EdgeInsets.all(16.0),
      padding: EdgeInsets.all(16.0),
      decoration: BoxDecoration(
          color: Theme.of(context).backgroundColor,
          borderRadius: BorderRadius.all(Radius.circular(10.0))),
      width: 400.0,
      height: 505.0,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.stretch,
        children: <Widget>[
          Container(
            alignment: Alignment.center,
            padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
            child: Text(this._appName,
                style: TextStyle(fontSize: 26, color: Colors.white)),
          ),
          Container(
            padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
            child: Text(this._uniqueCompanyID,
                style: Theme.of(context).primaryTextTheme.bodyText1),
          ),
          Container(
            padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
            child: TextField(
                style: Theme.of(context).primaryTextTheme.bodyText2,
                controller: _uniqueIDController),
          ),
          Container(
            padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
            child: Text(this._userNameText,
                style: Theme.of(context).primaryTextTheme.bodyText1),
          ),
          Container(
            padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 20.0),
            child: TextField(
                style: Theme.of(context).primaryTextTheme.bodyText2,
                controller: _usernameController),
          ),
          Container(
            padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 20.0),
            child: Text(this._passwordText,
                style: Theme.of(context).primaryTextTheme.bodyText1),
          ),
          TextField(
            obscureText: true,
            style: Theme.of(context).primaryTextTheme.bodyText2,
            controller: _passwordController,
          ),
          Divider(
            color: Colors.white,
            thickness: 1.0,
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
                      this._loginText,
                      style: Theme.of(context).primaryTextTheme.button,
                    ),
                    onPressed: () async {
                      await pr.show();
                      debugPrint("Starting the onPressed request now");
                      var login = await userProvider.requestLogin(
                          _uniqueIDController.text,
                          _usernameController.text,
                          _passwordController.text);
                      if (pr.isShowing()) {
                        await pr.hide();
                      }
                      if (login.isLoggedIn()) {
                        Navigator.pushReplacement(
                          context,
                          MaterialPageRoute(
                              builder: (context) => MainMenu(this._language)),
                        );
                      } else {
                        debugPrint(
                            "The user is not logged in now! ${_uniqueIDController.text}");
                        DialogHelper pdh = DialogHelper();
                        pdh.showMessageDialog(
                            getText("user_not_logged", this._language),
                            context,
                            this._language);
                      }
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
                        CompanyProvider cp = CompanyProvider();
                        cp.editCompanyResponse = null;
                        Navigator.push(
                          context,
                          MaterialPageRoute(
                              builder: (context) =>
                                  CompanyWidget(this._language)),
                        );
                      },
                      child: Text(
                        this._newCompany,
                        style: Theme.of(context).primaryTextTheme.button,
                      ),
                    )),
              ],
            ),
          ),
        ],
      ),
    ));
  }
}

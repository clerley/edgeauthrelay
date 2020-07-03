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

import 'package:authfe/appbar/menudrawer.dart';
import 'package:authfe/i18n/language.dart';
import 'package:authfe/views/viewhelper.dart';
import 'package:flutter/material.dart';

import 'mainmenu.dart';
import '../model/usermodel.dart';

class UpdatePasswordView extends StatefulWidget {
  final String _language;

  UpdatePasswordView(this._language);

  @override
  State<StatefulWidget> createState() => _UpdatePasswordViewState(_language);
}

class _UpdatePasswordViewState extends State<UpdatePasswordView> {
  final String _language;

  _UpdatePasswordViewState(this._language);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(getText("title", this._language)),
      ),
      body: SingleChildScrollView(
        child: _UpdatePasswordViewBody(this._language),
      ),
      drawer: DistAuthDrawer(this._language),
    );
  }
}

class _UpdatePasswordViewBody extends StatefulWidget {
  final String _language;

  _UpdatePasswordViewBody(this._language);

  @override
  State<StatefulWidget> createState() =>
      _UpdatePasswordViewBodyState(this._language);
}

class _UpdatePasswordViewBodyState extends State<_UpdatePasswordViewBody> {
  TextEditingController _currentPassword = TextEditingController();
  TextEditingController _newPassword = TextEditingController();
  TextEditingController _confirmPassword = TextEditingController();
  String _username;

  String _language;

  _UpdatePasswordViewBodyState(this._language);

  @override
  void initState() {
    UserProvider usersProvider = new UserProvider();
    if (usersProvider.edittingUser != null) {
      this._username = '(${usersProvider.edittingUser.username})';
    } else {
      this._username = "";
    }
    super.initState();
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
              Row(
                children: [
                  Container(
                    padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                    child: Text(
                      getText("change-password", this._language),
                      style: Theme.of(context).primaryTextTheme.bodyText1,
                    ),
                  ),
                  Container(
                      padding: EdgeInsets.fromLTRB(10.0, 0.0, 0.0, 10.0),
                      child: Text(
                        _username,
                        style: Theme.of(context).primaryTextTheme.bodyText1,
                      )),
                ],
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
                child: Text(getText("password", this._language)),
              ),
              Container(
                child: TextField(
                  style: Theme.of(context).primaryTextTheme.bodyText2,
                  controller: _currentPassword,
                  obscureText: true,
                ),
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
                child: Text(getText("new-password", this._language)),
              ),
              Container(
                child: TextField(
                  style: Theme.of(context).primaryTextTheme.bodyText2,
                  controller: _newPassword,
                  obscureText: true,
                ),
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
                child: Text(getText("confirmPassword", this._language)),
              ),
              Container(
                child: TextField(
                  style: Theme.of(context).primaryTextTheme.bodyText2,
                  controller: _confirmPassword,
                  obscureText: true,
                ),
              ),
              Center(
                child: Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: <Widget>[
                    Container(
                      padding: EdgeInsets.all(5.0),
                      child: OutlineButton(
                          textColor: Colors.white,
                          child: Text(
                            getText("save", LANG_ENGLISH),
                            style: Theme.of(context).primaryTextTheme.button,
                          ),
                          onPressed: () {
                            _updatePassword();
                          },
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(30.0),
                          )),
                    ),
                    Container(
                        padding: EdgeInsets.all(5.0),
                        child: OutlineButton(
                          textColor: Colors.white,
                          onPressed: () {
                            Navigator.pushReplacement(
                                context,
                                MaterialPageRoute(
                                    builder: (context) =>
                                        MainMenu(this._language)));
                          },
                          child: Text(
                            getText("cancel", LANG_ENGLISH),
                            style: Theme.of(context).primaryTextTheme.button,
                          ),
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(30.0),
                          ),
                        )),
                  ],
                ),
              ),
            ]),
      ),
    );
  }

  _updatePassword() async {
    DialogHelper dh = DialogHelper();
    var msg = getText('password-changed', this._language);
    var req = new UpdatePasswordRequest();
    req.confirmPassword = _confirmPassword.text;
    req.currentPassword = _currentPassword.text;
    req.newPassword = _newPassword.text;

    //Get the provider....
    UserProvider userProvider = UserProvider();
    try {
      if (userProvider.edittingUser == null) {
        req.username = userProvider.login.user.username;
      } else {
        req.username = userProvider.edittingUser.username;
      }
      var resp = await userProvider.updatePassword(req);
      if (!resp) {
        debugPrint('The password has not changed, please try again');
        msg = getText('pwd-update-failed', this._language);
      }
    } catch (e) {
      debugPrint('The following exception occurred: $e');
      msg = getText('pwd-update-failed', this._language);
    }

    dh.showMessageDialog(msg, context, this._language);
  }
}

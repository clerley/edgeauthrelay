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

import 'package:authfe/model/usermodel.dart';
import 'package:authfe/views/companyview.dart';
import 'package:authfe/views/updatepassword.dart';
import 'package:flutter/material.dart';
import '../i18n/language.dart';
import '../views/permissions.dart';
import '../views/roles.dart';
import '../views/users.dart';
import '../main.dart';

class DistAuthDrawer extends StatelessWidget {
  final String _language;

  DistAuthDrawer(this._language);

  @override
  Widget build(BuildContext context) {
    return Drawer(
      // Add a ListView to the drawer. This ensures the user can scroll
      // through the options in the drawer if there isn't enough vertical
      // space to fit everything.
      child: ListView(
        // Important: Remove any padding from the ListView.
        padding: EdgeInsets.zero,
        children: <Widget>[
          DrawerHeader(
            child: Text(
              getText("title", this._language),
              style: TextStyle(fontSize: 22.0),
            ),
            decoration: BoxDecoration(
              color: Theme.of(context).backgroundColor,
            ),
          ),
          ListTile(
            title: Text(getText("roles", this._language)),
            onTap: () {
              Navigator.pushReplacement(
                  context,
                  MaterialPageRoute(
                      builder: (context) => RolesView(_language)));
            },
          ),
          ListTile(
            title: Text(getText("permissions", this._language)),
            onTap: () {
              Navigator.pushReplacement(
                context,
                MaterialPageRoute(
                    builder: (context) => PermissionsView(_language)),
              );
            },
          ),
          ListTile(
            title: Text(getText("users", this._language)),
            onTap: () {
              Navigator.pushReplacement(
                  context,
                  MaterialPageRoute(
                      builder: (context) => UsersView(_language)));
            },
          ),
          ListTile(
            title: Text(getText("change-password", this._language)),
            onTap: () {
              Navigator.pushReplacement(
                context,
                MaterialPageRoute(
                    builder: (context) => UpdatePasswordView(this._language)),
              );
            },
          ),
          ListTile(
            title: Text(getText("company", this._language)),
            onTap: () {
              Navigator.pushReplacement(
                context,
                MaterialPageRoute(
                    builder: (context) => CompanyViewOnly(this._language)),
              );
            },
          ),
          ListTile(
            title: Text(getText("logout", this._language)),
            onTap: () async {
              UserProvider userProvider = UserProvider();
              debugPrint('UserProvider logged out');
              await userProvider.logout();
              Navigator.pushReplacement(
                context,
                MaterialPageRoute(
                    builder: (context) =>
                        MyHomePage(title: getText("title", this._language))),
              );
            },
          ),
        ],
      ),
    );
  }
}

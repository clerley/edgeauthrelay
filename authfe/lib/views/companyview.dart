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
import 'package:authfe/model/companymodel.dart';
import 'package:authfe/model/usermodel.dart';
import 'package:authfe/views/companysubsidiaries.dart';
import 'package:authfe/views/viewhelper.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'company.dart';
import 'mainmenu.dart';

class CompanyViewOnly extends StatefulWidget {
  final String _language;

  CompanyViewOnly(this._language);

  @override
  State<StatefulWidget> createState() => _CompanyViewState(this._language);
}

class _CompanyViewState extends State<CompanyViewOnly> {
  final String _language;

  _CompanyViewState(this._language);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: Text(getText("title", this._language)),
        ),
        body: SingleChildScrollView(
          child: _CompanyViewBody(this._language),
        ),
        drawer: DistAuthDrawer(this._language));
  }
}

class _CompanyViewBody extends StatefulWidget {
  final String _language;

  _CompanyViewBody(this._language);

  @override
  State<StatefulWidget> createState() => _CompanyViewBodyState(this._language);
}

class _CompanyViewBodyState extends State<_CompanyViewBody> {
  final String _language;

  _CompanyViewBodyState(this._language) {
    _uniqueID = TextEditingController();
    _name = TextEditingController();
    _address1 = TextEditingController();
    _address2 = TextEditingController();
    _state = TextEditingController();
    _city = TextEditingController();
    _zip = TextEditingController();
  }

  CompanyProvider companyProvider;

  TextEditingController _uniqueID;
  TextEditingController _name;
  TextEditingController _address1;
  TextEditingController _address2;
  TextEditingController _state;
  TextEditingController _city;
  TextEditingController _zip;

  @override
  void initState() {
    CompanyProvider companyProvider = CompanyProvider();
    UserProvider userProvider = UserProvider();
    if (userProvider.login.isLoggedIn()) {
      if (userProvider.login.loginRequest != null &&
          userProvider.login.loginRequest.uniqueID != null) {
        companyProvider
            .getCompany(userProvider.login.loginRequest.uniqueID)
            .then<GetCompanyResponse>((value) {
          if (value is GetCompanyResponse) {
            _getCompanyResponseCallback(value);
          }
          return value;
        });
      }
    }

    super.initState();
  }

  _getCompanyResponseCallback(GetCompanyResponse companyResponse) {
    if (companyResponse.status == "Success") {
      print("Received the company response!");
      companyProvider.editCompanyResponse = companyResponse;

      setState(() {
        if (companyResponse.company.uniqueID != null) {
          _uniqueID.text = companyResponse.company.uniqueID;
        }
        _name.text = companyResponse.company.name;
        _address1.text = companyResponse.company.address1;
        _address2.text = companyResponse.company.address2;
        _state.text = companyResponse.company.state;
        _city.text = companyResponse.company.city;
        _zip.text = companyResponse.company.zip;
      });
    } else {
      companyProvider.editCompanyResponse = null;
    }
  }

  @override
  Widget build(BuildContext context) {
    companyProvider = Provider.of<CompanyProvider>(context);
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
              padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
              child: Text(
                getText("company", this._language),
                style: Theme.of(context).primaryTextTheme.bodyText1,
              ),
            ),
            Container(
              padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
              child: Text(getText("uniqueID", this._language)),
            ),
            Container(
              child: TextField(
                style: Theme.of(context).primaryTextTheme.bodyText2,
                controller: _uniqueID,
              ),
            ),
            Container(
              padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
              child: Text(
                getText("name", this._language),
              ),
            ),
            Container(
              child: TextField(
                style: Theme.of(context).primaryTextTheme.bodyText2,
                controller: _name,
              ),
            ),
            Container(
              padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
              child: Text(
                getText("address", this._language),
              ),
            ),
            Container(
              padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
              child: TextField(
                  style: Theme.of(context).primaryTextTheme.bodyText2,
                  controller: _address1),
            ),
            Container(
              padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
              child: TextField(
                style: Theme.of(context).primaryTextTheme.bodyText2,
                controller: _address2,
              ),
            ),
            Row(
              children: <Widget>[
                Expanded(
                  flex: 24,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                    child: Text(
                      getText("city", this._language),
                    ),
                  ),
                ),
                Spacer(),
                Expanded(
                  flex: 12,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                    child: Text(
                      getText("state", this._language),
                    ),
                  ),
                ),
                Spacer(),
                Expanded(
                  flex: 24,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                    child: Text(
                      getText("zip", this._language),
                    ),
                  ),
                ),
              ],
            ),
            Row(
              children: <Widget>[
                Expanded(
                  flex: 24,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                    child: TextField(
                      style: Theme.of(context).primaryTextTheme.bodyText2,
                      controller: _city,
                    ),
                  ),
                ),
                Spacer(),
                Expanded(
                  flex: 12,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                    child: TextField(
                        style: Theme.of(context).primaryTextTheme.bodyText2,
                        controller: _state),
                  ),
                ),
                Spacer(),
                Expanded(
                  flex: 24,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                    child: TextField(
                      style: Theme.of(context).primaryTextTheme.bodyText2,
                      controller: _zip,
                    ),
                  ),
                ),
              ],
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: <Widget>[
                Container(
                  padding: EdgeInsets.fromLTRB(0.0, 0.0, 12.0, 0.0),
                  child: OutlineButton(
                    textColor: Colors.white,
                    child: Text(
                      getText("edit", this._language),
                      style: Theme.of(context).primaryTextTheme.button,
                    ),
                    onPressed: () {
                      _editCompany();
                    },
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(30.0),
                    ),
                  ),
                ),
                Container(
                  padding: EdgeInsets.fromLTRB(0.0, 0.0, 12.0, 0.0),
                  child: OutlineButton(
                    textColor: Colors.white,
                    child: Text(
                      getText("cancel", this._language),
                      style: Theme.of(context).primaryTextTheme.button,
                    ),
                    onPressed: () {
                      companyProvider.editCompanyResponse = null;
                      Navigator.pushReplacement(
                          context,
                          MaterialPageRoute(
                              builder: (context) => MainMenu(this._language)));
                    },
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(30.0),
                    ),
                  ),
                ),
                OutlineButton(
                    textColor: Colors.white,
                    child: Text(
                      getText("list_mgr_company", this._language),
                      style: Theme.of(context).primaryTextTheme.button,
                    ),
                    onPressed: () {
                      _showSubsidiaries();
                    },
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(30.0),
                    )),
              ],
            ),
          ],
        ),
      ),
    );
  }

  _editCompany() {
    var companyProvider = CompanyProvider();
    if (companyProvider.editCompanyResponse != null) {
      Navigator.pushReplacement(
          context,
          MaterialPageRoute(
              builder: (context) => CompanyWidget(this._language)));
    } else {
      DialogHelper dh = DialogHelper();
      dh.showMessageDialog(
          getText("", this._language), context, this._language);
    }
  }

  _showSubsidiaries() {
    var companyProvider = CompanyProvider();
    if (companyProvider.editCompanyResponse != null) {
      companyProvider.companyID =
          companyProvider.editCompanyResponse.company.companyID;
      Navigator.pushReplacement(
          context,
          MaterialPageRoute(
            builder: (context) => CompanySubsidiariesView(this._language),
          ));
      return;
    }

    DialogHelper dialogHelper = DialogHelper();
    dialogHelper.showMessageDialog(
        getText("group_not_found", this._language), context, this._language);
  }
}

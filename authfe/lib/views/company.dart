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

import 'dart:developer';

import 'package:authfe/appbar/menudrawer.dart';
import 'package:authfe/model/companymodel.dart';
import 'package:authfe/views/companyview.dart';
import 'package:authfe/views/viewhelper.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:uuid/uuid.dart';
import 'package:uuid/uuid_util.dart';
import '../i18n/language.dart';

class CompanyWidget extends StatefulWidget {
  final String _language;

  CompanyWidget(this._language);

  @override
  _CompanyWidgetState createState() => _CompanyWidgetState(this._language);
}

class _CompanyWidgetState extends State<CompanyWidget> {
  String _language;
  String _title;
  _CompanyWidgetState(this._language) {
    this._title = getText("title", this._language);
  }

  @override
  Widget build(BuildContext context) {
    CompanyProvider companyProvider = CompanyProvider();

    if (companyProvider.editCompanyResponse == null) {
      return Scaffold(
          appBar: AppBar(
            title: Text(this._title),
          ),
          body: SingleChildScrollView(
            child: CompanyBody(this.widget._language),
          ));
    }

    return Scaffold(
      appBar: AppBar(
        title: Text(this._title),
      ),
      body: SingleChildScrollView(
        child: CompanyBody(this.widget._language),
      ),
      drawer: DistAuthDrawer(this._language),
    );
  }
}

class CompanyBody extends StatefulWidget {
  final String _language;

  CompanyBody(this._language);

  @override
  State<StatefulWidget> createState() => _CompanyState(_language);
}

class _CompanyState extends State<CompanyBody> {
  String _language;
  String _company;
  String _addressText;
  String _cityText;
  String _stateText;
  String _zipText;
  String _nameText;
  String _isLocationText;
  String _remotelyManagedText;
  bool _remoteAuth;
  bool _isLocation;
  String _unit;
  List<DropdownMenuItem<String>> _unitsMenuItems = [];
  Company company;

  TextEditingController _uniqueIDTec;
  TextEditingController _nameTec;
  TextEditingController _addressTec;
  TextEditingController _cityTec;
  TextEditingController _stateTec;
  TextEditingController _zipTec;
  TextEditingController _authRelayTec;
  TextEditingController _jwtDurationTec;
  TextEditingController _passExpTec;
  TextEditingController _address1Tec;
  TextEditingController _passwordTec;
  TextEditingController _confirmPasswordTec;
  TextEditingController _groupOwnerText;
  TextEditingController _regisCodeText;
  TextEditingController _apiKeyText;

  _CompanyState(this._language) {
    this._company = getText("company", _language);
    this._cityText = getText("city", _language);
    this._stateText = getText("state", _language);
    this._zipText = getText("zip", _language);
    this._nameText = getText("name", _language);
    this._addressText = getText("address", _language);
    this._isLocationText = getText("isLocation", _language);
    this._remotelyManagedText = getText("remotelyManaged", _language);
    this._remoteAuth = false;
    this._isLocation = false;
    this._unit = "Minute";

    company = Company();
    _uniqueIDTec = TextEditingController();
    _nameTec = TextEditingController();
    _addressTec = TextEditingController();
    _cityTec = TextEditingController();
    _stateTec = TextEditingController();
    _zipTec = TextEditingController();
    _authRelayTec = TextEditingController();
    _jwtDurationTec = TextEditingController();
    _passExpTec = TextEditingController();
    _address1Tec = TextEditingController();
    _passwordTec = TextEditingController();
    _confirmPasswordTec = TextEditingController();
    _groupOwnerText = TextEditingController();
    _regisCodeText = TextEditingController();
    _apiKeyText = TextEditingController();

    //Here we will check if we are editting the company
    var companyProvider = CompanyProvider();
    if (companyProvider.editCompanyResponse != null) {
      company = companyProvider.editCompanyResponse.company;
      companyProvider.editCompanyResponse = null;
      //print(company.toJson());
      print('Using the editable company response!');
    }
  }

  //_CompanyState.withCompany(this._language, this._company);

  List<DropdownMenuItem<String>> _getUnitMenuItems() {
    List<DropdownMenuItem<String>> lst = [];
    DropdownMenuItem<String> item =
        new DropdownMenuItem<String>(child: Text("Hour"), value: "Hour");
    lst.add(item);

    item = new DropdownMenuItem<String>(child: Text("Day"), value: "Day");
    lst.add(item);

    item = new DropdownMenuItem<String>(child: Text("Week"), value: "Week");
    lst.add(item);

    item = new DropdownMenuItem<String>(child: Text("Month"), value: "Month");
    lst.add(item);

    item = new DropdownMenuItem<String>(child: Text("Minute"), value: "Minute");
    lst.add(item);

    debugPrint("The number of elements: ${lst.length}");
    return lst;
  }

  @override
  void initState() {
    this._unitsMenuItems = _getUnitMenuItems();

    if (company == null) {
      log("The company object is null, creating a new company now!");
      if (companyProvider.editCompanyResponse != null &&
          companyProvider.editCompanyResponse.company != null) {
        company = companyProvider.editCompanyResponse.company;
        companyProvider.editCompanyResponse = null;
      } else {
        company = Company();
      }
    }
    _uniqueIDTec.text = company.uniqueID;
    _nameTec.text = company.name;
    _addressTec.text = company.address1;
    _cityTec.text = company.city;
    _stateTec.text = company.state;
    _zipTec.text = company.zip;
    _authRelayTec.text = company.authRelay;
    _jwtDurationTec.text = company.jwtDuration.toString();
    _passExpTec.text = company.passwordExpiration.toString();
    _address1Tec.text = company.address2;
    _groupOwnerText.text = company.groupOwnerID;
    _regisCodeText.text = company.regisCode;
    _apiKeyText.text = company.apiKey;

    _isLocation = company.isLocation;
    _remoteAuth = company.remotelyManaged;
    _apiKeyText.text = company.apiKey;

    _unit = company.passwordUnit;
    if (_unit == null || _unit.isEmpty) {
      _unit = "Month";
    }

    super.initState();
  }

  CompanyProvider companyProvider;

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
              padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
              child: Text(
                this._company,
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
                controller: _uniqueIDTec,
              ),
            ),
            Container(
              padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
              child: Text(
                this._nameText,
              ),
            ),
            Container(
              child: TextField(
                style: Theme.of(context).primaryTextTheme.bodyText2,
                controller: _nameTec,
              ),
            ),
            Container(
              padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 0.0),
              child: Text(
                this._addressText,
              ),
            ),
            Container(
              padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
              child: TextField(
                  style: Theme.of(context).primaryTextTheme.bodyText2,
                  controller: _addressTec),
            ),
            Container(
              padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
              child: TextField(
                style: Theme.of(context).primaryTextTheme.bodyText2,
                controller: _address1Tec,
              ),
            ),
            Row(
              children: <Widget>[
                Expanded(
                  flex: 24,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                    child: Text(
                      this._cityText,
                    ),
                  ),
                ),
                Spacer(),
                Expanded(
                  flex: 12,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                    child: Text(
                      this._stateText,
                    ),
                  ),
                ),
                Spacer(),
                Expanded(
                  flex: 24,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                    child: Text(this._zipText),
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
                      controller: _cityTec,
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
                        controller: _stateTec),
                  ),
                ),
                Spacer(),
                Expanded(
                  flex: 24,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                    child: TextField(
                      style: Theme.of(context).primaryTextTheme.bodyText2,
                      controller: _zipTec,
                    ),
                  ),
                ),
              ],
            ),
            Container(
              padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
              child: Row(
                children: <Widget>[
                  Checkbox(
                    onChanged: (val) {
                      setState(() {
                        this._isLocation = val;
                      });
                    },
                    value: this._isLocation,
                  ),
                  Text(this._isLocationText),
                  Checkbox(
                    onChanged: (val) {
                      setState(() {
                        this._remoteAuth = val;
                      });
                    },
                    value: this._remoteAuth,
                  ),
                  Text(this._remotelyManagedText),
                ],
              ),
            ),
            Container(
              padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
              child: Text(
                getText("authrelay", this._language),
              ),
            ),
            Container(
              child: TextField(
                style: Theme.of(context).primaryTextTheme.bodyText2,
                controller: _authRelayTec,
              ),
            ),
            Container(
              padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
              child: Text(
                getText("apikey", this._language),
              ),
            ),
            Row(
              children: [
                Expanded(
                  child: Container(
                    child: TextField(
                      style: Theme.of(context).primaryTextTheme.bodyText2,
                      controller: _apiKeyText,
                    ),
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.all(8.0),
                  child: OutlinedButton(
                    style: ButtonStyle(
                      foregroundColor:
                          MaterialStateProperty.all<Color>(Colors.white),
                    ),
                    child: Icon(Icons.autorenew),
                    onPressed: () {
                      _generateApiKey();
                    },
                  ),
                ),
              ],
            ),
            Container(
                margin: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 0.0),
                padding: EdgeInsets.all(10.0),
                width: 900.0,
                decoration: BoxDecoration(
                  color: Theme.of(context).backgroundColor,
                  borderRadius: BorderRadius.all(Radius.circular(10.0)),
                  border: Border.all(color: Color(0xff506d90)),
                ),
                child: Column(
                  children: <Widget>[
                    Row(
                      children: <Widget>[
                        Expanded(
                          flex: 24,
                          child: Padding(
                            padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                            child: Text(getText("jwtDuration", this._language)),
                          ),
                        ),
                        Spacer(),
                        Expanded(
                          flex: 24,
                          child: Padding(
                            padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                            child: Text(getText("passwordExp", this._language)),
                          ),
                        ),
                        Spacer(),
                        Expanded(
                          flex: 24,
                          child: Padding(
                            padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                            child:
                                Text(getText("passwordUnit", this._language)),
                          ),
                        )
                      ],
                    ),
                    Row(
                      children: <Widget>[
                        Expanded(
                          flex: 24,
                          child: TextField(
                            keyboardType: TextInputType.number,
                            style: Theme.of(context).primaryTextTheme.bodyText2,
                            controller: _jwtDurationTec,
                          ),
                        ),
                        Spacer(),
                        Expanded(
                          flex: 24,
                          child: TextField(
                            keyboardType: TextInputType.number,
                            style: Theme.of(context).primaryTextTheme.bodyText2,
                            controller: _passExpTec,
                          ),
                        ),
                        Spacer(),
                        Expanded(
                          flex: 24,
                          //child: TextField(keyboardType: TextInputType.number,),
                          child: new DropdownButton<String>(
                            value: _unit,
                            items: _unitsMenuItems,
                            onChanged: (String newValue) {
                              setState(() {
                                _unit = newValue;
                              });
                            },
                          ),
                        )
                      ],
                    )
                  ],
                )),
            Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Container(
                  padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
                  child: Text(
                    getText('password', this._language),
                  ),
                ),
                TextField(
                  controller: _passwordTec,
                  style: Theme.of(context).primaryTextTheme.bodyText2,
                  obscureText: true,
                ),
                Container(
                    padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
                    child: Text(
                      getText('confirmPassword', this._language),
                    )),
                TextField(
                  controller: _confirmPasswordTec,
                  style: Theme.of(context).primaryTextTheme.bodyText2,
                  obscureText: true,
                ),
              ],
            ),
            Container(
              margin: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 0.0),
              padding: EdgeInsets.all(10.0),
              width: 900.0,
              decoration: BoxDecoration(
                color: Theme.of(context).backgroundColor,
                borderRadius: BorderRadius.all(Radius.circular(10.0)),
                border: Border.all(
                  color: Color(0xff506d90),
                ),
              ),
              child: Column(
                children: [
                  Row(
                    children: [
                      Expanded(
                        child: Container(
                          alignment: Alignment.centerLeft,
                          padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
                          child: Text(
                            getText('group-owner-id', this._language),
                          ),
                        ),
                      ),
                      Expanded(
                        child: Container(
                          alignment: Alignment.centerLeft,
                          padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
                          child: Text(
                            getText('regis-code', this._language),
                          ),
                        ),
                      ),
                    ],
                  ),
                  Row(
                    children: [
                      Expanded(
                        child: Container(
                          padding: EdgeInsets.fromLTRB(2.0, 1.0, 2.0, 1.0),
                          alignment: Alignment.centerLeft,
                          child: TextField(
                            controller: _groupOwnerText,
                            style: Theme.of(context).primaryTextTheme.bodyText2,
                          ),
                        ),
                      ),
                      Expanded(
                        child: Container(
                          padding: EdgeInsets.fromLTRB(2.0, 1.0, 2.0, 1.0),
                          alignment: Alignment.centerLeft,
                          child: TextField(
                            controller: _regisCodeText,
                            style: Theme.of(context).primaryTextTheme.bodyText2,
                          ),
                        ),
                      ),
                    ],
                  ),
                ],
              ),
            ),
            Row(
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
                      getText("save", this._language),
                      style: Theme.of(context).primaryTextTheme.button,
                    ),
                    onPressed: () {
                      saveCompany();
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
                    child: Text(
                      getText("cancel", this._language),
                      style: Theme.of(context).primaryTextTheme.button,
                    ),
                    onPressed: () {
                      if (company.isInsertable() && Navigator.canPop(context)) {
                        Navigator.pop(context);
                      } else {
                        Navigator.pushReplacement(
                            context,
                            MaterialPageRoute(
                                builder: (context) =>
                                    CompanyViewOnly(this._language)));
                      }
                    },
                  ),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }

  void saveCompany() async {
    company.address1 = _addressTec.text;
    company.address2 = _address1Tec.text;
    company.authRelay = _authRelayTec.text;
    company.city = _cityTec.text;
    company.isLocation = _isLocation;
    company.jwtDuration = int.parse(_jwtDurationTec.text);
    company.name = _nameTec.text;
    company.passwordExpiration = int.parse(_passExpTec.text);
    company.passwordUnit = _unit;
    company.remotelyManaged = _remoteAuth;
    company.state = _stateTec.text;
    company.uniqueID = _uniqueIDTec.text;
    company.zip = _zipTec.text;
    company.apiKey = _apiKeyText.text;
    if (companyProvider.companyID != null &&
        companyProvider.companyID.isNotEmpty) {
      company.groupOwnerID = companyProvider.companyID;
    }

    //We don't need the password if we are not inserting the company
    DialogHelper pdh = DialogHelper();
    var pd = pdh.createProgressDialog(
        getText('please_wait', this._language), context);

    if (company.isInsertable()) {
      company.passwordHandler.password = _passwordTec.text;
      company.passwordHandler.confirmPassword = _confirmPasswordTec.text;
      await pd.show();
      var rsp = await companyProvider.addCompany(company);
      await pd.hide();
      if (!rsp) {
        log("An error occured while adding a new company");
        pdh.showMessageDialog(
            getText("error_add_cmp", this._language), context, this._language);
      } else {
        pdh.showMessageDialog(
            getText("new_cmp_create", this._language), context, this._language);
      }
      //Ready to insert another company.
      _clearFieldsForInsert();
    } else {
      await pd.show();
      var rsp = await companyProvider.updateCompany(company);
      await pd.hide();
      if (rsp.status != "Success") {
        log("An error occured while adding a new company");
        pdh.showMessageDialog(
            getText("error_upt_cmp", this._language), context, this._language);
      } else {
        pdh.showMessageDialog(getText("company_updated", this._language),
            context, this._language);
      }
    }
  }

  _clearFieldsForInsert() {
    company = Company();

    setState(() {
      _uniqueIDTec.text = "";
      _nameTec.text = "";
      _addressTec.text = "";
      _cityTec.text = "";
      _stateTec.text = "";
      _zipTec.text = "";
      _authRelayTec.text = "";
      _jwtDurationTec.text = "30";
      _passExpTec.text = "";
      _address1Tec.text = "";
      _passwordTec.text = "";
      _confirmPasswordTec.text = "";
      _regisCodeText.text = "";
      _apiKeyText.text = "";
    });
  }

  _generateApiKey() {
    var uuid = Uuid();
    var v4Crypto = uuid.v4(options: {'rng': UuidUtil.cryptoRNG});
    _apiKeyText.text = v4Crypto.toString();
  }
}

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
import 'package:authfe/model/permissionmodel.dart';
import 'package:authfe/views/viewhelper.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../appbar/menudrawer.dart';
import '../i18n/language.dart';
import 'mainmenu.dart';
import 'searchpermission.dart';

class PermissionsView extends StatefulWidget {
  final String _language;
  final Permission _perm;

  PermissionsView(this._language) : _perm = null;

  PermissionsView.forEditing(this._language, this._perm);

  @override
  State<StatefulWidget> createState() {
    if (_perm != null) {
      return _PermissionsState.forEditting(this._language, this._perm);
    }

    return _PermissionsState(this._language);
  }
}

class _PermissionsState extends State<PermissionsView> {
  final String _language;
  final Permission _perm;

  _PermissionsState(this._language) : _perm = null;

  _PermissionsState.forEditting(this._language, this._perm);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(getText("title", this._language)),
      ),
      body: SingleChildScrollView(
        child: _PermissionBody.forEditing(this._language, this._perm),
      ),
      drawer: DistAuthDrawer(this._language),
    );
  }
}

class _PermissionBody extends StatefulWidget {
  final String _language;
  final Permission _perm;

  _PermissionBody(this._language) : _perm = null;

  _PermissionBody.forEditing(this._language, this._perm);

  @override
  State<StatefulWidget> createState() =>
      _PermissionBodyState.forEditting(this._language, this._perm);
}

class _PermissionBodyState extends State<_PermissionBody> {
  final String _language;
  TextEditingController _permController;
  TextEditingController _descrController;
  Permission _perm;

  _PermissionBodyState(this._language) {
    _perm = null;
  }

  _PermissionBodyState.forEditting(this._language, this._perm);

  @override
  void initState() {
    super.initState();
    this._permController = TextEditingController();
    this._descrController = TextEditingController();
    if (this._perm != null) {
      _permController.text = _perm.permission;
      _descrController.text = _perm.description;
    }
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
                  getText("permissions", this._language),
                  style: Theme.of(context).primaryTextTheme.bodyText1,
                ),
              ),
              Container(
                child: Text(getText("permission", this._language)),
              ),
              Container(
                child: TextField(
                  style: Theme.of(context).primaryTextTheme.bodyText2,
                  controller: _permController,
                ),
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 0.0),
                child: Text(
                  getText("description", this._language),
                ),
              ),
              Container(
                child: TextField(
                  style: Theme.of(context).primaryTextTheme.bodyText2,
                  controller: _descrController,
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
                            getText("add", this._language),
                            style: Theme.of(context).primaryTextTheme.button,
                          ),
                          onPressed: () {
                            addPermission();
                          },
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(30.0),
                          )),
                    ),
                    Container(
                      padding: EdgeInsets.all(5.0),
                      child: OutlineButton(
                          textColor: Colors.white,
                          child: Text(
                            getText("save", this._language),
                            style: Theme.of(context).primaryTextTheme.button,
                          ),
                          onPressed: () {
                            savePermission();
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
                                      SearchPermissions(this._language)),
                            );
                          },
                          child: Text(
                            getText("search", this._language),
                            style: Theme.of(context).primaryTextTheme.button,
                          ),
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(30.0),
                          ),
                        )),
                    Container(
                        padding: EdgeInsets.all(5.0),
                        child: OutlineButton(
                          textColor: Colors.white,
                          onPressed: () {
                            Navigator.pushReplacement(
                                context,
                                MaterialPageRoute(
                                    builder: (context) => MainMenu(this._language)));
                          },
                          child: Text(
                            getText("cancel", this._language),
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

  savePermission() async {
    if(this._perm != null) {
      this._perm.description = _descrController.text;
      this._perm.permission = _permController.text;
      PermissionProvider permProvider = Provider.of<PermissionProvider>(context);
      InsertPermissionResponse resp = await permProvider.updatePermission(this._perm);
      var pdh = DialogHelper();
      if(resp.status == "Success") {
        pdh.showMessageDialog(
          getText("perm_upd_success", this._language), context, this._language);
      } else {
        pdh.showMessageDialog(
          getText("perm_upd_error", this._language), context, this._language);
      }
    }
  }

  addPermission() async {
    Permission perm = Permission();
    perm.id = "";
    perm.description = _descrController.text;
    perm.permission = _permController.text;

    var pdh = DialogHelper();
    var pd = pdh.createProgressDialog(
        getText("please_wait", this._language), context);
    pd.show();
    PermissionProvider permProvider = Provider.of<PermissionProvider>(context);
    var resp = await permProvider.insertPermission(perm);
    if (resp.status == "Success") {
      _permController.text = "";
      _descrController.text = "";
      pdh.showMessageDialog(
          getText("perm_ins_success", this._language), context, this._language);
    } else {
      pdh.showMessageDialog(
          getText("perm_ins_error", this._language), context, this._language);
    }
    pd.hide();
  }
}

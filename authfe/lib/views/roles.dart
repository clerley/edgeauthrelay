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
import 'package:authfe/model/rolesmodel.dart';
import 'package:authfe/views/viewhelper.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../appbar/menudrawer.dart';
import '../i18n/language.dart';
import 'mainmenu.dart';
import 'searchrole.dart';

class RolesView extends StatefulWidget {
  final String _language;
  final Role _role;

  RolesView(this._language) : this._role = null;

  RolesView.withRole(this._language, this._role);

  @override
  State<StatefulWidget> createState() =>
      _RolesState(this._language, this._role);
}

class _RolesState extends State<RolesView> {
  final String _language;
  Role _role;

  _RolesState(this._language, this._role);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(getText("title", this._language)),
      ),
      body: SingleChildScrollView(
        child: _RoleBody(this._language, this._role),
      ),
      drawer: DistAuthDrawer(this._language),
    );
  }
}

class _RoleBody extends StatefulWidget {
  final String _language;
  final Role _role;

  _RoleBody(this._language, this._role);

  @override
  State<StatefulWidget> createState() =>
      _RoleBodyState.withRole(this._role, this._language);
}

class _RoleBodyState extends State<_RoleBody> {
  final String _language;

  Role role;
  TextEditingController _description;

  _RoleBodyState(this._language) {
    this.role = Role();
  }

  _RoleBodyState.withRole(this.role, this._language) {
    if (this.role == null) {
      this.role = Role();
    }
  }

  @override
  void initState() {
    _description = TextEditingController();
    if (this.role != null && this.role.description != null) {
      _description.text = role.description;
    }

    var permissionProvider = PermissionProvider();
    Future.sync(() async => await permissionProvider.listPermissions(0, 1000));

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
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                child: Text(
                  getText("roles", this._language),
                  style: Theme.of(context).primaryTextTheme.bodyText1,
                ),
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                child: Text(getText("description", this._language)),
              ),
              Container(
                child: TextField(
                  style: Theme.of(context).primaryTextTheme.bodyText2,
                  controller: _description,
                ),
              ),
              Center(
                child: Consumer<PermissionProvider>(
                    builder: (context, permissionProvider, child) {
                  return DataTable(
                    onSelectAll: (value) => allRolesSelected(value),
                    columns: [
                      DataColumn(label: Text("")),
                      DataColumn(
                          label: Text(getText("description", this._language)))
                    ],
                    rows: _getDataSource(),
                  );
                }),
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
                            addRole();
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
                            updateRole();
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
                                      SearchRoles(this._language)),
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
                                    builder: (context) =>
                                        MainMenu(this._language)));
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

  List<DataRow> _getDataSource() {
    PermissionProvider permissionsProvider = PermissionProvider();
    if (!permissionsProvider.isCached()) {
      Future.sync(
          () async => await permissionsProvider.listPermissions(0, 10000));
    }

    List<Permission> cachedPermissions =
        permissionsProvider.getCachedPermissions();
    var dataRows = [];
    if (cachedPermissions.length > 0) {
      for (var i = 0; i < cachedPermissions.length; i++) {
        var permission = cachedPermissions[i];
        var row = DataRow(
          selected: role.hasPermission(permission),
          cells: [],
          onSelectChanged: (value) {
            permissionSelected(permission, value);
          },
        );
        DataCell cell = new DataCell(Text(permission.permission));
        row.cells.add(cell);

        cell = new DataCell(Text(permission.description));
        row.cells.add(cell);
        dataRows.add(row);
      }
    }
    return dataRows;
  }

  bool permissionSelected(Permission perm, bool value) {
    if (role != null) {
      if (role.hasPermission(perm)) {
        role.removePermission(perm);
      } else {
        role.addPermission(perm);
      }
    }

    PermissionProvider permProv = PermissionProvider();
    permProv.doNotification();

    return value;
  }

  void addPermission(Permission perm) {
    if (this.role != null) {
      this.role.addPermission(perm);
    }
  }

  addRole() {
    _updateRole("add");
  }

  updateRole() {
    _updateRole("update");
  }

  _updateRole(String updateType) async {
    bool success = false;
    role.description = _description.text;
    RolesProvider rolesProvider = RolesProvider();

    var response;
    if (this.role != null &&
        updateType != "update" &&
        this.role.isInsertable()) {
      response = await rolesProvider.insertRole(role);
    } else if (this.role != null &&
        updateType == "update" &&
        !this.role.isInsertable()) {
      response = await rolesProvider.saveRole(role);
    }

    if (response != null && response.status == "Success") {
      success = true;
    }

    String msg;
    DialogHelper dialogHelper = DialogHelper();
    if (updateType == "update") {
      msg = getText("update_role_failed", this._language);
      if (success) {
        msg = getText("update_role_success", this._language);
      }
    } else {
      msg = getText("add_role_failed", this._language);
      if (success) {
        msg = getText("add_role_success", this._language);
      }
    }

    dialogHelper.showMessageDialog(msg, context, this._language);
  }

  allRolesSelected(bool isSelected) {
    if (this.role == null) {
      return;
    }

    PermissionProvider permissionsProvider = PermissionProvider();
    var cachedPermissions = permissionsProvider.getCachedPermissions();
    if (cachedPermissions != null) {
      for (var i = 0; i < cachedPermissions.length; i++) {
        if (isSelected) {
          this.role.addPermission(cachedPermissions[i]);
        } else {
          this.role.removePermission(cachedPermissions[i]);
        }
      }
    }

    permissionsProvider.doNotification();
  }
}

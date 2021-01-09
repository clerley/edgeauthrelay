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
import 'package:authfe/model/usermodel.dart';
import 'package:authfe/views/mainmenu.dart';
import 'package:authfe/views/updatepassword.dart';
import 'package:authfe/views/viewhelper.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../appbar/menudrawer.dart';
import '../i18n/language.dart';
import 'searchuser.dart';

class UsersView extends StatefulWidget {
  final String _language;

  UsersView(this._language);

  @override
  State<StatefulWidget> createState() => _UsersState(this._language);
}

class _UsersState extends State<UsersView> {
  final String _language;

  _UsersState(this._language);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(getText("title", this._language)),
      ),
      body: SingleChildScrollView(
        child: _UserBody(this._language),
      ),
      drawer: DistAuthDrawer(this._language),
    );
  }
}

class _UserBody extends StatefulWidget {
  final String _language;

  _UserBody(this._language);

  @override
  State<StatefulWidget> createState() => _UserBodyState(this._language);
}

class _UserBodyState extends State<_UserBody> {
  final String _language;
  User _user = User("");
  TextEditingController _username = TextEditingController();
  TextEditingController _name = TextEditingController();
  TextEditingController _password = TextEditingController();
  TextEditingController _confirmPassword = TextEditingController();
  TextEditingController _secret = TextEditingController();
  bool _isThing = false;
  List<DropdownMenuItem<String>> _roles = [];
  String _roleSelected = "";
  List<Permission> _permissions = [];
  List<DataRow> _rows = [];

  _UserBodyState(this._language);

  Future<List<DropdownMenuItem<String>>> _loadRolesDropDown() async {
    List<DropdownMenuItem<String>> lst = [];
    RolesProvider rolesProvider = RolesProvider();
    if (!rolesProvider.isCached()) {
      await rolesProvider.listRoles(0, 5000);
    }
    List<Role> roles = rolesProvider.getCached();
    if (roles == null) {
      return lst;
    }

    DropdownMenuItem<String> item =
        new DropdownMenuItem<String>(child: Text(""), value: "");
    lst.add(item);

    for (var i = 0; i < roles.length; i++) {
      var role = roles[i];
      item = new DropdownMenuItem<String>(
          child: Text(role.description), value: role.id);
      lst.add(item);
      if (_user.hasRole(role.id)) {
        _roleSelected = role.id;
      }
    }

    setState(() {
      this._roles = lst;
    });

    return lst;
  }

  _loadPermissions() async {
    PermissionProvider permProvider = PermissionProvider();
    await permProvider.listPermissions(0, 1000);

    if (!permProvider.isCached()) {
      return;
    }

    List<DataRow> tmpLst = [];
    this._permissions = permProvider.getCachedPermissions();
    for (var i = 0; i < _permissions.length; i++) {
      var perm = _permissions[i];
      DataRow dw = DataRow(
        cells: [],
        onSelectChanged: (value) => _addPermission(perm),
        selected: _isPermission(perm),
      );
      DataCell dc = DataCell(
        Text(perm.permission),
      );
      dw.cells.add(dc);
      dc = DataCell(
        Text(perm.description),
      );
      dw.cells.add(dc);
      tmpLst.add(dw);
    }

    setState(() {
      this._rows = tmpLst;
    });
  }

  @override
  void initState() {
    UserProvider usersProv = UserProvider();
    if (usersProv.edittingUser != null) {
      this._user = usersProv.edittingUser;
      usersProv.edittingUser = null;
      _username.text = _user.username;
      _name.text = _user.name;
      _secret.text = _user.secret;
      _isThing = _user.isThing;
      var rolesList = _user.listRoles();
      if (rolesList.length > 0) {
        _roleSelected = rolesList[0];
      }
    }
    super.initState();
    _loadRolesDropDown();
    _loadPermissions();
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
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 5.0),
                child: Text(
                  getText("users", this._language),
                  style: Theme.of(context).primaryTextTheme.bodyText1,
                ),
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
                child: Text(getText("username", this._language)),
              ),
              Row(
                children: [
                  Expanded(
                    child: Container(
                      child: TextField(
                        style: Theme.of(context).primaryTextTheme.bodyText2,
                        controller: _username,
                      ),
                    ),
                  ),
                  OutlineButton(
                    textColor: Colors.white,
                    child: Stack(
                      children: <Widget>[
                        Align(
                            alignment: Alignment.centerLeft,
                            child: Icon(Icons.security)),
                      ],
                    ),
                    onPressed: () {
                      _changeOtherUserPassword();
                    },
                    shape: CircleBorder(),
                  ),
                ],
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
                child: Text(getText("name", this._language)),
              ),
              Container(
                child: TextField(
                    style: Theme.of(context).primaryTextTheme.bodyText2,
                    controller: _name),
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
              ),
              Row(children: [
                Container(
                    child: Checkbox(
                        onChanged: (bool value) {
                          setState(() {
                            this._isThing = value;
                          });
                        },
                        value: _isThing)),
                Container(
                  child: Text(getText("isthing", this._language)),
                ),
              ]),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 5.0),
                child: Text(getText("roles", this._language)),
              ),
              Row(
                mainAxisAlignment: MainAxisAlignment.start,
                mainAxisSize: MainAxisSize.max,
                children: [
                  Container(
                    width: 500,
                    child: DropdownButton<String>(
                      value: _roleSelected,
                      items: _roles,
                      onChanged: (String newValue) {
                        setState(() {
                          _roleSelected = newValue;
                        });
                      },
                    ),
                  ),
                ],
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 5.0),
                child: Text(getText("permissions", this._language)),
              ),
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Consumer<UserProvider>(
                    builder: (context, userProvider, child) {
                      return DataTable(
                        columns: [
                          DataColumn(
                            label: Text(
                              getText("permission", this._language),
                            ),
                          ),
                          DataColumn(
                            label: Text(
                              getText("description", this._language),
                            ),
                          ),
                        ],
                        rows: _rows,
                      );
                    },
                  ),
                ],
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 5.0),
                child: Text(getText("password", this._language)),
              ),
              Container(
                child: TextField(
                  style: Theme.of(context).primaryTextTheme.bodyText2,
                  obscureText: true,
                  controller: _password,
                ),
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 5.0),
                child: Text(getText("confirmPassword", this._language)),
              ),
              Container(
                child: TextField(
                  style: Theme.of(context).primaryTextTheme.bodyText2,
                  obscureText: true,
                  controller: _confirmPassword,
                ),
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 5.0),
                child: Text(getText("secret", this._language)),
              ),
              Container(
                child: TextField(
                    style: Theme.of(context).primaryTextTheme.bodyText2,
                    controller: _secret),
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
                            _addUser();
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
                            _saveUser();
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
                                      SearchUsers(this._language)),
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

  _changeOtherUserPassword() {
    if (this._user.isInsertable) {
      debugPrint('The user is insertable, the password cannot be changed!');
      return;
    }

    UserProvider provider = UserProvider();
    if (provider.edittingUser == null) {
      provider.edittingUser = this._user;
    }

    Navigator.pushReplacement(
        context,
        MaterialPageRoute(
          builder: (context) => UpdatePasswordView(this._language),
        ));
  }

  _clearFieldsForNewAddition() {
    this._isThing = false;
    this._name.text = "";
    this._username.text = "";
    this._secret.text = "";
    this._password.text = "";
    this._confirmPassword.text = "";
    this._user = User("");
  }

  //This excludes the username/password/confirmPassword.
  //Those fields only make sense on inserts
  _readGenericUserInfo() {
    _user.isThing = _isThing;
    _user.name = _name.text;
    _user.addRole(_roleSelected);
    _user.secret = _secret.text;
  }

  _addUser() async {
    if (_user == null) {
      _user = User("");
    }
    _readGenericUserInfo();
    _user.username = _username.text;
    _user.password = _password.text;
    _user.confirmPassword = _confirmPassword.text;
    if (_user.isInsertable) {
      UserProvider userProvider = UserProvider();
      var rsp = await userProvider.insertUser(_user);
      if (rsp.status == "Success") {
        _clearFieldsForNewAddition();
        return;
      }
    }

    DialogHelper dh = DialogHelper();
    dh.showMessageDialog(
        getText("add_user_failed", this._language), context, this._language);
  }

  _saveUser() async {
    if (_user == null) {
      debugPrint('The user cannot be saved because it has not been defined');
      return;
    }

    String msg = getText("updt_user_success", this._language);
    _readGenericUserInfo();
    if (_user.isUpdatable) {
      UserProvider userProvider = UserProvider();
      var rsp = await userProvider.updateUser(_user);
      if (rsp.status != "Success") {
        msg = getText("updt_user_failed", this._language);
      }
    }

    DialogHelper dh = DialogHelper();
    dh.showMessageDialog(msg, context, this._language);
  }

  _addPermission(Permission perm) {
    if (_user.hasPermission(perm)) {
      _user.removePermission(perm);
    } else {
      _user.addPermission(perm);
    }

    _loadPermissions();
  }

  bool _isPermission(Permission perm) {
    if (_user == null) {
      return false;
    }

    return _user.hasPermission(perm);
  }
}

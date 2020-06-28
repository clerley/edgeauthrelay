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
import 'package:authfe/views/mainmenu.dart';
import 'package:flutter/material.dart';
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
  User user = User("");

  _UserBodyState(this._language);

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
              Container(
                child: TextField(
                    style: Theme.of(context).primaryTextTheme.bodyText2),
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
                child: Text(getText("name", this._language)),
              ),
              Container(
                child: TextField(
                    style: Theme.of(context).primaryTextTheme.bodyText2),
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
              ),
              Row(children: [
                Container(
                    child: Checkbox(onChanged: (bool value) {}, value: true)),
                Container(
                  child: Text(getText("isthing", this._language)),
                ),
              ]),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 5.0),
                child: Text(getText("password", this._language)),
              ),
              Container(
                child: TextField(
                    style: Theme.of(context).primaryTextTheme.bodyText2,
                    obscureText: true,),
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 5.0),
                child: Text(getText("confirmPassword", this._language)),
              ),
              Container(
                child: TextField(
                    style: Theme.of(context).primaryTextTheme.bodyText2,
                    obscureText: true,),
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 5.0),
                child: Text(getText("secret", this._language)),
              ),
              Container(
                child: TextField(
                    style: Theme.of(context).primaryTextTheme.bodyText2),
              ),
              Center(
                child: DataTable(
                  columns: [
                    DataColumn(label: Text("")),
                    DataColumn(label: Text(getText("name", this._language)))
                  ],
                  rows: _getDataSource(),
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
                          onPressed: () {},
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
                          onPressed: () {},
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

  bool _test1Checked = false;
  List<DataRow> _getDataSource() {
    var dataRows = List<DataRow>();
    var row = DataRow(cells: []);
    DataCell cell = new DataCell(Checkbox(
        onChanged: (bool value) {
          setState(() {
            this._test1Checked = value;
          });
        },
        value: this._test1Checked));
    row.cells.add(cell);
    cell = new DataCell(Text('Testing 1'));
    row.cells.add(cell);
    dataRows.add(row);
    return dataRows;
  }
}

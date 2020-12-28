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

import 'package:authfe/model/rolesmodel.dart';
import 'package:authfe/views/roles.dart';
import 'package:flutter/material.dart';
import '../i18n/language.dart';
import '../appbar/menudrawer.dart';

class SearchRoles extends StatefulWidget {
  final String _language;

  SearchRoles(this._language);

  @override
  State<StatefulWidget> createState() => _SearchRolesState(this._language);
}

class _SearchRolesState extends State<SearchRoles> {
  final String _language;

  _SearchRolesState(this._language);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(getText("title", this._language)),
      ),
      body: SingleChildScrollView(
        child: _SearchRoleBody(this._language),
      ),
      drawer: DistAuthDrawer(this._language),
    );
  }
}

class _SearchRoleBody extends StatefulWidget {
  final String _language;

  _SearchRoleBody(this._language);

  @override
  State<StatefulWidget> createState() => _SearchBodyView(this._language);
}

class _SearchBodyView extends State<_SearchRoleBody> {
  final String _language;
  TextEditingController _searchText;

  _SearchBodyView(this._language);

  @override
  void initState() {
    _searchText = TextEditingController();
    RolesProvider provider = RolesProvider();
    if (!provider.isCached()) {
      provider.listRoles(0, 1000);
    }

    _rows = <DataRow>[];
    super.initState();

    _getDataRows();
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
                    getText("search_role", this._language),
                    style: Theme.of(context).primaryTextTheme.bodyText1,
                  ),
                ),
                Container(
                  child: TextField(
                    style: Theme.of(context).primaryTextTheme.bodyText2,
                    controller: _searchText,
                  ),
                ),
                Center(
                  child: Container(
                      padding: EdgeInsets.all(5.0),
                      child: OutlineButton(
                        textColor: Colors.white,
                        onPressed: () {
                          filterRows();
                        },
                        child: Text(
                          getText("search", this._language),
                          style: Theme.of(context).primaryTextTheme.button,
                        ),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(30.0),
                        ),
                      )),
                ),
                Center(
                  child: Container(
                      child: DataTable(
                    columns: [
                      DataColumn(
                        label: Text(getText("id", this._language)),
                      ),
                      DataColumn(
                          label: Text(getText("description", this._language))),
                    ],
                    rows: _rows,
                  )),
                ),
              ])),
    );
  }

  List<DataRow> _rows;
  _getDataRows() async {
    List<DataRow> tempRows = <DataRow>[];

    var provider = RolesProvider();
    ListRolesResponse listRolesResponse = await provider.listRoles(0, 1000);

    if (listRolesResponse.status != "Success") {
      print("The response was not successful!");
      return tempRows;
    }

    if (listRolesResponse.roles == null ||
        listRolesResponse.roles.length == 0) {
      print("The roles have not been retrieve!");
      return tempRows;
    }

    List<Role> roles = listRolesResponse.roles;

    for (var i = 0; i < roles.length; i++) {
      var role = roles[i];
      var dataRow = DataRow(
        onSelectChanged: (value) => rowSelected(role, value),
        cells: <DataCell>[
          DataCell(Text(role.id)),
          DataCell(Text(role.description)),
        ],
      );
      tempRows.add(dataRow);
    }

    setState(() {
      this._rows = tempRows;
    });
    return _rows;
  }

  filterRows() async {
    var searchText = _searchText.text;
    if (searchText == null || searchText.isEmpty) {
      _getDataRows();
      return;
    }

    RolesProvider rolesProvider = RolesProvider();
    List<Role> roles = rolesProvider.filterByDescription(searchText);
    List<DataRow> tempRows = <DataRow>[];

    for (var i = 0; i < roles.length; i++) {
      var role = roles[i];
      var dataRow = DataRow(
        onSelectChanged: (value) => rowSelected(role, value),
        cells: <DataCell>[
          DataCell(Text(role.id)),
          DataCell(Text(role.description)),
        ],
      );
      tempRows.add(dataRow);
    }

    setState(() {
      this._rows = tempRows;
    });
  }

  rowSelected(Role role, bool selected) {
    if (selected) {
      Navigator.pushReplacement(
          context,
          MaterialPageRoute(
            builder: (context) => RolesView.withRole(this._language, role),
          ));
    }
  }
}

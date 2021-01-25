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

import 'package:authfe/model/permissionmodel.dart';
import 'package:flutter/material.dart';
import '../i18n/language.dart';
import '../appbar/menudrawer.dart';
import 'permissions.dart';

class SearchPermissions extends StatefulWidget {
  final String _language;

  SearchPermissions(this._language);

  @override
  State<StatefulWidget> createState() =>
      _SearchPermissionsState(this._language);
}

class _SearchPermissionsState extends State<SearchPermissions> {
  final String _language;

  _SearchPermissionsState(this._language);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(getText("title", this._language)),
      ),
      body: SingleChildScrollView(
        child: _SearchPermissionBody(this._language),
      ),
      drawer: DistAuthDrawer(this._language),
    );
  }
}

class _SearchPermissionBody extends StatefulWidget {
  final String _language;

  _SearchPermissionBody(this._language);

  @override
  State<StatefulWidget> createState() => _SearchBodyView(this._language);
}

class _SearchBodyView extends State<_SearchPermissionBody> {
  final String _language;
  TextEditingController _searchText;

  _SearchBodyView(this._language) {
    _searchText = TextEditingController();
  }

  @override
  void initState() {
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
                    getText("search_perm", this._language),
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
                    child: OutlinedButton(
                      style: ButtonStyle(
                        foregroundColor:
                            MaterialStateProperty.all<Color>(Colors.white),
                      ),
                      onPressed: () {
                        _filterRows();
                      },
                      child: Text(
                        getText("search", this._language),
                        style: Theme.of(context).primaryTextTheme.button,
                      ),
                    ),
                  ),
                ),
                Center(
                  child: Container(
                    child: DataTable(
                      columns: [
                        DataColumn(
                          label: Text(getText("description", this._language)),
                        ),
                        DataColumn(
                            label: Text(getText("permission", this._language))),
                      ],
                      rows: rows,
                    ),
                  ),
                ),
              ])),
    );
  }

  List<DataRow> rows = [];
  List<DataRow> fullList = [];

  _filterRows() async {
    if (_searchText.text.length == 0) {
      rows = fullList;
      setState(() {
        rows = fullList;
      });
    }

    List<DataRow> tmpList = [];
    for (int i = 0; i < fullList.length; i++) {
      if (fullList[i].cells[0].child is Text) {
        var textField = fullList[i].cells[0].child;

        if (textField.toString().indexOf(_searchText.text) >= 0) {
          tmpList.add(fullList[i]);
          continue;
        }
      }

      if (fullList[i].cells[1].child is Text) {
        var textField = fullList[i].cells[1].child;

        if (textField.toString().indexOf(_searchText.text) >= 0) {
          tmpList.add(fullList[i]);
        }
      }
    }

    if (tmpList.length > 0) {
      setState(() {
        rows = tmpList;
      });
    }
  }

  _getDataRows() async {
    List<DataRow> localRows = [];
    PermissionProvider permProvider = PermissionProvider();
    ListPermissionResponse permResp =
        await permProvider.listPermissions(0, 1000);

    if (permResp.permissions == null) {
      log('The permissions have not been retrieved yet!');
      return;
    }

    if (permResp.permissions.length == 0) {
      return;
    }

    for (int i = 0; i < permResp.permissions.length; i++) {
      var element = permResp.permissions[i];
      var drow = DataRow(
        onSelectChanged: (value) => rowSelected(element.id),
        cells: <DataCell>[
          DataCell(Text(element.description)),
          DataCell(Text(element.permission)),
        ],
      );
      localRows.add(drow);
    }

    fullList = localRows;
    setState(() {
      rows = localRows;
    });
  }

  rowSelected(String selected) {
    PermissionProvider permissions = PermissionProvider();
    Permission perm = permissions.findPermissionById(selected);
    if (perm != null) {
      Navigator.pushReplacement(
          context,
          MaterialPageRoute(
            builder: (context) => PermissionsView.forEditing(_language, perm),
          ));
    }
  }
}

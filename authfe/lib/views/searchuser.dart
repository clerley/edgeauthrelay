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

import 'package:flutter/material.dart';
import '../i18n/language.dart';
import '../appbar/menudrawer.dart';

class SearchUsers extends StatefulWidget {
  final String _language;

  SearchUsers(this._language);

  @override
  State<StatefulWidget> createState() =>
      _SearchUsersState(this._language);
}

class _SearchUsersState extends State<SearchUsers> {
  final String _language;

  _SearchUsersState(this._language);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(getText("title", this._language)),
      ),
      body: SingleChildScrollView(
        child: _SearchUserBody(this._language),
      ),
      drawer: DistAuthDrawer(this._language),
    );
  }
}

class _SearchUserBody extends StatefulWidget {
  final String _language;

  _SearchUserBody(this._language);

  @override
  State<StatefulWidget> createState() => _SearchBodyView(this._language);
}

class _SearchBodyView extends State<_SearchUserBody> {
  final String _language;

  _SearchBodyView(this._language);

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
                    getText("search_user", this._language),
                    style: Theme.of(context).primaryTextTheme.bodyText1,
                  ),
                ),
                Container(
                  child: TextField(
                      style: Theme.of(context).primaryTextTheme.bodyText2),
                ),
                Center(
                  child: Container(
                      padding: EdgeInsets.all(5.0),
                      child: OutlineButton(
                        textColor: Colors.white,
                        onPressed: () {},
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
                    child: DataTable(columns: [
                      DataColumn(
                        label: Text(getText("username", this._language)),
                      ),
                      DataColumn(
                          label: Text(getText("name", this._language))),
                    ], 
                    rows: _getDataRows(),
                    ),
                  ),
                ),
              ])),
    );
  }

  _getDataRows() {
    var drs = [ 
                      DataRow(
                        cells: <DataCell>[
                          DataCell(Text("...")),
                          DataCell(Text('...')),
                        ],
                      )
                      
    ];

    return drs;
  }
}

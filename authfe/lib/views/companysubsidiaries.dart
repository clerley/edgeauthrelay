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
import 'package:flutter/material.dart';

import 'mainmenu.dart';

class CompanySubsidiariesView extends StatefulWidget {
  final String _language;

  CompanySubsidiariesView(this._language);

  @override
  State<StatefulWidget> createState() =>
      _CompanySubsidiariesViewState(this._language);
}

class _CompanySubsidiariesViewState extends State<CompanySubsidiariesView> {
  final String _language;

  _CompanySubsidiariesViewState(this._language);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: Text(getText("title", this._language)),
        ),
        body: SingleChildScrollView(
          child: _CompanySubsidiariesBody(this._language),
        ),
        drawer: DistAuthDrawer(this._language),
        floatingActionButton: FloatingActionButton(
          backgroundColor: Colors.red,
          onPressed: () => print("Test"),
          child: Icon(Icons.add),
        ));
  }
}

class _CompanySubsidiariesBody extends StatefulWidget {
  final String _language;

  _CompanySubsidiariesBody(this._language);

  @override
  State<StatefulWidget> createState() =>
      _CompanySubsidiariesBodyState(this._language);
}

class _CompanySubsidiariesBodyState extends State<_CompanySubsidiariesBody> {
  final String _language;
  List<DataRow> _rows = [];

  _CompanySubsidiariesBodyState(this._language);

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
            Row(
              children: [
                Expanded(
                  child: Text(
                    getText("search", this._language),
                    style: Theme.of(context).primaryTextTheme.button,
                  ),
                ),
                OutlineButton(
                    textColor: Colors.white,
                    child: Text(
                      getText("search", this._language),
                      style: Theme.of(context).primaryTextTheme.button,
                    ),
                    onPressed: () {},
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(30.0),
                    )),
              ],
            ),
            TextField(
              style: Theme.of(context).primaryTextTheme.bodyText2,
            ),
            DataTable(
              columns: [
                DataColumn(
                  label: Text(getText("uniqueID", this._language)),
                ),
                DataColumn(
                  label: Text(getText("name", this._language)),
                ),
                DataColumn(
                  label: Text(getText("address", this._language)),
                ),
              ],
              rows: _rows,
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                OutlineButton(
                    textColor: Colors.white,
                    child: Text(
                      getText("edit", this._language),
                      style: Theme.of(context).primaryTextTheme.button,
                    ),
                    onPressed: () {},
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(30.0),
                    )),
                OutlineButton(
                    textColor: Colors.white,
                    child: Text(
                      getText("cancel", this._language),
                      style: Theme.of(context).primaryTextTheme.button,
                    ),
                    onPressed: () {
                      Navigator.pushReplacement(
                          context,
                          MaterialPageRoute(
                              builder: (context) => MainMenu(this._language)));
                    },
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(30.0),
                    )),
              ],
            )
          ],
        ),
      ),
    );
  }
}

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
import '../appbar/menudrawer.dart';
import '../i18n/language.dart';

class PermissionsView extends StatefulWidget {

  final String _language;

  PermissionsView(this._language);


  @override
  State<StatefulWidget> createState() => _PermissionsState(this._language);

}

class _PermissionsState extends State<PermissionsView> {

  final String _language;

  _PermissionsState(this._language);

  @override
  Widget build(BuildContext context) {
    
     return Scaffold(
      appBar: AppBar(
        title: Text(getText("title", this._language)),
      ),
      body: SingleChildScrollView(child: _PermissionBody(this._language),)
      ,
      drawer: DistAuthDrawer(this._language), 
      );
 
  }
}

class _PermissionBody extends StatefulWidget {

  final String _language;

  _PermissionBody(this._language);

  @override
  State<StatefulWidget> createState() => _PermissionBodyState(this._language);

}

class _PermissionBodyState extends State<_PermissionBody> {

  final String _language;

  _PermissionBodyState(this._language);

  @override
  Widget build(BuildContext context) {
     return  Center(
              child: Container(
                  child: Column(children: <Widget>[

                    Text(getText("permissions", this._language)),

                  ],
                ),
              ),
            );
  }

}
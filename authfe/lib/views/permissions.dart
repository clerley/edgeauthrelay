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
     return Center(
      child: Container(
        margin: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 0.0),
        padding:EdgeInsets.all(10.0),
        width: 900.0,
        decoration: BoxDecoration(
          color: Theme.of(context).backgroundColor,
          borderRadius: BorderRadius.all(Radius.circular(10.0))
        ),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.start,
          crossAxisAlignment: CrossAxisAlignment.start,
          children: <Widget>[
            Container(
              padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
              child: Text(getText("permissions", this._language), 
                      style: Theme.of(context).primaryTextTheme.bodyText1,),
            ),

            Container(
              child: Text(getText("permission", this._language)),
            ),

            Container(
              child: TextField(
                  style: Theme.of(context).primaryTextTheme.bodyText2),
            ),

            Container(
              padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 0.0),
              child: Text(getText("description", this._language),),
            ),

            Container(
              child: TextField(
                  style: Theme.of(context).primaryTextTheme.bodyText2),
            ),

            Center(
                child: Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: <Widget>[
                    Container(
                        padding: EdgeInsets.all(5.0),
                        child: OutlineButton(
                          textColor: Colors.white,
                          child: Text(getText("save", this._language), style: Theme.of(context).primaryTextTheme.button,),
                          onPressed: () { 
                            
                          },
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(30.0),)
                        ),
                    ),
                    Container(
                      padding: EdgeInsets.all(5.0),
                      child: OutlineButton(
                        textColor: Colors.white,
                        onPressed: () {
                          
                        },
                        child: Text(getText("cancel", this._language), style: Theme.of(context).primaryTextTheme.button,),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(30.0),
                        ),
                      )
                    ),
                  ],
                ),
              ),


          ]
        ),
      ),
    );
          
  }

}
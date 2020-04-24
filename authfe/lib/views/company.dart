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

class CompanyWidget extends StatefulWidget {

  final String _language;

  CompanyWidget(this._language);

  @override
  _CompanyWidgetState createState() => _CompanyWidgetState(this._language);
}

class _CompanyWidgetState extends State<CompanyWidget> {

  String _language;
  String _title;
  _CompanyWidgetState(this._language) {
    this._title = getText("title", this._language);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(this._title),
      ),
      body: SingleChildScrollView(child: CompanyBody(this.widget._language),) 
      );
  }
}


class CompanyBody extends StatefulWidget {
  
  final String _language;

  CompanyBody(this._language);

  @override
  State<StatefulWidget> createState() => _CompanyState(_language);

}

class _CompanyState extends State<CompanyBody> {
  String _language;
  String _company;
  String _addressText;
  String _cityText;
  String _stateText;
  String _zipText;
  String _nameText;
  String _isLocationText;
  String _remotelyManagedText;
  bool _remoteAuth;
  bool _isLocation;
  

  _CompanyState(this._language) {
    this._company = getText("company", _language);
    this._cityText = getText("city", _language);
    this._stateText = getText("state", _language);
    this._zipText = getText("zip", _language);
    this._nameText = getText("name", _language);
    this._addressText = getText("address", _language);
    this._isLocationText = getText("isLocation", _language);
    this._remotelyManagedText = getText("remotelyManaged", _language);
    this._remoteAuth = false;
    this._isLocation = false;
  }

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
              child: Text(this._company, 
                    style: TextStyle(color: Colors.white, fontSize: 22.0),),
            ),
            Container(
              child: Text(this._nameText,
                    style: TextStyle(color: Colors.white, fontSize: 22.0),),
            ),
            Container(
              padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
              child: TextField(
                decoration: InputDecoration(
                  border: OutlineInputBorder(),
                  fillColor: Colors.white,
                  filled: true,
                ),
              ),
            ),
            Container(
              child: Text(this._addressText, 
                            style: TextStyle(color: Colors.white, fontSize: 22.0)),),

            Container(
              padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
              child: TextField(
                decoration: InputDecoration(
                  border: OutlineInputBorder(),
                  fillColor: Colors.white,
                  filled: true,
                ),
              ),
            ),

            Container(
              padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
              child: TextField(
                decoration: InputDecoration(
                  border: OutlineInputBorder(),
                  fillColor: Colors.white,
                  filled: true,
                ),
              ),
            ),
            Row(children: <Widget>[
               Expanded(
                  flex: 24,
                  child: Container(
                      child: Text(this._cityText, 
                              style: TextStyle(color: Colors.white, fontSize: 22.0)),),
                ),

               Spacer(), 

               Expanded(
                  flex: 12,
                  child: Container(
                      child: Text(this._stateText, 
                              style: TextStyle(color: Colors.white, fontSize: 22.0)),),
                ),

               Spacer(),

               Expanded(
                  flex: 24,
                  child: Container(
                      child: Text(this._zipText, 
                              style: TextStyle(color: Colors.white, fontSize: 22.0)),),
                ),
               

            ],),

            Row(children: <Widget>[

              Expanded(
                flex: 24,
                  child: Container(
                  padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                  child: TextField(
                    decoration: InputDecoration(
                      border: OutlineInputBorder(),
                      fillColor: Colors.white,
                      filled: true,
                    ),
                  ),
                ),
              ),

              Spacer(),

              Expanded(
                flex: 12,
                  child: Container(
                  padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                  child: TextField(
                    decoration: InputDecoration(
                      border: OutlineInputBorder(),
                      fillColor: Colors.white,
                      filled: true,
                    ),
                  ),
                ),
              ),

              Spacer(), 
              
              Expanded(
                flex: 24,
                child: Container(
                  padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                  child: TextField(
                    decoration: InputDecoration(
                      border: OutlineInputBorder(),
                      fillColor: Colors.white,
                      filled: true,
                    ),
                  ),
                ),
              ),
            ],
            ),

            Container(
              child: Row(
                children: <Widget>[
                  Checkbox(
                    onChanged: (val) {
                      setState(() {
                        this._isLocation = val;
                      });
                    },
                    value: this._isLocation,
                  ),
                  Text(this._isLocationText,
                          style: TextStyle(fontSize: 22.0, color: Colors.white)),
                  Checkbox(
                    onChanged: (val) {
                      setState(() {
                        this._remoteAuth = val;
                      });
                    },
                    value: this._remoteAuth,
                  ),
                  Text(this._remotelyManagedText,
                          style: TextStyle(fontSize: 22.0, color: Colors.white)),

                ],
              ),
            ),

            Container(
              margin: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 0.0),
              padding:EdgeInsets.all(10.0),
              width: 900.0,
              decoration: BoxDecoration(
                            color: Colors.blueAccent,
                            borderRadius: BorderRadius.all(Radius.circular(10.0)),
              ),
              child: Column(
                children: <Widget>[
                  Row(
                    children: <Widget>[
                      Text(getText("settings", this._language),
                          style: TextStyle(color: Colors.white, fontSize: 22.0))
                    ],
                  ),
                  Row(children: <Widget>[
                      Expanded(
                        flex: 24,
                        child: Text(getText("jwtDuration", this._language), 
                                style: TextStyle(color: Colors.white, fontSize: 22.0)),
                      ),
                      Spacer(),
                      Expanded(
                        flex: 24,
                        child: Text(getText("passwordExp", this._language),
                                style: TextStyle(color: Colors.white, fontSize: 22.0)),
                      ),
                      Spacer(),
                      Expanded(
                        flex: 24,
                        child: Text(getText("passwordUnit", this._language),
                                style: TextStyle(color: Colors.white, fontSize: 22.0)),
                      )
                  ],), 
                  Row(children: <Widget>[
                      Expanded(
                        flex: 24,
                        child: TextField(
                                decoration: InputDecoration(
                                              border: OutlineInputBorder(),
                                              fillColor: Colors.white,
                                              filled: true,
                                          ),
                                )
                      ),
                      Spacer(),
                      Expanded(
                        flex: 24,
                        child: TextField(
                                decoration: InputDecoration(
                                              border: OutlineInputBorder(),
                                              fillColor: Colors.white,
                                              filled: true,
                                              
                                          ),
                                )
                      ),
                      Spacer(),
                      Expanded(
                        flex: 24,
                        child: TextField(
                                decoration: InputDecoration(
                                              border: OutlineInputBorder(),
                                              fillColor: Colors.white,
                                              filled: true,
                                          ),
                                )
                      )

                  ],)
                ],
              )

            ),
          ],
        ),
      ),
    );
  }

}
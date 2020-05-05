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
  String _unit;
  List<DropdownMenuItem<String>> _unitsMenuItems;

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
    this._unit = "Minute";
    this._unitsMenuItems = _getUnitMenuItems();
  }

  List<DropdownMenuItem<String>> _getUnitMenuItems() {
    List<DropdownMenuItem<String>> lst = new List<DropdownMenuItem<String>>();
    DropdownMenuItem<String> item = new DropdownMenuItem<String>(child: Text("Minute"),value: "Minute");
    lst.add(item);

    item = new DropdownMenuItem<String>(child: Text("Week"), value: "Week");
    lst.add(item);

    return lst;
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
              child: Text(this._company, style: Theme.of(context).primaryTextTheme.bodyText1,),
            ),
            
            Container(
              child: Text(getText("uniqueID", this._language)),
            ),

            Container(
              child: TextField(),
            ),

            Container(
              padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 0.0),
              child: Text(this._nameText,),
            ),

            Container(
              child: TextField(),
            ),
            
            Container(
              padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 0.0),
              child: Text(this._addressText,),
            ),

            Container(
              padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
              child: TextField(),
            ),

            Container(
              padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
              child: TextField(),
            ),
            Row(children: <Widget>[
               Expanded(
                  flex: 24,
                  child: Container(
                      child: Text(this._cityText,),
                  ),
                ),

               Spacer(), 

               Expanded(
                  flex: 12,
                  child: Container(
                      child: Text(this._stateText,),
                  ),
                ),

               Spacer(),

               Expanded(
                  flex: 24,
                  child: Container(
                      child: Text(this._zipText),
                  ),
                ),
               

            ],),

            Row(children: <Widget>[

              Expanded(
                flex: 24,
                  child: Container(
                  padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                  child: TextField(),
                ),
              ),

              Spacer(),

              Expanded(
                flex: 12,
                  child: Container(
                  padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                  child: TextField(),
                ),
              ),

              Spacer(), 
              
              Expanded(
                flex: 24,
                child: Container(
                  padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                  child: TextField(),
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
                  Text(this._isLocationText),
                  Checkbox(
                    onChanged: (val) {
                      setState(() {
                        this._remoteAuth = val;
                      });
                    },
                    value: this._remoteAuth,
                  ),
                  Text(this._remotelyManagedText),

                ],
              ),
            ),

            Container(
              child:Text(getText("authrelay", this._language),
              ),
            ),

            Container(
              child:TextField()
            ,),

            Container(
              margin: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 0.0),
              padding:EdgeInsets.all(10.0),
              width: 900.0,
              decoration: BoxDecoration(
                            color: Theme.of(context).backgroundColor,
                            borderRadius: BorderRadius.all(Radius.circular(10.0)),
                            border: Border.all(color:Theme.of(context).accentColor),
              ),
              child: Column(
                children: <Widget>[
                  Row(children: <Widget>[
                      Expanded(
                        flex: 24,
                        child: Text(getText("jwtDuration", this._language)),
                      ),
                      Spacer(),
                      Expanded(
                        flex: 24,
                        child: Text(getText("passwordExp", this._language)),
                      ),
                      Spacer(),
                      Expanded(
                        flex: 24,
                        child: Text(getText("passwordUnit", this._language)),
                      )
                  ],), 
                  Row(children: <Widget>[
                      Expanded(
                        flex: 24,
                        child: TextField(
                          keyboardType: TextInputType.number,
                          ),
                      ),
                      Spacer(),
                      Expanded(
                        flex: 24,
                        child: TextField(keyboardType: TextInputType.number,),
                      ),
                      Spacer(),
                      Expanded(
                        flex: 24,
                        //child: TextField(keyboardType: TextInputType.number,),
                        child: new DropdownButton<String>(
                              value: _unit,
                              items: _unitsMenuItems,
                              onChanged: (String newValue) {
                                setState(() {
                                  _unit = newValue;
                                }
                              );
                            },
                        ),
                      )

                  ],)
                ],
              )

            ),


            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: <Widget>[
              Container(
                  padding: EdgeInsets.all(5.0),
                  child: OutlineButton(
                  textColor: Colors.white,
                  child: Text(getText("save", this._language), style: Theme.of(context).primaryTextTheme.button,),
                  onPressed: () { 
                    print ("Testing It"); 
                  },
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(30.0),)
                ),
              ),

              Container(
                  padding: EdgeInsets.all(5.0),
                  child: OutlineButton(
                  textColor: Colors.white,
                  child: Text(getText("cancel", this._language), style: Theme.of(context).primaryTextTheme.button,),
                  onPressed: () { 
                    Navigator.pop(context);
                  },
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(30.0),)
                ),
              ),


            ],
            ),

          ],
        ),
      ),
    );
  }

}
import 'package:flutter/material.dart';
import '../i18n/language.dart';

class LoginWidget extends StatefulWidget {

  @override
  _LoginState createState() => _LoginState(LANG_ENGLISH);

}

class _LoginState extends State<LoginWidget> {

  String _language = LANG_ENGLISH;
  String _userNameText;
  String _passwordText; 
  String _appName;
  String _loginText;
  String _uniqueCompanyID;
  String _newCompany;

  _LoginState(this._language) {
    this._userNameText = getText("username", _language);
    this._passwordText = getText("password", _language);
    this._appName = getText("title", _language);
    this._loginText = getText("login", _language);
    this._uniqueCompanyID = getText("uniqueID", _language);
    this._newCompany = getText("newcompany", _language);
  }


  @override
  Widget build(BuildContext context) {

    return 
      Center(
        child: Container(
          padding:EdgeInsets.all(16.0),
          decoration: BoxDecoration(
            color: Colors.green,
            borderRadius: BorderRadius.all(Radius.circular(10.0))
            ),
          width: 400.0,
          height: 505.0,
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: <Widget>[
              Container(
                alignment: Alignment.center,
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
                child: Text(this._appName,
                        style: TextStyle(fontSize: 26, color: Colors.white)),
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
                child: Text(this._uniqueCompanyID,
                  style: TextStyle(fontSize: 22, color:Colors.white)),
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
                padding: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 10.0),
                child: Text(this._userNameText,
                  style: TextStyle(fontSize: 22, color:Colors.white)),
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 20.0),
                child: TextField(
                  decoration: InputDecoration( 
                    border: OutlineInputBorder(),
                    fillColor: Colors.white,
                    filled: true,
                  ),
                ),
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 20.0),
                child: Text(this._passwordText,
                  style: TextStyle(fontSize: 22, color: Colors.white)),
              ),
              TextField(
                obscureText: true,
                decoration: InputDecoration(
                  border: OutlineInputBorder(),
                  fillColor: Colors.white,
                  filled: true,
                ),
              ),
              Divider(
                color: Colors.white,
                thickness: 1.0,
              ),
              Center(
                child: Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: <Widget>[
                    Container(
                        padding: EdgeInsets.all(5.0),
                        child: OutlineButton(
                          textColor: Colors.white,
                          child: Text(this._loginText, style: TextStyle(color: Colors.white, fontSize: 16.0),),
                          onPressed: () { print ("Testing It"); },
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(30.0),)
                        ),
                    ),
                    Container(
                      padding: EdgeInsets.all(5.0),
                      child: OutlineButton(
                        textColor: Colors.white,
                        onPressed: () { print("Testing it again"); },
                        child: Text(this._newCompany, style: TextStyle(color: Colors.white, fontSize: 16.00),),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(30.0),
                        ),
                      )
                    ),
                  ],
                ),
              ),
            ],

          ),
        )
      );
  }

}
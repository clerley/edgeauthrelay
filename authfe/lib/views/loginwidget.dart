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

  _LoginState(this._language) {
    this._userNameText = getText("username", _language);
    this._passwordText = getText("password", _language);
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
          height: 400.0,
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: <Widget>[
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                child: Text(this._userNameText,
                  style: TextStyle(fontSize: 22, color:Colors.white)),
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 20.0),
                child: TextField(
                  decoration: InputDecoration( 
                    border: OutlineInputBorder(),
                  ),
                ),
              ),
              Container(
                padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                child: Text(this._passwordText,
                  style: TextStyle(fontSize: 22, color: Colors.white)),
              ),
              TextField(
                obscureText: true,
                decoration: InputDecoration(
                  border: OutlineInputBorder(),
                ),
              ),
            ],

          ),
        )
      );
  }

}
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
          color: Colors.blue,
          width: 350.0,
          height: 350.0,
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: <Widget>[
              Text(this._userNameText),
              TextField(),
              Text(this._passwordText),
              TextField(),
            ],

          ),
        )
      );
  }

}
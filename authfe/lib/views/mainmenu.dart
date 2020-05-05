import "package:flutter/material.dart";
import "../i18n/language.dart";


class MainMenu extends StatefulWidget {

  String _language;

  MainMenu(this._language);

  @override
  State<StatefulWidget> createState() => _MainMenuState(this._language);

}

class _MainMenuState extends State<MainMenu> {

  String _language;
  String _title;

  _MainMenuState(this._language) {
    this._title = getText("title", this._language);
  }


  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(this._title),
      ),
      body: SingleChildScrollView(child: _MainMenuBody(),) 
      );
  }

}

class _MainMenuBody extends StatelessWidget {


  @override
  Widget build(BuildContext context) {
    return Column();
  }

}
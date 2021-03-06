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

import "package:flutter/material.dart";
import "../i18n/language.dart";
import "../appbar/menudrawer.dart";


class MainMenu extends StatefulWidget {

  final String _language;

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
      body: SingleChildScrollView(child: _MainMenuBody(),
      ) ,
      drawer: DistAuthDrawer(this._language),
      );
  }

}

class _MainMenuBody extends StatelessWidget {


  @override
  Widget build(BuildContext context) {
    return Column(

    );
  }

}
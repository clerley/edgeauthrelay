import 'package:flutter/material.dart';


class CompanyWidget extends StatefulWidget {
  
  final String _language;

  CompanyWidget(this._language);

  @override
  State<StatefulWidget> createState() => _CompanyState(_language);

}

class _CompanyState extends State<CompanyWidget> {

  String _language;

  _CompanyState(this._language);

  @override
  Widget build(BuildContext context) {

    return Center(
      child: Container(
          padding:EdgeInsets.all(16.0),
          decoration: BoxDecoration(
            color: Colors.green,
            borderRadius: BorderRadius.all(Radius.circular(10.0))
          ),
        child: Column(
          children: <Widget>[


          ],
        ),
      ),
    );
  }

}
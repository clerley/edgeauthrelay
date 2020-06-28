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

import 'package:authfe/views/settingsview.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'i18n/language.dart';
import 'model/companymodel.dart';
import 'model/rolesmodel.dart';
import 'views/loginwidget.dart';
import 'model/usermodel.dart';
import 'model/permissionmodel.dart';


void main() => runApp(MultiProvider(
                          providers:[
                            ChangeNotifierProvider<UserProvider>(create: (_) => UserProvider(),),
                            ChangeNotifierProvider<CompanyProvider>(create: (_) => CompanyProvider(),),
                            ChangeNotifierProvider<PermissionProvider>(create: (_) => PermissionProvider(),),
                            ChangeNotifierProvider<RolesProvider>(create: (_) => RolesProvider(),),
                          ],
                          child: MyApp()));

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    var titleText = getText('title', LANG_ENGLISH);
        return MaterialApp(
          title: titleText,
          theme: ThemeData(
            brightness: Brightness.dark,
            appBarTheme: AppBarTheme(
              brightness: Brightness.dark,
              color: Color(0xff07203e),
            ),
            
            primaryColor: Color(0xff7b92ae),
            accentColor: Color(0xff506d90),
            backgroundColor: Color(0xff18365a),
            
            fontFamily: 'Arial',
            primaryTextTheme: TextTheme(
                bodyText1: TextStyle(fontSize: 22.0, color:Colors.white),
                button: TextStyle(fontSize: 16.0, color:Colors.white),
                bodyText2: TextStyle(fontSize: 16.0, color: Colors.yellow),
            ),

            inputDecorationTheme: InputDecorationTheme(
              //fillColor: Color(0xff7b92ae),
              fillColor: Color(0xff222831),
              filled: true,
            ),

            toggleableActiveColor: Color(0xff506d90),

          ),
          home: MyHomePage(title: titleText),
    );
  }
}

class MyHomePage extends StatefulWidget {
  MyHomePage({Key key, this.title}) : super(key: key);
  final String title;

  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
        actions: [IconButton(
          icon: Icon(Icons.settings), 
          onPressed: () {
            Navigator.pushReplacement(
                               context,
                               MaterialPageRoute(
                                 builder: (context) => SettingsView()));
          },
        )],
      ),
      body: SingleChildScrollView(child: LoginWidget(),), 
      );
  }
}

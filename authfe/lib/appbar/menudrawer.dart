import 'package:flutter/material.dart';
import '../i18n/language.dart';

Drawer getDrawer(String language) {

  return Drawer(
    // Add a ListView to the drawer. This ensures the user can scroll
    // through the options in the drawer if there isn't enough vertical
    // space to fit everything.
    child: ListView(
      // Important: Remove any padding from the ListView.
      padding: EdgeInsets.zero,
      children: <Widget>[
        DrawerHeader(
          child: Text(getText("title", language)),
          decoration: BoxDecoration(
            color: Colors.blue,
          ),
        ),
        ListTile(
          title: Text(getText("roles", language)),
          onTap: () {
            // Update the state of the app.
            // ...
          },
        ),
        ListTile(
          title: Text(getText("permissions", language)),
          onTap: () {
            // Update the state of the app.
            // ...
          },
        ),
                ListTile(
          title: Text(getText("users", language)),
          onTap: () {
            // Update the state of the app.
            // ...
          },
        ),
      ],
    ),
  );
}

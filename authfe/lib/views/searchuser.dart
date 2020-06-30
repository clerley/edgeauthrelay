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

import 'package:authfe/model/usermodel.dart';
import 'package:authfe/views/users.dart';
import 'package:flutter/material.dart';
import '../i18n/language.dart';
import '../appbar/menudrawer.dart';

class SearchUsers extends StatefulWidget {
  final String _language;

  SearchUsers(this._language);

  @override
  State<StatefulWidget> createState() => _SearchUsersState(this._language);
}

class _SearchUsersState extends State<SearchUsers> {
  final String _language;

  _SearchUsersState(this._language);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(getText("title", this._language)),
      ),
      body: SingleChildScrollView(
        child: _SearchUserBody(this._language),
      ),
      drawer: DistAuthDrawer(this._language),
    );
  }
}

class _SearchUserBody extends StatefulWidget {
  final String _language;

  _SearchUserBody(this._language);

  @override
  State<StatefulWidget> createState() => _SearchBodyView(this._language);
}

class _SearchBodyView extends State<_SearchUserBody> {
  final String _language;
  List<DataRow> _rows = [];
  TextEditingController _searchText = TextEditingController();

  _SearchBodyView(this._language);

  @override
  void initState() {
    super.initState();

    _getDataRows();
  }

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Container(
          margin: EdgeInsets.fromLTRB(0.0, 10.0, 0.0, 0.0),
          padding: EdgeInsets.all(10.0),
          width: 900.0,
          decoration: BoxDecoration(
              color: Theme.of(context).backgroundColor,
              borderRadius: BorderRadius.all(Radius.circular(10.0))),
          child: Column(
              mainAxisAlignment: MainAxisAlignment.start,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: <Widget>[
                Container(
                  padding: EdgeInsets.fromLTRB(0.0, 0.0, 0.0, 10.0),
                  child: Text(
                    getText("search_user", this._language),
                    style: Theme.of(context).primaryTextTheme.bodyText1,
                  ),
                ),
                Container(
                  child: TextField(
                    style: Theme.of(context).primaryTextTheme.bodyText2,
                    controller: _searchText,
                  ),
                ),
                Center(
                  child: Container(
                      padding: EdgeInsets.all(5.0),
                      child: OutlineButton(
                        textColor: Colors.white,
                        onPressed: () {
                          _searchUser();
                        },
                        child: Text(
                          getText("search", this._language),
                          style: Theme.of(context).primaryTextTheme.button,
                        ),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(30.0),
                        ),
                      )),
                ),
                Center(
                  child: Container(
                    child: DataTable(
                      columns: [
                        DataColumn(
                          label: Text(getText("username", this._language)),
                        ),
                        DataColumn(
                            label: Text(getText("name", this._language))),
                      ],
                      rows: _rows,
                    ),
                  ),
                ),
              ])),
    );
  }

  handleSelected(User user) {
    UserProvider().edittingUser = user;
    Navigator.pushReplacement(
      context,
      MaterialPageRoute(
        builder: (context) => UsersView(this._language),
      ),
    );
  }

  _searchUser() {
    String searchText = _searchText.text;
    searchText = searchText.toLowerCase();
    List<User> filtered = [];
    List<User> userList = UserProvider().getCachedListOfUsers();
    if (userList != null) {
      for (var i = 0; i < userList.length; i++) {
        var user = userList[i];
        if (user.name != null &&
            user.name.toLowerCase().indexOf(searchText) >= 0) {
          filtered.add(user);
        } else if (user.username != null &&
            user.username.toLowerCase().indexOf(searchText) >= 0) {
          filtered.add(user);
        }
      }
    }

    _populateUserList(filtered);
  }

  _populateUserList(List<User> usersList) {
    List<DataRow> tempRows = [];
    if (usersList != null) {
      for (var i = 0; i < usersList.length; i++) {
        User user = usersList[i];
        DataRow row = DataRow(
          cells: [],
          onSelectChanged: (value) => handleSelected(user),
        );
        DataCell cell = DataCell(
          Text(user.username),
        );
        row.cells.add(cell);

        cell = DataCell(
          Text(user.name),
        );
        row.cells.add(cell);
        tempRows.add(row);
      }

      setState(() {
        this._rows = tempRows;
      });
    }
  }

  _getDataRows() async {
    debugPrint(
        'Initiating getDataRows. For some reason this method does not seem to return');
    UserProvider usersProvider = UserProvider();
    //For now we will just keep at 1000. Eventually we will have to
    //implement pagination but. I want to keep it simple.
    List<User> usersList = usersProvider.getCachedListOfUsers();
    if (usersList != null && usersList.length == 0) {
      UsersList rsp = await usersProvider.listUsers(0, 1000);
      if (rsp.status == "Success") {
        debugPrint('Users from 0 through 1000 have been retrieved!');
        usersList = usersProvider.getCachedListOfUsers();
      } else {
        debugPrint('No users have been returned, the status is ${rsp.status}');
        return;
      }
    }

    _populateUserList(usersList);
  }
}

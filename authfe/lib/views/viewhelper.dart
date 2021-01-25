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

import 'package:authfe/i18n/language.dart';
import 'package:authfe/model/companymodel.dart';
import 'package:flutter/material.dart';
import 'package:progress_dialog/progress_dialog.dart';

class DialogHelper {
  static final DialogHelper _theInstance = DialogHelper._internalConstructor();

  DialogHelper._internalConstructor();

  factory DialogHelper() {
    return _theInstance;
  }

  ProgressDialog createProgressDialog(String msg, var context) {
    var pr = new ProgressDialog(context);
    pr.style(
        message: msg,
        borderRadius: 10.0,
        backgroundColor: Colors.white,
        progressWidget: CircularProgressIndicator(),
        elevation: 10.0,
        insetAnimCurve: Curves.easeInOut,
        progress: 0.0,
        maxProgress: 100.0,
        progressTextStyle: TextStyle(
            color: Colors.black, fontSize: 13.0, fontWeight: FontWeight.w400),
        messageTextStyle: TextStyle(
            color: Colors.black, fontSize: 19.0, fontWeight: FontWeight.w600));
    return pr;
  }

  void showMessageDialog(String msg, BuildContext context, String language) {
    // flutter defined function
    showDialog(
      context: context,
      builder: (BuildContext context) {
        // return object of type Dialog
        return AlertDialog(
          title: new Text(getText("warning", language)),
          content: new Text(msg),
          actions: <Widget>[
            OutlinedButton(
              style: ButtonStyle(
                foregroundColor: MaterialStateProperty.all<Color>(Colors.white),
              ),
              child: Text(
                getText("close", language),
                style: Theme.of(context).primaryTextTheme.button,
              ),
              onPressed: () {
                Navigator.of(context).pop();
              },
            ),
          ],
        );
      },
    );
  }

  void showCompanyInfo(Company company, BuildContext context, String language) {
    // flutter defined function
    showDialog(
      context: context,
      builder: (BuildContext context) {
        // return object of type Dialog
        return AlertDialog(
          title: Center(
            child: new Text(
              getText("company", language),
              style: TextStyle(color: Colors.red),
            ),
          ),
          content: Container(
            height: 300,
            width: 400,
            child: Column(
              children: [
                Padding(
                  padding: const EdgeInsets.fromLTRB(8.0, 5.0, 8.0, 2.5),
                  child: Row(
                    children: [
                      new Text(
                        getText('uniqueID', language),
                        style: TextStyle(color: Colors.red),
                      ),
                    ],
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.fromLTRB(8.0, 2.5, 8.0, 2.5),
                  child: Row(
                    children: [
                      Text(company.companyID),
                      Text(' - '),
                      Text(company.uniqueID),
                    ],
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.fromLTRB(8.0, 5.0, 8.0, 2.5),
                  child: Row(
                    children: [
                      Text(
                        getText('name', language),
                        style: TextStyle(color: Colors.red),
                      ),
                    ],
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.fromLTRB(8.0, 2.5, 8.0, 2.5),
                  child: Row(
                    children: [
                      Text(company.name),
                    ],
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.fromLTRB(8.0, 5.0, 8.0, 2.5),
                  child: Row(
                    children: [
                      Text(
                        getText('address', language),
                        style: TextStyle(color: Colors.red),
                      ),
                    ],
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.fromLTRB(8.0, 2.5, 8.0, 2.5),
                  child: Row(
                    children: [
                      Text(company.getFullAddress()),
                    ],
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.fromLTRB(8.0, 10.0, 8.0, 2.5),
                  child: Row(
                    children: [
                      Text(company.regisCode, style: TextStyle(fontSize: 22)),
                    ],
                  ),
                ),
              ],
            ),
          ),
          actions: <Widget>[
            OutlinedButton(
              style: ButtonStyle(
                foregroundColor: MaterialStateProperty.all<Color>(Colors.white),
              ),
              child: Icon(Icons.lock_open),
              onPressed: () {
                (Company company) async {
                  var hd = createProgressDialog(
                      getText('please-wait', language), context);
                  await hd.show();
                  CompanyProvider companyProvider = CompanyProvider();
                  CompanyRegistrationResponse regResp =
                      await companyProvider.enableRegistration(company);
                  await hd.hide();
                  if (regResp.status == "Success") {
                    company.regisCode = regResp.regisCode;
                    debugPrint(
                        'The value of the registration code is:[${company.regisCode}]');
                    Navigator.of(context).pop();
                  }
                }(company);
              },
            ),
            OutlinedButton(
              style: ButtonStyle(
                foregroundColor: MaterialStateProperty.all<Color>(Colors.white),
              ),
              child: Text(
                getText("close", language),
                style: Theme.of(context).primaryTextTheme.button,
              ),
              onPressed: () {
                Navigator.of(context).pop();
              },
            ),
          ],
        );
      },
    );
  }
}

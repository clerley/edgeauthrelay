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

import 'dart:core';

const LANG_ENGLISH = "En";
const LANG_SPANISH = "Sp";

Map<String, Map<String, String>> languages = {
    LANG_ENGLISH: {"hello"           : "Hello",
                   "title"           : "Novare Security Systems",
                   "username"        : "Username",
                   "password"        : "Password",
                   "login"           : "Login",
                   "uniqueID"        : "Company ID",
                   "create"          : "Create",
                   "newcompany"      : "New Company",
                   "company"         : "Company",
                   "address"         : "Address",
                   "zip"             : "Zip",
                   "city"            : "City",
                   "state"           : "State",
                   "email"           : "Email",
                   "name"            : "Name",
                   "isLocation"      : "Is Location",
                   "remotelyManaged" : "Remotely Managed",
                   "authrelay"       : "Authorization Relay",
                   "settings"        : "Settings",
                   "jwtDuration"     : "JWT Duration",
                   "passwordExp"     : "Password Expiration",
                   "passwordUnit"    : "Password Unit",
                   "save"            : "Save",
                   "cancel"          : "Cancel",
                   "roles"           : "Roles",
                   "permissions"     : "Permissions",
                   "users"           : "Users",
                   "companies"       : "Companies", 
                   "logout"          : "Logout",    
                   "permission"      : "Permission",         
                   "description"     : "Description",  
                   "search"          : "Search",
                   "search_perm"     : "Search Permissions",
                   "search_role"     : "Search Roles",
                   "add"             : "Add",
                   }
  ,
    LANG_SPANISH: {"hello"           : "Hola!",
                   "title"           : "Novare Security Systems",
                   "username"        : "Usuario",
                   "password"        : "Sena",
                   "login"           : "Login",
                   "uniqueID"        : "ID Compania",
                   "create"          : "Criar",
                   "newcompany"      : "Compania Nueva",
                   "company"         : "Compania",
                   "address"         : "Address",
                   "zip"             : "Zip",
                   "city"            : "City",
                   "state"           : "State",
                   "email"           : "Email",
                   "name"            : "Nombre",
                   "isLocation"      : "isLocation",
                   "remotelyManaged" : "Remotely Managed",
                   "authrelay"       : "Authorization Relay",
                   "settings"        : "Configuraciones",
                   "jwtDuration"     : "JWT Duration",
                   "passwordExp"     : "Password Expiration",
                   "passwordUnit"    : "Password Unit",
                   "save"            : "Save",
                   "cancel"          : "Cancel",                  
                   "roles"           : "Roles",
                   "permissions"     : "Permissions",
                   "users"           : "Users",
                   "companies"       : "Companies",
                   "logout"          : "Logout",
                   "permission"      : "Permission", 
                   "description"     : "Description",
                   "search"          : "Search",
                   "search_perm"     : "Search Permissions",
                   "search_role"     : "Search Roles",
                   "add"             : "Add",
                   } 
};



String getText(String text, String language) {

  if(!languages.containsKey(language)) {
    return "Unknown";
  }


  var lan = languages[language];
  if(lan.containsKey(text)) {
    return lan[text];
  }

  return "Unknown";
}


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
                   "search_user"     : "Search Users",
                   "isthing"         : "Is Thing?",
                   "secret"          : "Secret",
                   "please_wait"     : "Please wait, processing request",
                   "confirmPassword" : "Confirm Password",
                   "warning"         : "Warning",
                   "close"           : "Close",
                   "error_add_cmp"   : "Error adding company. Not additional information available!",
                   "new_cmp_create"  : "A new company has been created!",
                   "user_not_logged" : "User was not properly logged in! Verify the company ID, username and password",
                   "perm_ins_success": "Permission was successfully inserted",
                   "perm_ins_error"  : "Error inserting the permission", 
                   "perm_upd_success": "Permission was successfully updated",
                   "perm_upd_error"  : "Error updating the permission",
                   "add_role_failed" : "Unable to add the role",
                   "update_role_failed": "Unable to update the role",
                   "update_role_success": "Role update was successful!",
                   "add_role_success": "Role was successfully added",
                   "id"              : "ID",
                   "edit"            : "Edit",
                   "error_upt_cmp"   : "Error updating company",
                   "cmp_updt_failed" : "The company update failed",
                   "company_updated" : "The company has been updated",
                   "not_editable"    : "An issue occurred while attempting to edit the company",
                   "list_mgr_company": "Subsidiaries",
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
                   "search_user"     : "Search Users",
                   "isthing"         : "Is Thing?",
                   "secret"          : "Secret",
                   "please_wait"     : "Please wait, processing request",
                   "confirmPassword" : "Confirm Password",
                   "warning"         : "Warning",
                   "close"           : "Close",
                   "error_add_cmp"   : "Error adding company. Not additional information available!",
                   "new_cmp_create"  : "A new company has been created!",
                   "user_not_logged" : "User was not properly logged in! Verify the company ID, username and password",
                   "perm_ins_success": "Permission was successfully inserted",
                   "perm_ins_error"  : "Error inserting the permission", 
                   "perm_upd_success": "Permission was successfully updated",
                   "perm_upd_error"  : "Error updating the permission",
                   "add_role_failed" : "Unable to add the role",
                   "update_role_failed": "Unable to update the role",
                   "update_role_success": "Role update was successful!",
                   "add_role_success": "Role was successfully added",
                   "id"              : "ID",
                   "edit"            : "Edit",
                   "error_upt_cmp"   : "Error updating company",
                   "cmp_updt_failed" : "The company update failed",
                   "company_updated" : "The company has been updated",
                   "not_editable"    : "The company is no editable",
                   "list_mgr_company": "Subsidiaries",
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



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


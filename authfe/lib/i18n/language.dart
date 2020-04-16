
import 'dart:core';

const LANG_ENGLISH = "En";
const LANG_SPANISH = "Sp";

Map<String, Map<String, String>> languages = {
    LANG_ENGLISH: {"hello"    : "Hello",
                   "title"    : "Novare Security Systems",
                   "username" : "Username",
                   "password" : "Password"
                   }
  ,
    LANG_SPANISH: {"hello"    : "Hola!",
                   "title"    : "Novare Security Systems",
                   "username" : "Usuario",
                   "password" : "Sena"}
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


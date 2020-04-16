

import 'package:flutter_test/flutter_test.dart';
import '../../lib/i18n/language.dart';

void main() {

  test('Test Language', () async {

    var rsp = getText("title", LANG_ENGLISH);
    expect("Novare Security Systems", rsp);

    rsp = getText("title", LANG_SPANISH);
    expect("Novare Security Systems", rsp);

  });


}
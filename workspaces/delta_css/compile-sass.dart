import 'dart:io';
//import 'package:sass/sass.dart' as sass;

void main(List<String> arguments) async {
  // clean before creating our dist
  //Directory("dist").delete(recursive: true);

  var srcDir = Directory("src");

  await for (var entity in srcDir.list(recursive: true)) {
    if (entity is File) {
      entity.exists();
    }
  }
  // generate our folder because sass library won't do it by itself
  //Directory("dist").create(recursive: false);

  //var result = sass.compileToResult(arguments[0]).css;
  //File(arguments[1]).writeAsStringSync(result);
}

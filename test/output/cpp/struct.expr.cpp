#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/defer.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/file.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/std.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/var.cpp"
#include <string>
defer onExitFuncs;
std::map<std::string, var> structMap;
var genStruct(std::string structName) {
  var structValue = structMap[structName];
  return structValue;
}

int main() {

  var another = {};
  another["woah"] = 0;

  var yeah = {};
  yeah["woah"] = 8;

  var Thing = {};
  Thing["fieldA"] = 0;
  Thing["stringField"] = "";
  Thing["false_field"] = false;
  Thing["anotherFielderino"] = 0.000000;
  var thing_SZblKKEHly = {};
  thing_SZblKKEHly["woah"] = 0;
  Thing["thing"] = thing_SZblKKEHly;

  var something = {};
  something["fieldA"] = 666;
  something["stringField"] = "";
  something["false_field"] = false;
  something["anotherFielderino"] = 0.000000;
  var thing_OovnTVYvuv = {};
  thing_OovnTVYvuv["woah"] = 0;
  something["thing"] = thing_OovnTVYvuv;
}

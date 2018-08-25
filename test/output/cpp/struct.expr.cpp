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
  yeah["woah"] = 0;

  var Thing = {};
  Thing["fieldA"] = 0;
  Thing["stringField"] = "";
  Thing["false_field"] = false;
  Thing["anotherFielderino"] = 0.000000;
  var thing_iBhKOFOZSD = {};
  thing_iBhKOFOZSD["woah"] = 0;
  Thing["thing"] = thing_iBhKOFOZSD;

  var something = {};
  something["fieldA"] = 0;
  something["stringField"] = "";
  something["false_field"] = false;
  something["anotherFielderino"] = 0.000000;
  var thing_QjrDsiwBII = {};
  thing_QjrDsiwBII["woah"] = 0;
  something["thing"] = thing_QjrDsiwBII;

  Println("something", something);

  var something2 = {};
  something2["fieldA"] = 912559;
  something2["stringField"] = "";
  something2["false_field"] = false;
  something2["anotherFielderino"] = 0.000000;
  var thing_phPnKXCATQ = {};
  thing_phPnKXCATQ["woah"] = 0;
  something2["thing"] = thing_phPnKXCATQ;

  Println("something2", something2);

  var something3 = {};
  something3["fieldA"] = 0;
  something3["stringField"] = "chyah brah";
  something3["false_field"] = false;
  something3["anotherFielderino"] = 0.000000;
  var thing_MbtjuwJCci = {};
  thing_MbtjuwJCci["woah"] = 0;
  something3["thing"] = thing_MbtjuwJCci;

  Println("something3", something3);
}

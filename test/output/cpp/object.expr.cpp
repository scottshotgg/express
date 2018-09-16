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

  var obj = {};
  obj["something"] = "here";
  var hey_HbOdYZPuda = {};
  hey_HbOdYZPuda["me"] = true;
  hey_HbOdYZPuda["anIntVariable"] = 69;
  obj["hey"] = hey_HbOdYZPuda;

  var objs[] = {};
  {
    var _wliELUSHTE = {};
    _wliELUSHTE["another"] = "object";
    objs[0] = _wliELUSHTE;
  }
  {
    var obj_hUtAFfhwjO = {};
    obj_hUtAFfhwjO["something"] = "here";
    var hey_MiLRTWbnsp = {};
    hey_MiLRTWbnsp["me"] = true;
    hey_MiLRTWbnsp["anIntVariable"] = 69;
    obj_hUtAFfhwjO["hey"] = hey_MiLRTWbnsp;
    objs[1] = obj_hUtAFfhwjO;
  }
  {
    var obj_AdncpFqHZT = {};
    obj_AdncpFqHZT["something"] = "here";
    var hey_LlooSNEfKS = {};
    hey_LlooSNEfKS["me"] = true;
    hey_LlooSNEfKS["anIntVariable"] = 69;
    obj_AdncpFqHZT["hey"] = hey_LlooSNEfKS;
    objs[2] = obj_AdncpFqHZT;
  }
}

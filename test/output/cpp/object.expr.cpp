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
  var hey_HhaFBNDrfE = {};
  hey_HhaFBNDrfE["me"] = true;
  hey_HhaFBNDrfE["anIntVariable"] = 69;
  obj["hey"] = hey_HhaFBNDrfE;

  var objs[] = {};
  {
    var _GMbqGlCnIW = {};
    _GMbqGlCnIW["another"] = "object";
    objs[0] = _GMbqGlCnIW;
  }
  {
    var obj_UXgprvNahx = {};
    obj_UXgprvNahx["something"] = "here";
    var hey_YEXMGlANik = {};
    hey_YEXMGlANik["me"] = true;
    hey_YEXMGlANik["anIntVariable"] = 69;
    obj_UXgprvNahx["hey"] = hey_YEXMGlANik;
    objs[1] = obj_UXgprvNahx;
  }
  {
    var obj_dFLHnVnNmF = {};
    obj_dFLHnVnNmF["something"] = "here";
    var hey_XsWmqzUgsx = {};
    hey_XsWmqzUgsx["me"] = true;
    hey_XsWmqzUgsx["anIntVariable"] = 69;
    obj_dFLHnVnNmF["hey"] = hey_XsWmqzUgsx;
    objs[2] = obj_dFLHnVnNmF;
  }
}

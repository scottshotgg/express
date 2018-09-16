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
  var hey_SiuXGmICyr = {};
  hey_SiuXGmICyr["me"] = true;
  hey_SiuXGmICyr["anIntVariable"] = 69;
  obj["hey"] = hey_SiuXGmICyr;

  var objs[] = {};
  {
    var _mXjCZUZYyf = {};
    _mXjCZUZYyf["another"] = "object";
    objs[0] = _mXjCZUZYyf;
  }
  {
    var obj_wrvYvTzXGZ = {};
    obj_wrvYvTzXGZ["something"] = "here";
    var hey_aclskXftVG = {};
    hey_aclskXftVG["me"] = true;
    hey_aclskXftVG["anIntVariable"] = 69;
    obj_wrvYvTzXGZ["hey"] = hey_aclskXftVG;
    objs[1] = obj_wrvYvTzXGZ;
  }
  {
    var obj_iTmSAFmXmr = {};
    obj_iTmSAFmXmr["something"] = "here";
    var hey_tYmimobctm = {};
    hey_tYmimobctm["me"] = true;
    hey_tYmimobctm["anIntVariable"] = 69;
    obj_iTmSAFmXmr["hey"] = hey_tYmimobctm;
    objs[2] = obj_iTmSAFmXmr;
  }
}

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
  var hey_xLcFozYJMb = {};
  hey_xLcFozYJMb["me"] = true;
  hey_xLcFozYJMb["anIntVariable"] = 69;
  obj["hey"] = hey_xLcFozYJMb;

  var objs[] = {};
  {
    var _cQqtHRjCil = {};
    _cQqtHRjCil["another"] = "object";
    objs[0] = _cQqtHRjCil;
  }
  {
    var obj_ipbNvCKUbF = {};
    obj_ipbNvCKUbF["something"] = "here";
    var hey_FisqAFkcKk = {};
    hey_FisqAFkcKk["me"] = true;
    hey_FisqAFkcKk["anIntVariable"] = 69;
    obj_ipbNvCKUbF["hey"] = hey_FisqAFkcKk;
    objs[1] = obj_ipbNvCKUbF;
  }
  {
    var obj_ZZnqIMqttV = {};
    obj_ZZnqIMqttV["something"] = "here";
    var hey_mYzClBsIFN = {};
    hey_mYzClBsIFN["me"] = true;
    hey_mYzClBsIFN["anIntVariable"] = 69;
    obj_ZZnqIMqttV["hey"] = hey_mYzClBsIFN;
    objs[2] = obj_ZZnqIMqttV;
  }
}

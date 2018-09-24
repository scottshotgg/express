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
  var hey_geqzKQWRwo = {};
  hey_geqzKQWRwo["me"] = true;
  hey_geqzKQWRwo["anIntVariable"] = 69;
  obj["hey"] = hey_geqzKQWRwo;

  var objs[] = {};
  {
    var _MCWNmyFwLV = {};
    _MCWNmyFwLV["another"] = "object";
    objs[0] = _MCWNmyFwLV;
  }
  {
    var obj_NSKkMLjdyU = {};
    obj_NSKkMLjdyU["something"] = "here";
    var hey_HmQBEpeFok = {};
    hey_HmQBEpeFok["me"] = true;
    hey_HmQBEpeFok["anIntVariable"] = 69;
    obj_NSKkMLjdyU["hey"] = hey_HmQBEpeFok;
    objs[1] = obj_NSKkMLjdyU;
  }
  {
    var obj_PEfxobmkCe = {};
    obj_PEfxobmkCe["something"] = "here";
    var hey_HCMtaEcpbC = {};
    hey_HCMtaEcpbC["me"] = true;
    hey_HCMtaEcpbC["anIntVariable"] = 69;
    obj_PEfxobmkCe["hey"] = hey_HCMtaEcpbC;
    objs[2] = obj_PEfxobmkCe;
  }
}

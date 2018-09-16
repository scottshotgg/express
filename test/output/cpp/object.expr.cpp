#include "/home/scottshotgg/Development/go/src/github.com/scottshotgg/express/lib/defer.cpp"
#include "/home/scottshotgg/Development/go/src/github.com/scottshotgg/express/lib/file.cpp"
#include "/home/scottshotgg/Development/go/src/github.com/scottshotgg/express/lib/std.cpp"
#include "/home/scottshotgg/Development/go/src/github.com/scottshotgg/express/lib/var.cpp"
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
  var hey_lJhhNgSQRv = {};
  hey_lJhhNgSQRv["me"] = true;
  hey_lJhhNgSQRv["anIntVariable"] = 69;
  obj["hey"] = hey_lJhhNgSQRv;

  var objs[] = {};
  {
    var _mqPpsepUyT = {};
    _mqPpsepUyT["another"] = "object";
    objs[0] = _mqPpsepUyT;
  }
  {
    var obj_fpokmdBXTD = {};
    obj_fpokmdBXTD["something"] = "here";
    var hey_mxTvHfxcCC = {};
    hey_mxTvHfxcCC["me"] = true;
    hey_mxTvHfxcCC["anIntVariable"] = 69;
    obj_fpokmdBXTD["hey"] = hey_mxTvHfxcCC;
    objs[1] = obj_fpokmdBXTD;
  }
  {
    var obj_ggVHwgdoUI = {};
    obj_ggVHwgdoUI["something"] = "here";
    var hey_kSxEBcQrho = {};
    hey_kSxEBcQrho["me"] = true;
    hey_kSxEBcQrho["anIntVariable"] = 69;
    obj_ggVHwgdoUI["hey"] = hey_kSxEBcQrho;
    objs[2] = obj_ggVHwgdoUI;
  }
}

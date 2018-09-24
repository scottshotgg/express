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

var increment(var i) {
  defer onReturnFuncs;
  {
    defer onLeaveFuncs;

    var _mfdeOrpvpO = {};
    _mfdeOrpvpO["something"] = "else";
    return _mfdeOrpvpO;
  }
}
int main() {

  var _gnyboxDgFJ = {};
  _gnyboxDgFJ["another"] = "object";

  increment(_gnyboxDgFJ);
}

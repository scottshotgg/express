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

var increment(var i) {
  defer onReturnFuncs;
  {
    defer onLeaveFuncs;

    var _aGxDptipWN = {};
    _aGxDptipWN["something"] = "else";
    return _aGxDptipWN;
  }
}
int main() {

  var _IarNahrYJy = {};
  _IarNahrYJy["another"] = "object";

  increment(_IarNahrYJy);
}

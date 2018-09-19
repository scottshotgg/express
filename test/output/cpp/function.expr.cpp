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

    var _hZmuGJWcdp = {};
    _hZmuGJWcdp["something"] = "else";
    return _hZmuGJWcdp;
  }
}
int main() {

  var _TBPZpSqyqk = {};
  _TBPZpSqyqk["another"] = "object";

  increment(_TBPZpSqyqk);
}

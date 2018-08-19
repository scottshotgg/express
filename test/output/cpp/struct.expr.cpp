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

int printStuff(int k) {
  defer onReturnFuncs;
  {
    defer onLeaveFuncs;

    {
      int i = 0;
      while (i < k) {
        {
          defer onLeaveFuncs;

          onExitFuncs.deferStack.push([=](...) { Println("on exit", i); });

          onReturnFuncs.deferStack.push([=](...) { Println("on return", i); });

          onLeaveFuncs.deferStack.push([=](...) { Println("on leave", i); });

          onReturnFuncs.deferStack.push([=](...) { Println("defer", i); });
        }
        i += 1;
      }
    }

    return 0;
  }
}

var increment(var i) {
  defer onReturnFuncs;
  {
    defer onLeaveFuncs;

    var _OfjbvSDgfb = {};
    _OfjbvSDgfb["something"] = "else";
    return _OfjbvSDgfb;
  }
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
  var thing_iBJvCQVvuK = {};
  thing_iBJvCQVvuK["woah"] = 0;
  Thing["thing"] = thing_iBJvCQVvuK;

  var something = {};
  something["fieldA"] = 0;
  something["stringField"] = "";
  something["false_field"] = false;
  something["anotherFielderino"] = 0.000000;
  var thing_jqwwkiUlbG = {};
  thing_jqwwkiUlbG["woah"] = 0;
  something["thing"] = thing_jqwwkiUlbG;

  Println("something", something);

  var something2 = {};
  something2["fieldA"] = 912559;
  something2["stringField"] = "";
  something2["false_field"] = false;
  something2["anotherFielderino"] = 0.000000;
  var thing_ogJtsMFzxw = {};
  thing_ogJtsMFzxw["woah"] = 0;
  something2["thing"] = thing_ogJtsMFzxw;

  Println("something2", something2);

  var something3 = {};
  something3["fieldA"] = 0;
  something3["stringField"] = "chyah brah";
  something3["false_field"] = false;
  something3["anotherFielderino"] = 0.000000;
  var thing_UBClBzODIK = {};
  thing_UBClBzODIK["woah"] = 0;
  something3["thing"] = thing_UBClBzODIK;

  Println("something3", something3);
}

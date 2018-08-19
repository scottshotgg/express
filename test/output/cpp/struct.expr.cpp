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

    var _VUegFoEXjc = {};
    _VUegFoEXjc["something"] = "else";
    return _VUegFoEXjc;
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
  var thing_iLeSuAOfdg = {};
  thing_iLeSuAOfdg["woah"] = 0;
  Thing["thing"] = thing_iLeSuAOfdg;

  var something = {};
  something["fieldA"] = 0;
  something["stringField"] = "";
  something["false_field"] = false;
  something["anotherFielderino"] = 0.000000;
  var thing_dJALjDjijI = {};
  thing_dJALjDjijI["woah"] = 0;
  something["thing"] = thing_dJALjDjijI;

  Println("something", something);

  var something2 = {};
  something2["fieldA"] = 912559;
  something2["stringField"] = "";
  something2["false_field"] = false;
  something2["anotherFielderino"] = 0.000000;
  var thing_BAReiIuZQj = {};
  thing_BAReiIuZQj["woah"] = 0;
  something2["thing"] = thing_BAReiIuZQj;

  Println("something2", something2);

  var something3 = {};
  something3["fieldA"] = 0;
  something3["stringField"] = "chyah brah";
  something3["false_field"] = false;
  something3["anotherFielderino"] = 0.000000;
  var thing_qdgKDiefgN = {};
  thing_qdgKDiefgN["woah"] = 0;
  something3["thing"] = thing_qdgKDiefgN;

  Println("something3", something3);
}

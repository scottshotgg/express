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

    var _xersmSsRmO = {};
    _xersmSsRmO["something"] = "else";
    return _xersmSsRmO;
  }
}

var someFunction(int arg1) {
  defer onReturnFuncs;
  {
    defer onLeaveFuncs;

    {
      int arrayBoi_1534316920[] = {2, 4, 5, 9};
      int k = 0;
      int k_1534316920 = 0;
      while (4) {
        {
          defer onLeaveFuncs;

          k = arrayBoi_1534316920[k_1534316920];

          onReturnFuncs.deferStack.push([=](...) { Println("value: ", k); });
        }
        k_1534316920 += 1;
      }
    }

    var _SAuiNIjEsl = {};
    _SAuiNIjEsl["value"] = 1000;
    return _SAuiNIjEsl;
  }
}

int main() {

  onExitFuncs.deferStack.push([=](...) { Println("hey its me", ""); });

  onExitFuncs.deferStack.push([=](...) { Println("does this work", ""); });
}

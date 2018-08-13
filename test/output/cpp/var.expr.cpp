#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/defer.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/file.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/std.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/var.cpp"
#include <string>
defer onExitFuncs;

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
    var _RmmaeIcMxV = {};
    _RmmaeIcMxV["something"] = "else";
    return _RmmaeIcMxV;
  }
}

var someFunction(int arg1) {
  defer onReturnFuncs;
  {
    defer onLeaveFuncs;

    {
      int arrayBoi_1534201942[] = {2, 4, 5, 9};
      int k = 0;
      int k_1534201942 = 0;
      while (4) {
        {
          defer onLeaveFuncs;
          k = arrayBoi_1534201942[k_1534201942];
          onReturnFuncs.deferStack.push([=](...) { Println("value: ", k); });
        }
        k_1534201942 += 1;
      }
    }
    var _QklKVfUFZo = {};
    _QklKVfUFZo["value"] = 1000;
    return _QklKVfUFZo;
  }
}

int main() {
  var hi_my_type_is = "what";
  hi_my_type_is = 666;
  hi_my_type_is = true;
  hi_my_type_is = 2.71828;
}

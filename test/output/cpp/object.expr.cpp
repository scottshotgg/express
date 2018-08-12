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
    var _hVZXpHISeH = {};
    _hVZXpHISeH["something"] = "else";
    return _hVZXpHISeH;
  }
}

int main() {
  var obj = {};
  obj["something"] = "here";
  var hey = {};
  hey["me"] = true;
  hey["anIntVariable"] = 69;
  obj["hey"] = hey;
  var objs[] = {};
  {
    var _CjpHUMUwzU = {};
    _CjpHUMUwzU["another"] = "object";
    objs[0] = _CjpHUMUwzU;
  }
  {
    var obj_NVmYBdKEkN = {};
    obj_NVmYBdKEkN["something"] = "here";
    var hey = {};
    hey["me"] = true;
    hey["anIntVariable"] = 69;
    obj_NVmYBdKEkN["hey"] = hey;
    objs[1] = obj_NVmYBdKEkN;
  }
  {
    var obj_kdIodKiCus = {};
    obj_kdIodKiCus["something"] = "here";
    var hey = {};
    hey["me"] = true;
    hey["anIntVariable"] = 69;
    obj_kdIodKiCus["hey"] = hey;
    objs[2] = obj_kdIodKiCus;
  }
}

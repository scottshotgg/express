#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/defer.cpp"
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
          onExitFuncs.deferStack.push([=](...) { println("on exit", i); });
          onReturnFuncs.deferStack.push([=](...) { println("on return", i); });
          onReturnFuncs.deferStack.push([=](...) { println("defer", i); });
          onLeaveFuncs.deferStack.push([=](...) { println("on leave", i); });
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
    var _DOuORuLQrp = {};
    _DOuORuLQrp["something"] = "else";
    return _DOuORuLQrp;
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
    var _bBEEsuPKLq = {};
    _bBEEsuPKLq["another"] = "object";
    objs[0] = _bBEEsuPKLq;
  }
  {
    var obj_GCBFNRSbfO = {};
    obj_GCBFNRSbfO["something"] = "here";
    var hey = {};
    hey["me"] = true;
    hey["anIntVariable"] = 69;
    obj_GCBFNRSbfO["hey"] = hey;
    objs[1] = obj_GCBFNRSbfO;
  }
  {
    var obj_wLeQHIwPNJ = {};
    obj_wLeQHIwPNJ["something"] = "here";
    var hey = {};
    hey["me"] = true;
    hey["anIntVariable"] = 69;
    obj_wLeQHIwPNJ["hey"] = hey;
    objs[2] = obj_wLeQHIwPNJ;
  }
}

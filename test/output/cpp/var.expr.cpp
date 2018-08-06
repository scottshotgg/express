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
    var _XpglhUvVjQ = {};
    _XpglhUvVjQ["something"] = "else";
    return _XpglhUvVjQ;
  }
}

var someFunction(int arg1) {
  defer onReturnFuncs;
  {
    defer onLeaveFuncs;

    {
      int arrayBoi_1533588450[] = {2, 4, 5, 9};
      int k = 0;
      int k_1533588450 = 0;
      while (4) {
        {
          defer onLeaveFuncs;
          k = arrayBoi_1533588450[k_1533588450];
          onReturnFuncs.deferStack.push([=](...) { println("value: ", k); });
        }
        k_1533588450 += 1;
      }
    }
    var _mhBkjwmqMY = {};
    _mhBkjwmqMY["value"] = 1000;
    return _mhBkjwmqMY;
  }
}

int main() {
  var hi_my_type_is = "what";
  hi_my_type_is = 666;
  hi_my_type_is = true;
  hi_my_type_is = 2.71828;
}

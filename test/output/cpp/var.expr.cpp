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

var someFunction(int arg1) {
  defer onReturnFuncs;
  {
    defer onLeaveFuncs;

    {
      int arrayBoi_1534341376[] = {2, 4, 5, 9};
      int k = 0;
      int k_1534341376 = 0;
      while (k_1534341376 < 4) {
        {
          defer onLeaveFuncs;

          k = arrayBoi_1534341376[k_1534341376];

          onReturnFuncs.deferStack.push(
              [=](...) { Println("value of k: ", k); });

          if (arg1 < k) {
            defer onLeaveFuncs;

            var _lhwgZYwoSI = {};
            _lhwgZYwoSI["value"] = k;
            return _lhwgZYwoSI;
          }
        }
        k_1534341376 += 1;
      }
    }

    var _hBTzACTzhA = {};
    _hBTzACTzhA["something"] = 0;
    return _hBTzACTzhA;
  }
}

int main() {

  var hi_my_type_is = "what";

  hi_my_type_is = 666;

  hi_my_type_is = true;

  hi_my_type_is = 2.71828;
}

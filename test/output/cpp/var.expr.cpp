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

    var _qagxBYLiDh = {};
    _qagxBYLiDh["something"] = "else";
    return _qagxBYLiDh;
  }
}
var someFunction(int arg1) {
  defer onReturnFuncs;
  {
    defer onLeaveFuncs;

    {
      int arrayBoi_1534716692[] = {2, 4, 5, 9};
      int k = 0;
      int k_1534716692 = 0;
      while (k_1534716692 < 4) {
        {
          defer onLeaveFuncs;

          k = arrayBoi_1534716692[k_1534716692];

          onLeaveFuncs.deferStack.push(
              [=](...) { Println("value of k: ", k); });

          onReturnFuncs.deferStack.push(
              [=](...) { Println("value of k: ", k); });

          if (arg1 < k) {
            defer onLeaveFuncs;

            var _sDUCvSBOxr = {};
            _sDUCvSBOxr["value"] = k;
            return _sDUCvSBOxr;
          }
        }
        k_1534716692 += 1;
      }
    }

    var _vyOxdbQnLP = {};
    _vyOxdbQnLP["something"] = 0;
    return _vyOxdbQnLP;
  }
}
int main() {

  var hi_my_type_is = "what";

  hi_my_type_is = 666;

  hi_my_type_is = true;

  hi_my_type_is = 2.71828;
}

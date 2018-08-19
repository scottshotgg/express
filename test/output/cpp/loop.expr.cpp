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
int main() {

  int f = 0;

  {
    int i = 1;
    while (i < 10) {
      {
        defer onLeaveFuncs;

        f = i;

        int h = 1;
      }
      i += 1;
    }
  }

  {
    int arrayBoi_1534716689[] = {1, 2, 4};
    int i = 0;
    int i_1534716689 = 0;
    while (i_1534716689 < 3) {
      {
        defer onLeaveFuncs;

        i = i_1534716689;

        f = i;

        int h = 1;
      }
      i_1534716689 += 1;
    }
  }

  int countdown[] = {9, 8, 7, 5, 4, 3, 2, 1};

  {
    int i = 0;
    int i_1534716689 = 0;
    while (i_1534716689 < 8) {
      {
        defer onLeaveFuncs;

        i = countdown[i_1534716689];

        f = i;

        int h = 1;
      }
      i_1534716689 += 1;
    }
  }
}

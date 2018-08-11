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
    var _lLVcINPwbo = {};
    _lLVcINPwbo["something"] = "else";
    return _lLVcINPwbo;
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
    int arrayBoi_1534015483[] = {1, 2, 4};
    int i = 0;
    int i_1534015483 = 0;
    while (3) {
      {
        defer onLeaveFuncs;
        i = i_1534015483;
        f = i;
        int h = 1;
      }
      i_1534015483 += 1;
    }
  }
  int countdown[] = {9, 8, 7, 5, 4, 3, 2, 1};

  {
    int i = 0;
    int i_1534015483 = 0;
    while (8) {
      {
        defer onLeaveFuncs;
        i = countdown[i_1534015483];
        f = i;
        int h = 1;
      }
      i_1534015483 += 1;
    }
  }
}

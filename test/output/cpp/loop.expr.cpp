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
    int arrayBoi_1533588446[] = {1, 2, 4};
    int i = 0;
    int i_1533588446 = 0;
    while (3) {
      {
        defer onLeaveFuncs;
        i = i_1533588446;
        f = i;
        int h = 1;
      }
      i_1533588446 += 1;
    }
  }
  int countdown[] = {9, 8, 7, 5, 4, 3, 2, 1};

  {
    int i = 0;
    int i_1533588446 = 0;
    while (8) {
      {
        defer onLeaveFuncs;
        i = countdown[i_1533588446];
        f = i;
        int h = 1;
      }
      i_1533588446 += 1;
    }
  }
}

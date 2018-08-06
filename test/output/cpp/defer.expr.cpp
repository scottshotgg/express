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

int main() {
  onExitFuncs.deferStack.push(
      [=](...) { println("im actually at the end", ""); });

  println("me first", "");

  println(printStuff(3));
  onExitFuncs.deferStack.push(
      [=](...) { println("im actually at the start", ""); });
}

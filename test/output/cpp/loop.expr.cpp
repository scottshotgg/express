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
<<<<<<< HEAD

    var _xersmSsRmO = {};
    _xersmSsRmO["something"] = "else";
    return _xersmSsRmO;
=======
    var _zVjyJcdhMU = {};
    _zVjyJcdhMU["something"] = "else";
    return _zVjyJcdhMU;
>>>>>>> does not work; committing code before click back
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
<<<<<<< HEAD
    int arrayBoi_1534316912[] = {1, 2, 4};
    int i = 0;
    int i_1534316912 = 0;
    while (3) {
      {
        defer onLeaveFuncs;

        i = i_1534316912;

=======
    int arrayBoi_1534206767[] = {1, 2, 4};
    int i = 0;
    int i_1534206767 = 0;
    while (3) {
      {
        defer onLeaveFuncs;
        i = i_1534206767;
>>>>>>> does not work; committing code before click back
        f = i;

        int h = 1;
      }
<<<<<<< HEAD
      i_1534316912 += 1;
=======
      i_1534206767 += 1;
>>>>>>> does not work; committing code before click back
    }
  }

  int countdown[] = {9, 8, 7, 5, 4, 3, 2, 1};

  {
    int i = 0;
<<<<<<< HEAD
    int i_1534316912 = 0;
    while (8) {
      {
        defer onLeaveFuncs;

        i = countdown[i_1534316912];

=======
    int i_1534206767 = 0;
    while (8) {
      {
        defer onLeaveFuncs;
        i = countdown[i_1534206767];
>>>>>>> does not work; committing code before click back
        f = i;

        int h = 1;
      }
<<<<<<< HEAD
      i_1534316912 += 1;
=======
      i_1534206767 += 1;
>>>>>>> does not work; committing code before click back
    }
  }
}

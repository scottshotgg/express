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

int main() {

  var obj = {};
  obj["something"] = "here";
  var hey_DBQPWjcLGf = {};
  hey_DBQPWjcLGf["me"] = true;
  hey_DBQPWjcLGf["anIntVariable"] = 69;
  obj["hey"] = hey_DBQPWjcLGf;

  var objs[] = {};
  {
    var _PHZhPGNrur = {};
    _PHZhPGNrur["another"] = "object";
    objs[0] = _PHZhPGNrur;
  }
  {
    var obj_LMzKFfAFZn = {};
    obj_LMzKFfAFZn["something"] = "here";
    var hey_zJcufPhZMh = {};
    hey_zJcufPhZMh["me"] = true;
    hey_zJcufPhZMh["anIntVariable"] = 69;
    obj_LMzKFfAFZn["hey"] = hey_zJcufPhZMh;
    objs[1] = obj_LMzKFfAFZn;
  }
  {
    var obj_pgWwIuxvLe = {};
    obj_pgWwIuxvLe["something"] = "here";
    var hey_NMvgxCfGok = {};
    hey_NMvgxCfGok["me"] = true;
    hey_NMvgxCfGok["anIntVariable"] = 69;
    obj_pgWwIuxvLe["hey"] = hey_NMvgxCfGok;
    objs[2] = obj_pgWwIuxvLe;
  }
}

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

  var obj = {};
  obj["something"] = "here";
  var hey_MWRhdCXbLD = {};
  hey_MWRhdCXbLD["me"] = true;
  hey_MWRhdCXbLD["anIntVariable"] = 69;
  obj["hey"] = hey_MWRhdCXbLD;

  var objs[] = {};
  {
    var _SlgnjaZCSz = {};
    _SlgnjaZCSz["another"] = "object";
    objs[0] = _SlgnjaZCSz;
  }
  {
    var obj_LQRzzXFAVG = {};
    obj_LQRzzXFAVG["something"] = "here";
    var hey_DKOdVKXUee = {};
    hey_DKOdVKXUee["me"] = true;
    hey_DKOdVKXUee["anIntVariable"] = 69;
    obj_LQRzzXFAVG["hey"] = hey_DKOdVKXUee;
    objs[1] = obj_LQRzzXFAVG;
  }
  {
    var obj_ONfdLtrXJT = {};
    obj_ONfdLtrXJT["something"] = "here";
    var hey_dEAYwTJvqr = {};
    hey_dEAYwTJvqr["me"] = true;
    hey_dEAYwTJvqr["anIntVariable"] = 69;
    obj_ONfdLtrXJT["hey"] = hey_dEAYwTJvqr;
    objs[2] = obj_ONfdLtrXJT;
  }
}

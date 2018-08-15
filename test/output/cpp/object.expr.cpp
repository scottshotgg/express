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

    var _aXYFzYytAR = {};
    _aXYFzYytAR["something"] = "else";
    return _aXYFzYytAR;
  }
}

int main() {

  var obj = {};
  obj["something"] = "here";
  var hey_pWLaIextde = {};
  hey_pWLaIextde["me"] = true;
  hey_pWLaIextde["anIntVariable"] = 69;
  obj["hey"] = hey_pWLaIextde;

  var objs[] = {};
  {
    var _EUdvkgyRqM = {};
    _EUdvkgyRqM["another"] = "object";
    objs[0] = _EUdvkgyRqM;
  }
  {
    var obj_YhYjzolAqA = {};
    obj_YhYjzolAqA["something"] = "here";
    var hey_IXLLHiteZM = {};
    hey_IXLLHiteZM["me"] = true;
    hey_IXLLHiteZM["anIntVariable"] = 69;
    obj_YhYjzolAqA["hey"] = hey_IXLLHiteZM;
    objs[1] = obj_YhYjzolAqA;
  }
  {
    var obj_GIdIVUvFdq = {};
    obj_GIdIVUvFdq["something"] = "here";
    var hey_TZWaPFeSli = {};
    hey_TZWaPFeSli["me"] = true;
    hey_TZWaPFeSli["anIntVariable"] = 69;
    obj_GIdIVUvFdq["hey"] = hey_TZWaPFeSli;
    objs[2] = obj_GIdIVUvFdq;
  }
}

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

    var _VUegFoEXjc = {};
    _VUegFoEXjc["something"] = "else";
    return _VUegFoEXjc;
  }
}
int main() {

  var obj = {};
  obj["something"] = "here";
  var hey_GYuhEgiEGH = {};
  hey_GYuhEgiEGH["me"] = true;
  hey_GYuhEgiEGH["anIntVariable"] = 69;
  obj["hey"] = hey_GYuhEgiEGH;

  var objs[] = {};
  {
    var _qNGsqmcFXe = {};
    _qNGsqmcFXe["another"] = "object";
    objs[0] = _qNGsqmcFXe;
  }
  {
    var obj_aiuBFmnyKS = {};
    obj_aiuBFmnyKS["something"] = "here";
    var hey_nfPhRwGNqK = {};
    hey_nfPhRwGNqK["me"] = true;
    hey_nfPhRwGNqK["anIntVariable"] = 69;
    obj_aiuBFmnyKS["hey"] = hey_nfPhRwGNqK;
    objs[1] = obj_aiuBFmnyKS;
  }
  {
    var obj_kOZRcoZEdG = {};
    obj_kOZRcoZEdG["something"] = "here";
    var hey_UEiGTwudmG = {};
    hey_UEiGTwudmG["me"] = true;
    hey_UEiGTwudmG["anIntVariable"] = 69;
    obj_kOZRcoZEdG["hey"] = hey_UEiGTwudmG;
    objs[2] = obj_kOZRcoZEdG;
  }
}

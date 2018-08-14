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

  var obj = {};
  obj["something"] = "here";
  var hey_rbzVqmlkKb = {};
  hey_rbzVqmlkKb["me"] = true;
  hey_rbzVqmlkKb["anIntVariable"] = 69;
  obj["hey"] = hey_rbzVqmlkKb;

  var objs[] = {};
  {
<<<<<<< HEAD
    var _KKUeupHTAD = {};
    _KKUeupHTAD["another"] = "object";
    objs[0] = _KKUeupHTAD;
  }
  {
    var obj_KcxoRjjOnF = {};
    obj_KcxoRjjOnF["something"] = "here";
    var hey_IyNrtPknZs = {};
    hey_IyNrtPknZs["me"] = true;
    hey_IyNrtPknZs["anIntVariable"] = 69;
    obj_KcxoRjjOnF["hey"] = hey_IyNrtPknZs;
    objs[1] = obj_KcxoRjjOnF;
  }
  {
    var obj_XhOVKTKrDS = {};
    obj_XhOVKTKrDS["something"] = "here";
    var hey_JoKirTevrz = {};
    hey_JoKirTevrz["me"] = true;
    hey_JoKirTevrz["anIntVariable"] = 69;
    obj_XhOVKTKrDS["hey"] = hey_JoKirTevrz;
    objs[2] = obj_XhOVKTKrDS;
=======
    var _eiwPVbvYur = {};
    _eiwPVbvYur["another"] = "object";
    objs[0] = _eiwPVbvYur;
  }
  {
    var obj_ABCbDmZSYI = {};
    obj_ABCbDmZSYI["something"] = "here";
    var hey = {};
    hey["me"] = true;
    hey["anIntVariable"] = 69;
    obj_ABCbDmZSYI["hey"] = hey;
    objs[1] = obj_ABCbDmZSYI;
  }
  {
    var obj_CrMGMktHir = {};
    obj_CrMGMktHir["something"] = "here";
    var hey = {};
    hey["me"] = true;
    hey["anIntVariable"] = 69;
    obj_CrMGMktHir["hey"] = hey;
    objs[2] = obj_CrMGMktHir;
>>>>>>> does not work; committing code before click back
  }
}

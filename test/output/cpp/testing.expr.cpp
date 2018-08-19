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
var someFunction(int arg1) {
  defer onReturnFuncs;
  {
    defer onLeaveFuncs;

    {
      int arrayBoi_1534716692[] = {2, 4, 5, 9};
      int k = 0;
      int k_1534716692 = 0;
      while (k_1534716692 < 4) {
        {
          defer onLeaveFuncs;

          k = arrayBoi_1534716692[k_1534716692];

          onLeaveFuncs.deferStack.push(
              [=](...) { Println("value of k: ", k); });

          onReturnFuncs.deferStack.push(
              [=](...) { Println("value of k: ", k); });

          if (arg1 < k) {
            defer onLeaveFuncs;

            var _sDUCvSBOxr = {};
            _sDUCvSBOxr["value"] = k;
            return _sDUCvSBOxr;
          }
        }
        k_1534716692 += 1;
      }
    }

    var _vyOxdbQnLP = {};
    _vyOxdbQnLP["something"] = 0;
    return _vyOxdbQnLP;
  }
}
int main() {

  Println("return value", someFunction(7));

  Println();

  Println();

  onExitFuncs.deferStack.push(
      [=](...) { Println("return value", someFunction(9)); });

  var Engine = {};
  Engine["Displacement"] = 2.4;
  Engine["Type"] = "v6";
  Engine["HP"] = 235;
  Engine["OilType"] = "5w-20";
  var Warnings_ejeDMOLCMn = {};
  Engine["Warnings"] = Warnings_ejeDMOLCMn;
  Engine["Warnings"] = Warnings_ejeDMOLCMn;

  var things_jpfkbCBZTb = {};
  things_jpfkbCBZTb["thingy"] = 0;

  var normalWarnings_ttYBjTVCqJ = {};
  normalWarnings_ttYBjTVCqJ["lowOil"] = false;
  normalWarnings_ttYBjTVCqJ["lowGas"] = false;
  normalWarnings_ttYBjTVCqJ["highTemp"] = false;
  normalWarnings_ttYBjTVCqJ["lowTirePressure"] = false;

  var Car = {};
  var id_LWNrDgYnPP = {};
  Car["id"] = id_LWNrDgYnPP;
  Car["id"] = id_LWNrDgYnPP;

  Car["Type"] = "car";
  Car["New"] = false;
  Car["numOfWheels"] = 4;
  var engine_DFSjHaakAs = {};
  engine_DFSjHaakAs["Displacement"] = 2.4;
  engine_DFSjHaakAs["Type"] = "v6";
  engine_DFSjHaakAs["HP"] = 199;
  engine_DFSjHaakAs["OilType"] = "5w-20";
  var Warnings_QUYNgADWin = {};
  Warnings_QUYNgADWin["lowOil"] = false;
  Warnings_QUYNgADWin["lowGas"] = false;
  Warnings_QUYNgADWin["highTemp"] = false;
  Warnings_QUYNgADWin["lowTirePressure"] = false;
  engine_DFSjHaakAs["Warnings"] = Warnings_QUYNgADWin;
  engine_DFSjHaakAs["Warnings"] = Warnings_QUYNgADWin;

  Car["engine"] = engine_DFSjHaakAs;

  onExitFuncs.deferStack.push([=](...) { Println("Most cars:", Car); });

  var something = someFunction(7);

  Println(something);

  std::string name = "hey this is an id";

  var oldCar = {};
  var id_KVLjHCkOUm = "hey this is an id";
  oldCar["id"] = id_KVLjHCkOUm;

  oldCar["Type"] = "car";
  oldCar["New"] = false;
  oldCar["numOfWheels"] = 4;
  var engine_vfWrQouKuT = {};
  engine_vfWrQouKuT["Displacement"] = 2.4;
  engine_vfWrQouKuT["Type"] = "v6";
  engine_vfWrQouKuT["HP"] = 160;
  engine_vfWrQouKuT["OilType"] = "5w-20";
  engine_vfWrQouKuT["Warnings"] = something;
  oldCar["engine"] = engine_vfWrQouKuT;

  Println("Old cars:", oldCar);
}

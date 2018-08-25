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

var someFunction(int arg1) {
  defer onReturnFuncs;
  {
    defer onLeaveFuncs;

    {
      int arrayBoi_1535160066[] = {2, 4, 5, 9};
      int k = 0;
      int k_1535160066 = 0;
      while (k_1535160066 < 4) {
        {
          defer onLeaveFuncs;

          k = arrayBoi_1535160066[k_1535160066];

          onLeaveFuncs.deferStack.push(
              [=](...) { Println("value of k: ", k); });

          onReturnFuncs.deferStack.push(
              [=](...) { Println("value of k: ", k); });

          if (arg1 < k) {
            defer onLeaveFuncs;

            var _VbqVWpNekf = {};
            _VbqVWpNekf["value"] = k;
            return _VbqVWpNekf;
          }
        }
        k_1535160066 += 1;
      }
    }

    var _TiXjAxUAtI = {};
    _TiXjAxUAtI["something"] = 0;
    return _TiXjAxUAtI;
  }
}
int main() {

  int someArgument = 7;

  Println("return value", someFunction(someArgument));

  Println();

  Println();

  onExitFuncs.deferStack.push(
      [=](...) { Println("return value", someFunction(someArgument)); });

  someArgument = 9;

  onExitFuncs.deferStack.push(
      [=](...) { Println("return value", someFunction(someArgument)); });

  var Engine = {};
  Engine["Displacement"] = 2.4;
  Engine["Type"] = "v6";
  Engine["HP"] = 235;
  Engine["OilType"] = "5w-20";
  var Warnings_rWokffsYOJ = Warnings_rWokffsYOJ;
  Engine["Warnings"] = Warnings_rWokffsYOJ;

  var things_gMuJjNwgwx = {};
  things_gMuJjNwgwx["thingy"] = 0;

  var normalWarnings_EsVBmBrxGg = {};
  normalWarnings_EsVBmBrxGg["lowOil"] = false;
  normalWarnings_EsVBmBrxGg["lowGas"] = false;
  normalWarnings_EsVBmBrxGg["highTemp"] = false;
  normalWarnings_EsVBmBrxGg["lowTirePressure"] = false;

  var Car = {};
  var id_CfIWAWaNfL = {};
  Car["id"] = id_CfIWAWaNfL;
  Car["id"] = id_CfIWAWaNfL;

  Car["Type"] = "car";
  Car["New"] = false;
  Car["numOfWheels"] = 4;
  var engine_eUirjRLQEW = {};
  engine_eUirjRLQEW["Displacement"] = 2.4;
  engine_eUirjRLQEW["Type"] = "v6";
  engine_eUirjRLQEW["HP"] = 199;
  engine_eUirjRLQEW["OilType"] = "5w-20";
  var Warnings_bsKADBtrMH = {};
  Warnings_bsKADBtrMH["lowOil"] = false;
  Warnings_bsKADBtrMH["lowGas"] = false;
  Warnings_bsKADBtrMH["highTemp"] = false;
  Warnings_bsKADBtrMH["lowTirePressure"] = false;
  engine_eUirjRLQEW["Warnings"] = Warnings_bsKADBtrMH;
  engine_eUirjRLQEW["Warnings"] = Warnings_bsKADBtrMH;

  Car["engine"] = engine_eUirjRLQEW;

  onExitFuncs.deferStack.push([=](...) { Println("Most cars:", Car); });

  var something = someFunction(7);

  Println(something);

  std::string name = "hey this is an id";

  var oldCar = {};
  var id_yGcYxAqhxz = "hey this is an id";
  oldCar["id"] = id_yGcYxAqhxz;

  oldCar["Type"] = "car";
  oldCar["New"] = false;
  oldCar["numOfWheels"] = 4;
  var engine_kFoqXmEsro = {};
  engine_kFoqXmEsro["Displacement"] = 2.4;
  engine_kFoqXmEsro["Type"] = "v6";
  engine_kFoqXmEsro["HP"] = 160;
  engine_kFoqXmEsro["OilType"] = "5w-20";
  engine_kFoqXmEsro["Warnings"] = something;
  oldCar["engine"] = engine_kFoqXmEsro;

  Println("Old cars:", oldCar);
}

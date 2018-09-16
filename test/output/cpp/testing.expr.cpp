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

var someFunction(int arg1) {
  defer onReturnFuncs;
  {
    defer onLeaveFuncs;

    {
      int arrayBoi_1537004698[] = {2, 4, 5, 9};
      int k = 0;
      int k_1537004698 = 0;
      while (k_1537004698 < 4) {
        {
          defer onLeaveFuncs;

          k = arrayBoi_1537004698[k_1537004698];

          onLeaveFuncs.deferStack.push(
              [=](...) { Println("value of k: ", k); });

          onReturnFuncs.deferStack.push(
              [=](...) { Println("value of k: ", k); });

          if (arg1 < k) {
            defer onLeaveFuncs;

            var _dzyVFPcARN = {};
            _dzyVFPcARN["value"] = k;
            return _dzyVFPcARN;
          }
        }
        k_1537004698 += 1;
      }
    }

    var _kamabuhBlJ = {};
    _kamabuhBlJ["something"] = 0;
    return _kamabuhBlJ;
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
  var Warnings_GWjCrfHoGJ = Warnings_GWjCrfHoGJ;
  Engine["Warnings"] = Warnings_GWjCrfHoGJ;

  var things_KOCqWxOuva = {};
  things_KOCqWxOuva["thingy"] = 0;

  var normalWarnings_kbrlVLOXFS = {};
  normalWarnings_kbrlVLOXFS["lowOil"] = false;
  normalWarnings_kbrlVLOXFS["lowGas"] = false;
  normalWarnings_kbrlVLOXFS["highTemp"] = false;
  normalWarnings_kbrlVLOXFS["lowTirePressure"] = false;

  var Car = {};
  var id_AGRZhHldfx = {};
  Car["id"] = id_AGRZhHldfx;
  Car["id"] = id_AGRZhHldfx;

  Car["Type"] = "car";
  Car["New"] = false;
  Car["numOfWheels"] = 4;
  var engine_yRiaGPJUDo = {};
  engine_yRiaGPJUDo["Displacement"] = 2.4;
  engine_yRiaGPJUDo["Type"] = "v6";
  engine_yRiaGPJUDo["HP"] = 199;
  engine_yRiaGPJUDo["OilType"] = "5w-20";
  var Warnings_ndrXPRMeDU = {};
  Warnings_ndrXPRMeDU["lowOil"] = false;
  Warnings_ndrXPRMeDU["lowGas"] = false;
  Warnings_ndrXPRMeDU["highTemp"] = false;
  Warnings_ndrXPRMeDU["lowTirePressure"] = false;
  engine_yRiaGPJUDo["Warnings"] = Warnings_ndrXPRMeDU;
  engine_yRiaGPJUDo["Warnings"] = Warnings_ndrXPRMeDU;

  Car["engine"] = engine_yRiaGPJUDo;

  onExitFuncs.deferStack.push([=](...) { Println("Most cars:", Car); });

  var something = someFunction(7);

  Println(something);

  std::string name = "hey this is an id";

  var oldCar = {};
  var id_ANNWEcZVqf = "hey this is an id";
  oldCar["id"] = id_ANNWEcZVqf;

  oldCar["Type"] = "car";
  oldCar["New"] = false;
  oldCar["numOfWheels"] = 4;
  var engine_iMiuhnGcOc = {};
  engine_iMiuhnGcOc["Displacement"] = 2.4;
  engine_iMiuhnGcOc["Type"] = "v6";
  engine_iMiuhnGcOc["HP"] = 160;
  engine_iMiuhnGcOc["OilType"] = "5w-20";
  engine_iMiuhnGcOc["Warnings"] = something;
  oldCar["engine"] = engine_iMiuhnGcOc;

  Println("Old cars:", oldCar);
}

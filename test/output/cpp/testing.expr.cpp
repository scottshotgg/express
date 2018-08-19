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
      int arrayBoi_1534557084[] = {2, 4, 5, 9};
      int k = 0;
      int k_1534557084 = 0;
      while (k_1534557084 < 4) {
        {
          defer onLeaveFuncs;

          k = arrayBoi_1534557084[k_1534557084];

          onReturnFuncs.deferStack.push(
              [=](...) { Println("value of k: ", k); });

          if (arg1 < k) {
            defer onLeaveFuncs;

            var _RxfoiFnpbd = {};
            _RxfoiFnpbd["value"] = k;
            return _RxfoiFnpbd;
          }
        }
        k_1534557084 += 1;
      }
    }

    var _meCIXAxHdf = {};
    _meCIXAxHdf["something"] = 0;
    return _meCIXAxHdf;
  }
}

int main() {

  Println("return value", someFunction(7));

  Println();

  Println();

  Println("return value", someFunction(9));

  var Engine = {};
  Engine["Displacement"] = 2.4;
  Engine["Type"] = "v6";
  Engine["HP"] = 235;
  Engine["OilType"] = "5w-20";
  var Warnings_ZNzQrbZeVZ = {};
  Engine["Warnings"] = Warnings_ZNzQrbZeVZ;

  var things = {};
  things["thingy"] = 0;

  var normalWarnings = {};
  normalWarnings["lowOil"] = false;
  normalWarnings["lowGas"] = false;
  normalWarnings["highTemp"] = false;
  normalWarnings["lowTirePressure"] = false;

  var Car = {};
  Car["Type"] = "car";
  Car["New"] = false;
  Car["numOfWheels"] = 4;
  var engine_fNVjXXCyhs = {};
  engine_fNVjXXCyhs["Displacement"] = 2.4;
  engine_fNVjXXCyhs["Type"] = "v6";
  engine_fNVjXXCyhs["HP"] = 199;
  engine_fNVjXXCyhs["OilType"] = "5w-20";
  var Warnings_QXeoiowVUb = {};
  Warnings_QXeoiowVUb["lowOil"] = false;
  Warnings_QXeoiowVUb["lowGas"] = false;
  Warnings_QXeoiowVUb["highTemp"] = false;
  Warnings_QXeoiowVUb["lowTirePressure"] = false;
  engine_fNVjXXCyhs["Warnings"] = Warnings_QXeoiowVUb;
  Car["engine"] = engine_fNVjXXCyhs;

  var name_WHLhquJdHT = "hey this is an id";

  var oldWarnings = {};
  oldWarnings["antique"] = true;

  something(7);

  var oldCar = {};
  oldCar["Type"] = "car";
  oldCar["New"] = false;
  oldCar["numOfWheels"] = 4;
  var engine_nZxjUrCoTB = {};
  engine_nZxjUrCoTB["Displacement"] = 2.4;
  engine_nZxjUrCoTB["Type"] = "v6";
  engine_nZxjUrCoTB["HP"] = 160;
  engine_nZxjUrCoTB["OilType"] = "5w-20";
  var Warnings_LlsonCUKFK = {};
  Warnings_LlsonCUKFK["antique"] = true;
  engine_nZxjUrCoTB["Warnings"] = Warnings_LlsonCUKFK;
  oldCar["engine"] = engine_nZxjUrCoTB;

  Println("Most cars:", Car);

  Println("Old cars:", oldCar);
}

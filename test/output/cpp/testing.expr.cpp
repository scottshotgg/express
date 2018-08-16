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
      int arrayBoi_1534449547[] = {2, 4, 5, 9};
      int k = 0;
      int k_1534449547 = 0;
      while (k_1534449547 < 4) {
        {
          defer onLeaveFuncs;

          k = arrayBoi_1534449547[k_1534449547];

          onReturnFuncs.deferStack.push(
              [=](...) { Println("value of k: ", k); });

          if (arg1 < k) {
            defer onLeaveFuncs;

            var _YjietuaNMQ = {};
            _YjietuaNMQ["value"] = k;
            return _YjietuaNMQ;
          }
        }
        k_1534449547 += 1;
      }
    }

    var _WcVEbJHMZV = {};
    _WcVEbJHMZV["something"] = 0;
    return _WcVEbJHMZV;
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
  var Warnings_CakKkSLHGM = {};
  Engine["Warnings"] = Warnings_CakKkSLHGM;

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
  var engine_EnaMwfxewN = {};
  engine_EnaMwfxewN["Displacement"] = 2.4;
  engine_EnaMwfxewN["Type"] = "v6";
  engine_EnaMwfxewN["HP"] = 199;
  engine_EnaMwfxewN["OilType"] = "5w-20";
  var Warnings_oYUaoJnHNt = {};
  Warnings_oYUaoJnHNt["lowOil"] = false;
  Warnings_oYUaoJnHNt["lowGas"] = false;
  Warnings_oYUaoJnHNt["highTemp"] = false;
  Warnings_oYUaoJnHNt["lowTirePressure"] = false;
  engine_EnaMwfxewN["Warnings"] = Warnings_oYUaoJnHNt;
  Car["engine"] = engine_EnaMwfxewN;

  var oldWarnings = {};
  oldWarnings["antique"] = true;

  var oldCar = {};
  oldCar["Type"] = "car";
  oldCar["New"] = false;
  oldCar["numOfWheels"] = 4;
  var engine_BcdNATmDgE = {};
  engine_BcdNATmDgE["Displacement"] = 2.4;
  engine_BcdNATmDgE["Type"] = "v6";
  engine_BcdNATmDgE["HP"] = 160;
  engine_BcdNATmDgE["OilType"] = "5w-20";
  var Warnings_DCmULscuZt = {};
  Warnings_DCmULscuZt["antique"] = true;
  engine_BcdNATmDgE["Warnings"] = Warnings_DCmULscuZt;
  oldCar["engine"] = engine_BcdNATmDgE;

  Println("Most cars:", Car);

  Println("Old cars:", oldCar);
}

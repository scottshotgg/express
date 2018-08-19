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
var someFunction(int arg1) {
  defer onReturnFuncs;
  {
    defer onLeaveFuncs;

    {
      int arrayBoi_1534710374[] = {2, 4, 5, 9};
      int k = 0;
      int k_1534710374 = 0;
      while (k_1534710374 < 4) {
        {
          defer onLeaveFuncs;

          k = arrayBoi_1534710374[k_1534710374];

          onReturnFuncs.deferStack.push(
              [=](...) { Println("value of k: ", k); });

          if (arg1 < k) {
            defer onLeaveFuncs;

            var _gxXjofqyMw = {};
            _gxXjofqyMw["value"] = k;
            return _gxXjofqyMw;
          }
        }
        k_1534710374 += 1;
      }
    }

    var _UziPfMRtSs = {};
    _UziPfMRtSs["something"] = 0;
    return _UziPfMRtSs;
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
  var Warnings_NwKOWKEBPh = {};
  Engine["Warnings"] = Warnings_NwKOWKEBPh;

  var things_pduEwKDisF = {};
  things_pduEwKDisF["thingy"] = 0;

  var normalWarnings_zZvUfKjipX = {};
  normalWarnings_zZvUfKjipX["lowOil"] = false;
  normalWarnings_zZvUfKjipX["lowGas"] = false;
  normalWarnings_zZvUfKjipX["highTemp"] = false;
  normalWarnings_zZvUfKjipX["lowTirePressure"] = false;

  var Car = {};
  var id_AdbBXwWHuw = {};
  Car["id"] = id_AdbBXwWHuw;

  Car["Type"] = "car";
  Car["New"] = false;
  Car["numOfWheels"] = 4;
  var engine_QkxqNcZOWD = {};
  engine_QkxqNcZOWD["Displacement"] = 2.4;
  engine_QkxqNcZOWD["Type"] = "v6";
  engine_QkxqNcZOWD["HP"] = 199;
  engine_QkxqNcZOWD["OilType"] = "5w-20";
  var Warnings_pHBtsWBfno = {};
  Warnings_pHBtsWBfno["lowOil"] = false;
  Warnings_pHBtsWBfno["lowGas"] = false;
  Warnings_pHBtsWBfno["highTemp"] = false;
  Warnings_pHBtsWBfno["lowTirePressure"] = false;
  engine_QkxqNcZOWD["Warnings"] = Warnings_pHBtsWBfno;

  Car["engine"] = engine_QkxqNcZOWD;

  std::string name = "hey this is an id";

  var something = someFunction(7);

  Println(something);

  var oldCar = {};
  var id_YMzunsayhi = "hey this is an id";

  oldCar["Type"] = "car";
  oldCar["New"] = false;
  oldCar["numOfWheels"] = 4;
  var engine_lvObOleaXZ = {};
  engine_lvObOleaXZ["Displacement"] = 2.4;
  engine_lvObOleaXZ["Type"] = "v6";
  engine_lvObOleaXZ["HP"] = 160;
  engine_lvObOleaXZ["OilType"] = "5w-20";
  engine_lvObOleaXZ["Warnings"] = something;
  oldCar["engine"] = engine_lvObOleaXZ;

  Println("Most cars:", Car);

  Println("Old cars:", oldCar);
}

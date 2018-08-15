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
      int arrayBoi_1534363714[] = {2, 4, 5, 9};
      int k = 0;
      int k_1534363714 = 0;
      while (k_1534363714 < 4) {
        {
          defer onLeaveFuncs;

          k = arrayBoi_1534363714[k_1534363714];

          onReturnFuncs.deferStack.push(
              [=](...) { Println("value of k: ", k); });

          if (arg1 < k) {
            defer onLeaveFuncs;

            var _QVVCHDyEUS = {};
            _QVVCHDyEUS["value"] = k;
            return _QVVCHDyEUS;
          }
        }
        k_1534363714 += 1;
      }
    }

    var _rWEbDxcJRM = {};
    _rWEbDxcJRM["something"] = 0;
    return _rWEbDxcJRM;
  }
}

int main() {

  onExitFuncs.deferStack.push([=](...) { Println("hey its me", ""); });

  onExitFuncs.deferStack.push([=](...) { Println("does this work", ""); });

  Println("return value", someFunction(7));

  Println();

  Println();

  Println("return value", someFunction(9));

  var Engine = {};
  Engine["Displacement"] = 2.4;
  Engine["Type"] = "v6";
  Engine["HP"] = 235;
  Engine["OilType"] = "5w-20";
  var Warnings_SKiJOQueIf = {};
  Engine["Warnings"] = Warnings_SKiJOQueIf;

  var things = {};
  things["thingy"] = 0;

  var Car = {};
  Car["Type"] = "car";
  Car["New"] = false;
  Car["numOfWheels"] = 4;
  var engine_ynlFZpOWZK = {};
  engine_ynlFZpOWZK["Displacement"] = 2.4;
  engine_ynlFZpOWZK["Type"] = "v6";
  engine_ynlFZpOWZK["HP"] = 199;
  engine_ynlFZpOWZK["OilType"] = "5w-20";
  var Warnings_mYzneOcdqz = {};
  Warnings_mYzneOcdqz["lowOil"] = false;
  Warnings_mYzneOcdqz["lowGas"] = false;
  Warnings_mYzneOcdqz["highTemp"] = false;
  Warnings_mYzneOcdqz["lowTirePressure"] = false;
  Warnings_mYzneOcdqz["antique"] = false;
  engine_ynlFZpOWZK["Warnings"] = Warnings_mYzneOcdqz;

  Car["engine"] = engine_ynlFZpOWZK;

  var oldCar = {};
  oldCar["Type"] = "car";
  oldCar["New"] = false;
  oldCar["numOfWheels"] = 4;
  var engine_KJbwFwgDZI = {};
  engine_KJbwFwgDZI["Displacement"] = 2.4;
  engine_KJbwFwgDZI["Type"] = "v6";
  engine_KJbwFwgDZI["HP"] = 160;
  engine_KJbwFwgDZI["OilType"] = "5w-20";
  var Warnings_hYxbSDxWgx = {};
  engine_KJbwFwgDZI["Warnings"] = Warnings_hYxbSDxWgx;

  oldCar["engine"] = engine_KJbwFwgDZI;

  Println("Default Value for Car struct", Car);
}

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
int main() {

  var thingy = 7;

  Print("thingy =", thingy, "\n");

  Println(thingy);

  Println();

  thingy = 69.69;

  Print("thingy =", thingy, "\n");

  Println(thingy);

  Println();

  thingy = "woah woah woah";

  Print("thingy =", thingy, "\n");

  Println(thingy);

  Println();

  thingy = false;

  Print("thingy =", thingy, "\n");

  Println(thingy);

  Println();

  var thingyObject = {};
  thingyObject["im_just_a"] = "DEAD BOY";

  Println(thingyObject);

  var _mpzoJxReuj = {};
  _mpzoJxReuj["got_no"] = "BLOOD IN MY VEINS";

  Println(_mpzoJxReuj);

  Println();
}

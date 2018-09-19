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

int main() {

  var obj = {};
  obj["something"] = "here";
  var hey_YGCZbWMYbz = {};
  hey_YGCZbWMYbz["me"] = true;
  hey_YGCZbWMYbz["anIntVariable"] = 69;
  obj["hey"] = hey_YGCZbWMYbz;

  var objs[] = {};
  {
    var _irUaBApshi = {};
    _irUaBApshi["another"] = "object";
    objs[0] = _irUaBApshi;
  }
  {
    var obj_YhESdZcLil = {};
    obj_YhESdZcLil["something"] = "here";
    var hey_UkfeWRXNlP = {};
    hey_UkfeWRXNlP["me"] = true;
    hey_UkfeWRXNlP["anIntVariable"] = 69;
    obj_YhESdZcLil["hey"] = hey_UkfeWRXNlP;
    objs[1] = obj_YhESdZcLil;
  }
  {
    var obj_uZNUhvKdsF = {};
    obj_uZNUhvKdsF["something"] = "here";
    var hey_uDsMboBnny = {};
    hey_uDsMboBnny["me"] = true;
    hey_uDsMboBnny["anIntVariable"] = 69;
    obj_uZNUhvKdsF["hey"] = hey_uDsMboBnny;
    objs[2] = obj_uZNUhvKdsF;
  }
}

#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/ExpressRedo/lib/std.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/ExpressRedo/lib/var.cpp"
#include <string>
void declareSomething() { bool i = true; }

int main() {
  var obj = {};
  obj["something"] = "here";
  var hey = {};
  hey["me"] = true;
  hey["anIntVariable"] = 69;
  obj["hey"] = hey;
  var objs[] = {};
  {
    var _SOTdhyFFyv = {};
    _SOTdhyFFyv["another"] = "object";
    objs[0] = _SOTdhyFFyv;
  }
  {
    var obj_YSkjXlpSzV = {};
    obj_YSkjXlpSzV["something"] = "here";
    var hey = {};
    hey["me"] = true;
    hey["anIntVariable"] = 69;
    obj_YSkjXlpSzV["hey"] = hey;
    objs[1] = obj_YSkjXlpSzV;
  }
  {
    var obj_FFXnkfZvAP = {};
    obj_FFXnkfZvAP["something"] = "here";
    var hey = {};
    hey["me"] = true;
    hey["anIntVariable"] = 69;
    obj_FFXnkfZvAP["hey"] = hey;
    objs[2] = obj_FFXnkfZvAP;
  }
}

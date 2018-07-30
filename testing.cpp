#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/std.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/var.cpp"
#include <string>
var someFunction(int arg1) {

  {
    int arrayBoi_1532987607[] = {2, 4, 5, 9};
    int k = 0;
    int k_1532987607 = 0;
    while (k_1532987607 < 4) {
      {
        k = arrayBoi_1532987607[k_1532987607];
        if (arg1 < k) {
          var _KZcehGNfKr = {};
          _KZcehGNfKr["value"] = k;
          return _KZcehGNfKr;
        }
      }
      k_1532987607 += 1;
    }
  }
  var _szchZTqNQg = {};
  _szchZTqNQg["value"] = 1000;
  return _szchZTqNQg;
}

int main() { println(someFunction(5)); }

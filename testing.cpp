#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/std.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/var.cpp"
#include <string>
var someFunction(int arg1) {

  {
    int arrayBoi_1532978563[] = {2, 4, 5, 9};
    int k = 0;
    int k_1532978563 = 0;
    while (k_1532978563 < 4) {
      {
        k = arrayBoi_1532978563[k_1532978563];
        if (arg1 < k) {
          var _BCCsqULJyB = {};
          _BCCsqULJyB["value"] = k;
          return _BCCsqULJyB;
        }
      }
      k_1532978563 += 1;
    }
  }
  var _IJHOuZTrWj = {};
  _IJHOuZTrWj["value"] = 1000;
  return _IJHOuZTrWj;
}

int main() { someFunction(5); }

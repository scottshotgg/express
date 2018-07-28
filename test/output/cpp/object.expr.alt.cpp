#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/ExpressRedo/lib/std.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/ExpressRedo/lib/var.cpp"
#include <string>
void declareSomething() { bool i = true; }

int main() {
  var obj = map<string, var>{
      {"something", "here"},
      {"hey",
       map<string, var>{
           {"me", true},
           {"anIntVariable", 69},
       }},
  };



  int f = 0;

  {
    int arrayBoi_1532791625[] = {9, 8, 7};
    int i = 0;
    int i_1532791625 = 0;
    while (i_1532791625 < 3) {
      {
        i = arrayBoi_1532791625[i_1532791625];
        f = i;
      }
      i_1532791625 += 1;
    }
  }
  map<string, var> objs[] = {map<string, var>{
                                 {"another", "object"},
                             },
                             map<string, var>{
                                 {"something", "here"},
                                 {"hey",
                                  map<string, var>{
                                      {"me", true},
                                      {"anIntVariable", 69},
                                  }},
                             },
                             map<string, var>{
                                 {"something", "here"},
                                 {"hey",
                                  map<string, var>{
                                      {"me", true},
                                      {"anIntVariable", 69},
                                  }},
                             }};
}

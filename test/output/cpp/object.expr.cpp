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

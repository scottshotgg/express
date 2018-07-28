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
                                  }},
                             },
                             map<string, var>{
                                 {"something", "here"},
                                 {"hey",
                                  map<string, var>{
                                      {"me", true},
                                  }},
                             }};

    obj["something"] = true;
    objs[1]["hey"]["another_key_to_use"] = 666;
    cout << "objs1 " << objs[1] << endl;
    cout << "objs2 " << objs[2] << endl;
}

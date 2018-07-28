#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/ExpressRedo/lib/std.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/ExpressRedo/lib/var.cpp"
#include <string>
void declareSomething() { bool i = true; }

int main() {
  map<string, var> obj = map<string, var>{
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

    cout << objs[0] << endl;
    objs[0]["woah"] = true;
    cout << objs[0] << endl;

    cout << objs[1] << endl;
    cout << objs[2] << endl;
    objs[2]["something"] = "there";
    cout << objs[1] << endl;
    cout << objs[2] << endl;
}

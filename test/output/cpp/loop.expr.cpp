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

  int f = 0;

  {
    int i = 1;
    while (i < 10) {
      {
        defer onLeaveFuncs;

        f = i;

        int h = 1;
      }
      i += 1;
    }
  }

  {
    int arrayBoi_1537322069[] = {1, 2, 4};
    int i = 0;
    int i_1537322069 = 0;
    while (i_1537322069 < 3) {
      {
        defer onLeaveFuncs;

        i = i_1537322069;

        f = i;

        int h = 1;
      }
      i_1537322069 += 1;
    }
  }

  int countdown[] = {9, 8, 7, 5, 4, 3, 2, 1};

  {
    int i = 0;
    int i_1537322069 = 0;
    while (i_1537322069 < 8) {
      {
        defer onLeaveFuncs;

        i = countdown[i_1537322069];

        f = i;

        int h = 1;
      }
      i_1537322069 += 1;
    }
  }
}

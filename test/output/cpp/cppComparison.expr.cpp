#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/std.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/var.cpp"
#include <string>

int main() {
  int total = 0;

  {
    int i = 0;
    while (i < 1000000) {
      {
        total = 0;
      }
      i += 1;
    }
  }
}

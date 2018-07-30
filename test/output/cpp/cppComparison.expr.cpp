#include "/home/scottshotgg/Development/go/src/github.com/scottshotgg/express/lib/std.cpp"
#include "/home/scottshotgg/Development/go/src/github.com/scottshotgg/express/lib/var.cpp"
#include <string>

int main() {
  int total = 0;

  {
    int i = 0;
    while (i < 1000000000) {
      {
        total += i;
      }
      i += 1;
    }
  }
}

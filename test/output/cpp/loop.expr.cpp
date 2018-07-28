#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/ExpressRedo/lib/std.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/ExpressRedo/lib/var.cpp"
#include <string>
void declareSomething() { bool i = true; }

int main() {
  int f = 0;

  {
    int i = 1;
    while (i < 10) {
      {
        f = i;
        int h = 1;
      }
      i += 1;
    }
  }

  {
    int arrayBoi_1532747125[] = {1, 2, 4};
    int i = 0;
    int i_1532747125 = 0;
    while (i_1532747125 < 3) {
      {
        i = i_1532747125;
        f = i;
        int h = 1;
      }
      i_1532747125 += 1;
    }
  }
  int countdown[] = {9, 8, 7, 5, 4, 3, 2, 1};

  {
    int i = 0;
    int i_1532747125 = 0;
    while (i_1532747125 < 8) {
      {
        i = countdown[i_1532747125];
        f = i;
        int h = 1;
      }
      i_1532747125 += 1;
    }
  }
}

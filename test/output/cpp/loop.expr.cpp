#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/std.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/var.cpp"
#include <string>
int increment(int i) { return 100; }

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
    int arrayBoi_1532850496[] = {1, 2, 4};
    int i = 0;
    int i_1532850496 = 0;
    while (i_1532850496 < 3) {
      {
        i = i_1532850496;
        f = i;
        int h = 1;
      }
      i_1532850496 += 1;
    }
  }
  int countdown[] = {9, 8, 7, 5, 4, 3, 2, 1};

  {
    int i = 0;
    int i_1532850496 = 0;
    while (i_1532850496 < 8) {
      {
        i = countdown[i_1532850496];
        f = i;
        int h = 1;
      }
      i_1532850496 += 1;
    }
  }
}

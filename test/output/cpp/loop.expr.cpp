#include "var.cpp"
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
    int arrayBoi_1532455693[] = {1, 2, 4};
    int i = 0;
    int i_1532455693 = 0;
    while (i_1532455693 < 3) {
      {
        i = i_1532455693;
        f = i;
        int h = 1;
      }
      i_1532455693 += 1;
    }
  }
  int countdown[] = {9, 8, 7, 5, 4, 3, 2, 1};

  {
    int i = 0;
    int i_1532455693 = 0;
    while (i_1532455693 < 8) {
      {
        i = countdown[i_1532455693];
        f = i;
        int h = 1;
      }
      i_1532455693 += 1;
    }
  }
}

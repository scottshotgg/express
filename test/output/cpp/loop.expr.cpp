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
    int arrayBoi_1532545221[] = {1, 2, 4};
    int i = 0;
    int i_1532545221 = 0;
    while (i_1532545221 < 3) {
      {
        i = i_1532545221;
        f = i;
        int h = 1;
      }
      i_1532545221 += 1;
    }
  }
  int countdown[] = {9, 8, 7, 5, 4, 3, 2, 1};

  {
    int i = 0;
    int i_1532545221 = 0;
    while (i_1532545221 < 8) {
      {
        i = countdown[i_1532545221];
        f = i;
        int h = 1;
      }
      i_1532545221 += 1;
    }
  }
}

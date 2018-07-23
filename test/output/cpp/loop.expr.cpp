#include "json.hpp"
#include <string>
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
    int arrayBoi_1532324567[] = {1, 2, 4};
    int i = 0;
    int i_1532324567 = 0;
    while (i_1532324567 < 3) {
      {
        i = i_1532324567;
        f = i;
        int h = 1;
      }
      i_1532324567 += 1;
    }
  }
  int countdown[] = {9, 8, 7, 5, 4, 3, 2, 1};

  {
    int i = 0;
    int i_1532324567 = 0;
    while (i_1532324567 < 8) {
      {
        i = countdown[i_1532324567];
        f = i;
        int h = 1;
      }
      i_1532324567 += 1;
    }
  }
}

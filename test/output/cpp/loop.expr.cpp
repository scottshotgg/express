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
    int i = 0;
    int arrayBoi[] = {9, 8, 7, 5, 4, 3, 2, 1};
    int i_1532320836 = 0;
    while (i_1532320836 < 8) {
      {
        i = i_1532320836;
        f = i;
        int h = 1;
      }
      i_1532320836 += 1;
    }
  }

  {
    int i = 0;
    int arrayBoi[] = {9, 8, 7, 5, 4, 3, 2, 1};
    int i_1532320836 = 0;
    while (i_1532320836 < 8) {
      {
        i = arrayBoi[i_1532320836];
        f = i;
        int h = 1;
      }
      i_1532320836 += 1;
    }
  }
}

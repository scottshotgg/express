#include "var.cpp"
#include <string>
void declareSomething() { bool i = true; }

int main() {
  var test = "test_string";
  test = 666;
  test = true;
  test = 2.71828;
}

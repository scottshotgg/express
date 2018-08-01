#include <iostream>

// FIXME: could probably speed this up by using a stringstream and or a different way to print

// Base-case print function
template <typename T>
void print(T data) {
  std::cout << data;
}

void print(bool data) {
  if (data) {
    std::cout << "true";
  } else {
    std::cout << "false";
  }
}

// Recursive print macro
template <typename T, typename... Args>
void print(T first, Args... args) {
  print(first);
  std::cout << " ";
  print(args...);
}

// Print for arrays
// TODO: this might be moved to the `var` lib if vars are allowed to hold arrays
template <typename T>
void print(T rest[]) {
  std::cout << "[ ";
  int length = (sizeof(*rest)/sizeof(*rest)) + 1;
  for (int i = 0; i < length; i++) {
    print(rest[i]);
    std::cout << ", ";
  }
  print(rest[length]);
  std::cout << " ]";
}

// Recursive println macro
template <typename T, typename... Args>
void println(T first, Args... args) {
  print(first, args...);
  std::cout << std::endl;
}


#include <iostream>

// Empty Print function
void Print() {}

// Print function just for C strings
void Print(char* data) {
  std::cout << data;
}

// Print function to change bool types as actual "true" or "false" literal string values
void Print(bool data) {
  if (data) {
    std::cout << "true";
  } else {
    std::cout << "false";
  }
}

// Print function for arrays
// TODO: this might be moved to the `var` lib if vars are allowed to hold arrays
template <typename T>
void Print(T rest[]) {
  std::cout << "[ ";
  int length = (sizeof(*rest)/sizeof(*rest)) + 1;
  for (int i = 0; i < length; i++) {
    Print(rest[i]);
    std::cout << ", ";
  }
  Print(rest[length]);
  std::cout << " ]";
}

// Base-case Print function
template <typename T>
void Print(T data) {
  std::cout << data;
}

// Recursive Print function
template <typename T, typename... Args>
void Print(T first, Args... args) {
    Print(first);
    std::cout << " ";
    Print(args...);
}

// Empty Println macro
void Println() {
  std::cout << std::endl;
}

// Recursive Println macro
template <typename T, typename... Args>
void Println(T first, Args... args) {
  Print(first, args...);
  std::cout << std::endl;
}
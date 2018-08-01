#include <iostream>

// FIXME: could probably speed this up by using a stringstream and or a different way to print

// Empty print function
void print() {
  //cout << "im in here print empty" << endl;
}

void print(char* data) {
  //cout << "im in here char*" << endl;
  std::cout << data;
}

void print(bool data) {
  //cout << "im in here print bool" << endl;
  if (data) {
    std::cout << "true";
  } else {
    std::cout << "false";
  }
}

// template<> void print()

// Base-case print function
template <typename T>
void print(T data) {
  //cout << "im in here print T" << endl;
  std::cout << data;
}

// template <> 
// void print(std::string data){
//   cout << "im in here string" << endl;
//   std::cout << data;
// }

// template <> 
// void print(char* data){
//   cout << "im in here char*" << endl;
//   std::cout << data;
// }

// Recursive print macro
template <typename T, typename... Args>
void print(T first, Args... args) {
  //cout << "im in here print" << endl;
  // if (sizeof...(Args) - 1) {
  //cout << "im in here print" << endl;
    print(first);
    std::cout << " ";
    print(args...);
  // } else {
  // //cout << "im in here print" << endl;
  //   print(first);
  // }
}

// Print for arrays
// TODO: this might be moved to the `var` lib if vars are allowed to hold arrays
template <typename T>
void print(T rest[]) {
  //cout << "im in here rest" << endl;
  std::cout << "[ ";
  int length = (sizeof(*rest)/sizeof(*rest)) + 1;
  for (int i = 0; i < length; i++) {
    print(rest[i]);
    std::cout << ", ";
  }
  print(rest[length]);
  std::cout << " ]";
}

// template <> 
// void print(std::string data){

// };

// Recursive println macro
template <typename T, typename... Args>
void println(T first, Args... args) {
  //cout << "im in here println" << endl;
  print(first, args...);
  std::cout << std::endl;
}

// Recursive println macro
void println() {
  //cout << "im in here println empty" << endl;
  std::cout << std::endl;
}

// // Templated defer function
// // template <typename T, typename T1, typename T2> -- // if we want void to return something
// template <typename T, typename T1>
// void defer(T1 (*f)(T2)) {
//   deferFunc hey_its_me(nullptr, [](...){ T1() });
// }
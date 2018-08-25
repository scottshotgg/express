#include <iostream>
#include <array>

// FIXME: could probably speed this up by using a stringstream and or a different way to print

// Empty print function
void Print() {
  //cout << "im in here print empty" << endl;
}

void Print(std::string data) {
  //cout << "im in here char*" << endl;
  std::cout << data;
}

void Print(bool data) {
  //cout << "im in here print bool" << endl;
  if (data) {
    std::cout << "true";
  } else {
    std::cout << "false";
  }
}

// template<> void Print()

// Base-case print function
template <typename T>
void Print(T data) {
  //cout << "im in here print T" << endl;
  std::cout << data;
}

// template <> 
// void Print(std::string data){
//   cout << "im in here string" << endl;
//   std::cout << data;
// }

// template <> 
// void Print(char* data){
//   cout << "im in here char*" << endl;
//   std::cout << data;
// }

// Recursive print macro
template <typename T, typename... Args>
void Print(T first, Args... args) {
  //cout << "im in here print" << endl;
  // if (sizeof...(Args) - 1) {
  //cout << "im in here print" << endl;
    Print(first);
    std::cout << " ";
    Print(args...);
  // } else {
  // //cout << "im in here print" << endl;
  //   Print(first);
  // }
}

// Print for arrays
// TODO: this might be moved to the `var` lib if vars are allowed to hold arrays
template <typename T>
void Print(T rest[]) {
  //cout << "im in here rest" << endl;
  std::cout << "[ ";
  int length = sizeof(rest[0]) + 2;
  for (int i = 0; i < length; i++) {
    Print(rest[i]);
    std::cout << ", ";
  }
  Print(rest[length]);
  std::cout << " ]";
}

// template <> 
// void Print(std::string data){

// };1

// Recursive println macro
template <typename T, typename... Args>
void Println(T first, Args... args) {
  //cout << "im in here println" << endl;
  Print(first, args...);
  std::cout << std::endl;
}

template <typename T>
void Println(T rest[]) {
  Print(rest);
  std::cout << std::endl;
}

// Recursive println macro
void Println() {
  //cout << "im in here println empty" << endl;
  std::cout << std::endl;
}

// // Templated defer function
// // template <typename T, typename T1, typename T2> -- // if we want void to return something
// template <typename T, typename T1>
// void defer(T1 (*f)(T2)) {
//   deferFunc hey_its_me(nullptr, [](...){ T1() });
// }
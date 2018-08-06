#include <string>
#include <iostream>

using namespace std;

class File {
  public:
    // std::string Open(string filepath);
    std::string Read(int numChars);
    std::string ReadLine();
    File(FILE* fp) : file(fp) {}

  private:
    int readPointer = 0;
    FILE* file = nullptr;
};

// string file::Read() {

// }

File Open(string filepath, string mode) {
  return File{ fopen(filepath.c_str(), mode.c_str()) };
}

// Try to do something like this
// string File::ReadLine() {
//   string entireLine;
//   int buffAmount = 10;
//   char readBuff[buffAmount];
//   while(fgets(readBuff, buffAmount, file)) {
//     cout << "\"" << readBuff << "\"" << endl;
//     if (readBuff[9]) {
//       return "something";
//     }

//     cout << "last char" << "\"" << readBuff[9] << "\"" << endl;
//   }

//   return entireLine;
// }

string ReadFile(string filepath) {
  FILE* file = fopen(filepath.c_str(), "r");
  if (file == nullptr) {
    return "";
  }
  
  string entireFile;
  int buffAmount = 100;
  char readBuff[buffAmount];
  while(fgets(readBuff, buffAmount, file)) {
      entireFile += readBuff;
  }

  return entireFile;
}

void WriteFile(string filepath, string contents, bool writeOver) {
  FILE* file = fopen(filepath.c_str(), "w");
  if (file != nullptr) {
    if (!writeOver) {
      // TODO: need to return an error or
      // something here
    }
  }

  fputs(contents.c_str(), file);
}
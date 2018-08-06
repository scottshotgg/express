#include <string>
#include <iostream>

using namespace std;

class File {
  public:
    void Close();

    std::string Read(int);
    std::string ReadLine();

    void Write(string);
    void WriteLine(string);

    File() {}
    File(FILE* fp) : file(fp) {}

  private:
    FILE* file = nullptr;
    int readPointer = 0;
    
    // Assume that files are ascii encoded by default
    string _charset = "ascii";
    char _lastChar = 0;
    char _currentChar = 0;
    char _nextChar = 0;
};

string File::Read(int numOfChars) {
  // string entireFile;
  // char readBuff[numOfChars];
  // fgets(readBuff, numOfChars, this->file);

  // return readBuff;

  string entireLine;
  int c = fgetc(this->file);
  int count = 0;
  
  while(c != EOF && count < numOfChars) {
    count++;
    entireLine += c;
    c = fgetc(this->file);
  }

  return entireLine;
}

string File::ReadLine() {
  string entireLine;
  int c = fgetc(this->file);

  while(c != EOF && c != '\n') {
    entireLine += c;
    c = fgetc(this->file);
  }

  return entireLine;
}

void File::Write(string text) {
  // do something with the status?
  fputs(text.c_str(), this->file);
}

void File::WriteLine(string text) {
  Write(text + "\n");
}

// Need to put a status here
void File::Close() { fclose(this->file); }

File Open(string filepath, string mode) {
  return File {
    fopen(filepath.c_str(),
    mode.c_str())
  };
}

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
  fclose(file);

  return entireFile;
}

void WriteFile(string filepath, string contents, bool overwrite) {
  FILE* file = fopen(filepath.c_str(), "w");
  if (file != nullptr) {
    if (!overwrite) {
      // TODO: need to return an error or
      // something here
    }
  }

  fputs(contents.c_str(), file);
  fclose(file);
}
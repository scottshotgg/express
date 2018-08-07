#include <string>
#include <iostream>
#include <libgen.h>

using namespace std;

class File {
  public:
    // Constructors
    File() {}
    File(FILE* fp, string mode) : file(fp), mode(mode) {}
    File(FILE* fp, string mode, string filename, string filepath, bool open) :
      file(fp), mode(mode), filename(filename), filepath(filepath), open(open) {}

    // TODO: Need to make destructors

    // Getters
    bool IsOpen();

    // Close file pointer
    void Close();

    // Utility functions
    bool AtEOF();
    bool AtEOL();
    int  CurrentPosition();
    int  SeekToPosition(int);
    void ResetPosition();
    int  Length();
    int  Rename(string);
    int  Move(string);
    int  Delete(string);

    // Read functions
    std::string Read(int);
    std::string ReadNextChar();
    std::string ReadUntilNext(char);
    std::string ReadLine();

    // Write functions
    void Write(string);
    void WriteLine(string);

  private:
    FILE* file = nullptr;
    string mode = "r";
    string filename;
    string filepath;
    bool open = false;

    // All variables preceded with _ are experimental
    bool _eof = false;
    bool _eol = false;

    // Assume that files are ascii encoded by default
    string _charset = "ascii";
};

bool File::IsOpen() {
  return open;
}

// TODO: Need to put a status here
void File::Close() {
  open = false;
  fclose(this->file);
}

bool File::AtEOF() {
  return feof(this->file);
}

bool File::AtEOL() {
  bool eol = false;
  int c = fgetc(this->file);

  if (c != EOF && c == '\n') {
    eol = true;
  }

  fseek(this->file, 1, SEEK_CUR);
  return eol;
}

int File::CurrentPosition() {
  return ftell(this->file);
}

int File::SeekToPosition(int position) {
  return fseek(this->file, position, SEEK_SET);
}

void File::ResetPosition() {
  return rewind(this->file);
  // return fseek(this->file, 0, SEEK_SET);
}

// TODO: this could probably be something
// that we do when the file is opened
int File::Length() {
  int curr = CurrentPosition();

  fseek(this->file, 0, SEEK_END);
  int sz = CurrentPosition();

  SeekToPosition(curr);

  return sz;
}

int File::Rename(string newFilename) {
  return Move(filepath+"/"+newFilename);
}

int File::Move(string newFilepath) {
  return rename((filepath+"/"+filename).c_str(), (newFilepath).c_str());
}

int File::Delete(string newFilepath) {
  return remove((filepath+"/"+filename).c_str());
}

// int File::ChangeName(string newFilename) {
//   return Move(filepath+"/"+newFilename);
// }

// int File::ChangeLocation(string newFilepath) {
//   return rename((filepath+"/"+filename).c_str(), (newFilepath).c_str());
// }

// int File::Delete(string newFilepath) {
//   return remove((filepath+"/"+filename).c_str());
// }

string File::Read(int numOfChars) {
  string entireLine;
  int c = fgetc(this->file);
  int count = 0;
 
  while(!feof(this->file) && count < numOfChars) {
    count++;
    entireLine += c;
    c = fgetc(this->file);
  }

  return entireLine;
}

string File::ReadNextChar() {
  // TODO: is this safe?
  if (!feof(this->file)) {
    int c = fgetc(this->file);
    return (char*)&c;
  }

  return "";
}

// TODO: try using stdio.h::getdelim
// TODO: this needs to also return a bool/error indicating
//  whether or not it found the char
// Maybe return nothing and move the seek pointer back if nothing found
string File::ReadUntilNext(char lookingFor) {
  string entireLine;
  entireLine += fgetc(this->file);
  int c = fgetc(this->file);

  while(!feof(this->file)) {
    if (c != lookingFor) {
      entireLine += c;
      c = fgetc(this->file);
      continue;
    } else {
      // fseek(this->file, 1, SEEK_CUR);
      ungetc(c, this->file);
    }

    break;
  }

  return entireLine;
}

// TODO: try using stdio.h::getline
string File::ReadLine() {
  string entireLine;
  int c = fgetc(this->file);

  while(!feof(this->file) && c != '\n') {
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

// FIXME: All functions below this will later be moved to their own package
// Some of these operations will benefit from directly using the file stream
File Open(string filepath, string mode) {
  char* filepathCString = (char*)filepath.c_str();

  return File {
    fopen(filepathCString, mode.c_str()),
    mode.c_str(),
    basename(filepathCString),
    dirname(filepathCString),
    true
  };
}

string ReadFile(string filepath) {
  FILE* file = fopen(filepath.c_str(), "r");
  if (file == nullptr) {
    fclose(file);
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
      fclose(file);
    }
  }

  fputs(contents.c_str(), file);
  fclose(file);
}

// TODO: this needs to return a status/error/something
void CopyFile(string fromPath, string toPath, bool overwrite) {
  WriteFile(toPath, ReadFile(fromPath), overwrite);
}
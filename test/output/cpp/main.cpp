#include "std.cpp"


void TestPrint() {
// ten million ops
  int varOperationsAmount = 1000;

  auto t1 = std::chrono::high_resolution_clock::now();

  for (int i = 0; i < varOperationsAmount; i++) {
    print("WWW.FIRMCODES.COM");
  }
  // std::cout << "vz " << vz << std::endl;

  int varDuration = std::chrono::duration_cast<std::chrono::milliseconds>(
                        std::chrono::high_resolution_clock::now() - t1)
                        .count();

  std::cout << std::endl << "var type operations took " << varDuration << " milliseconds\n\n";
}

int main() {
  int lilIntBoi = 42;
  std::string lilstringBoi = "ayy girl waddup";
  bool lilBoolBoi = true;
  float lilFloatBoi = .1234567890;

  println("WWW.FIRMCODES.COM");
  println("WWW.FIRMCODES.COM", lilIntBoi);
  println("WWW.FIRMCODES.COM", lilstringBoi);
  println("WWW.FIRMCODES.COM", lilBoolBoi);
  println("WWW.FIRMCODES.COM", lilFloatBoi);
  println("WWW.FIRMCODES.COM", lilIntBoi, lilstringBoi, lilBoolBoi, lilFloatBoi);
}
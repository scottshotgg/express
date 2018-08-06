#include <iostream>
#include <stack>
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/std.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/var.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/file.cpp"

// class defer {
//   public:
//     std::stack <std::function<void()>> deferStack;
    
//     ~defer() {
//       while (!deferStack.empty()) {
//         (deferStack.top())();
//         deferStack.pop();
//       }
//     }
// };

int main() {
    // TODO: need to test if ReadLine reads out the newline in golang
    cout << ReadFile("../defer.cpp") << endl;
    WriteFile("../defer_copy.cpp", ReadFile("../defer.cpp"), false);

    File readFile = Open("../defer.cpp", "rw");

    File writeFile = Open("defer_copy.cpp", "w");
    writeFile.WriteLine(readFile.ReadLine());
    writeFile.Close();
    readFile.Close();
}
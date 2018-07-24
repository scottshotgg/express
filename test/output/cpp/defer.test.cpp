#include <memory>
#include <iostream>

using namespace std;
using defer = shared_ptr<void>;

void something() {
  cout << "something" << endl;
}

int main() {
  cout << "start" << endl;
  // Express needs to take care of putting these in reverse order
  // and assigning random names to the functions

  // TODO: need to see about implementing function arguments

  // Using a function that is already declared
  defer ____(nullptr, [](...){ something(); });

  // Anonymous statements
  defer ___(nullptr, [](...){ cout << "hi my name is scott" << endl; });
  defer _(nullptr, [](...){ cout << "end" << endl; });
}
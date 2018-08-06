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

    // cout << ReadFile("../defer.cpp") << endl;
    // WriteFile("../defer_copy.cpp", ReadFile("../defer.cpp"), false);

    File file = Open("../defer.cpp", "rw");
    cout << file.ReadLine() << endl;
    cout << file.ReadLine() << endl;
    cout << file.ReadLine() << endl;
//   defer deferFuncs;
  

//   std::cout << "start" << std::endl;
//   // Express needs to take care of putting these in reverse order
//   // and assigning random names to the functions

//   // TODO: need to see about implementing function arguments
//   deferFuncs.deferStack.push(([=](...){ println("hi", ""); }));

//   for (int i = 0; i < 10; i++) {
//     println("i:", i);
//     deferFuncs.deferStack.push(([=](...){ println(i); /* println(i); */ }));
//   }
}

/*
#include <functional>
#include <iostream>
#include <string>
#include <vector>
 
using namespace std;
 
void execute(const vector<function<void ()>>& fs)
{
    for (auto& f : fs)
        f();
}
 
void plain_old_func()
{
    cout << "I'm an old plain function" << endl;
}
 
class functor
{
    public:
        void operator()() const
        {
            cout << "I'm a functor" << endl;
        }
};
 
int main()
{
    vector<function<void ()>> x;
    x.push_back(plain_old_func);
     
    functor functor_instance;
    x.push_back(functor_instance);
    x.push_back([] ()
    {
        cout << "HI, I'm a lambda expression" << endl;
    });
     
    execute(x);
}
*/
#include <functional>
#include <stack>

class defer {
  public:
    std::stack <std::function<void()>> deferStack;

    ~defer() {
      while (!deferStack.empty()) {
        (deferStack.top())();
        deferStack.pop();
      }
    }
};
// Code formatting not generated by transpiler; included from VS Code for visibility.
// See main.expr.shit for transpiled version

#include <map>
#include <string>
struct Any
{
  std::string type;
  void *data;
};

int main()
{
  int f = 0;
  
  {
    int i = 1;
    while (i < 10)
    {
      {
        f = i;
        int h = 1;
        // Printf's not generated by transpiler; added for testing and visibility
        printf("Inside -\n");
        printf("f: %d\n", f);
        printf("h: %d\n", h);
      }
      i += 1;
    }
  }

  // Printf's not generated by transpiler; added for testing and visibility
  printf("\nOutside -\n");
  printf("f: %d\n", f);
}
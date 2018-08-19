#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/defer.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/file.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/std.cpp"
#include "/Users/sgg7269/Development/go/src/github.com/scottshotgg/express/lib/var.cpp"
#include <string>
#include <sstream>

// // template <typename T>
// tuple<string, T> tuple_list;

// // typedef vector<tuple<int,string>> tuple_list;
// typedef std::vector<tuple_list> props;

template<class K, class T>
std::map<K, T> dict(K key, T data) {
   std::map<K,T> map;
   map[key] = data;
   return map;
}


defer onExitFuncs;

// FIXME: could use a list of template tuples and a map of { <name> : <pointer> }
struct Thing {
  int fieldA = 990;
  string stringField;
  bool false_field;
  float anotherFielderino;

  string ToString() {
    std::stringstream ss;
    ss << "{ fieldA: " << this->fieldA << ", "
      << "stringField: \"" << this->stringField << "\", "
      << "false_field: " << this->false_field << ", "
      << "anotherFielderino: " << this->anotherFielderino << " }";

    return ss.str();
  }
};

template <typename T1, typename T2, typename T3, typename T4>
struct ThingT {
  T1 fieldA = 990;
  T2 stringField;
  T3 false_field;
  T4 anotherFielderino;
};

struct Thing2 {
  // This is how we'll do default values for the struct
  int stuff = 666;
  Thing ting;

  // string ToString() {
  //   stringstream ss;
  //   ss << this;
  //   string thing = ss.str();

  //   return thing;
  // }
};

std::ostream& operator<< (std::ostream &o, const Thing &t) {
  Thing tc = (Thing)t;

  return o << tc.ToString();
}

// std::ostream& operator<< (std::ostream &o, const Thing2 &t) {
//   return o
//       << "{ stuff: " << t.stuff << ", "
//       << "ting: " << t.ting << " }";
// }

int printStuff(int k) {
  defer onReturnFuncs;
  {
    defer onLeaveFuncs;

    {
      int i = 0;
      while (i < k) {
        {
          defer onLeaveFuncs;
          onExitFuncs.deferStack.push([=](...) { Println("on exit", i); });
          onReturnFuncs.deferStack.push([=](...) { Println("on return", i); });
          onLeaveFuncs.deferStack.push([=](...) { Println("on leave", i); });
          onReturnFuncs.deferStack.push([=](...) { Println("defer", i); });
        }
        i += 1;
      }
    }
    return 0;
  }
}

var increment(var i) {
  defer onReturnFuncs;
  {
    defer onLeaveFuncs;
    var _PxkNpMspyl = {};
    _PxkNpMspyl["something"] = "else";
    return _PxkNpMspyl;
  }
}

int main() {
  map thing = dict("me", 69);
  // cout <<  << endl;

  Thing something = {
    69,
    "yea its meh",
    true,
    66.6
  };

  struct mabob {
    string hey = "itsme";
  };

  mabob er;

  Println(something);
  // something.fieldA = "hell yeah its me";
  // Println(something);
  // something.fieldA = {};
  // something.fieldA["hellyeah"] = "fuckin rite";
  // Println(something);
  // Println(Thing2{
  //   667,
  //   something
  // });

  File file = Open("heyitsme.txt", "w");
  cout << something.ToString() << endl;
  file.WriteLine(something.ToString());
  file.Close();
}

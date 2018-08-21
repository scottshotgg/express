```
interface myInterface = {
    int A
    float B
    func GetA() int
    func GetB() float
}

interface myInheritedInterface = myInterface {
    bool  C
    string D
    func GetC() bool
    func GetD() string
}

// This will accept either myInterface or MyInheritedInterface
func someFunction(myInterface something) {}
```

> Interfaces can be enumerated (defined) from objects or structs at compile time, and only from structs at run time. To call an interface (read: structurally-typed) function using an object at runtime, you must explicitly assert (cast) the object to a struct type that fulfills the interface. In order to maintain the static type checking at compile time so the Express compiler can be certain in its ability to predict runtime results, dynamic variables need to be asserted so that the interface will be satisfied.

> or we could make the interface take an object and any fields that the object does not explicitly provide will be filled using default values and any function will return the default values... although this kinda breaks the meaning of interface, however, this is essentially what casting your object to a struct would do....


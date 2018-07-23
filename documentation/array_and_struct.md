```csharp
cufloat128s filter_coefficients = [3.3, 4.4, 5.5, 6.6]
constant unsigned float128 array coefficients = [3.3, 4.4, 5.5, 6.6]
constant unsigned float128[] coefficients = [3.3, 4.4, 5.5, 6.6]


struct thing = {
    // These will take the type
    // defined default values
    int    normieInt,
    bool   normieBool,
    float  normieFloat,
    string normieString `json:"normie_string"`,
    
    
    // These will be initialized to
    // these default values
    string field_A = "A",
    field1 `json:"field_one"` = 1
    isField := true
    objField `json:"an_object"`: {
        anotherField: "anotherField"
        2ndFIELD: 2;
        fields: [ 1, 2, 3, 4, 5 ],
    }
    "stringedField": "VALUE",
    int fieldReference: .field1,
    int* shadowField = &field1
    fLoAtInG|FiElDs: coefficients;
}
structTest := thing{}

// Doesn't work because this it's an object; not a struct.
objThing := {}
objTest := objThing{}

// Iterable composition
something := thing over things {
    thing + 2
}
```

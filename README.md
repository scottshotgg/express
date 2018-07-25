# Express

Express is a _strongly-typed_ language containing both _type_inference_ and  dynamically typed `var` (think JavaScript, except better) containing an _extremely_ lite runtime that currently does dynamic typing and RAII lifetime management. In the future it will include a greenthread scheduler a la Golang, atomization of operations, garbage collection, SQLite3/GraphQL, and a few other ideas that are currently only conceptual. Most features will be optional and able to be enabled and disabled at compile time and possible runtime.

All programs are currently _transpiled_ to C++ and then LLVM is subsequently invoked (along with `clang-format`) to produce the corresponding binary. At this time, transpiling is, _time-wise_, sufficiently more efficient than outputting LLVM tokens or building an intermediary using SSA/3AC through something like BinaryNinja. Later on, this will most likely be changed in favor  LLVM token production when features either become too much of a burden to implement and maintain in C++ or the transpiler development lags too much to adequently support the development of the language.

The lexer, syntax parser, semantic parser, and C++ transpiler are currently all implemented in `Go` and will later be converted to `Rust`, but a `JavaScript` implementation in Node is also being developed and will later on be consolidated into this repo after a reorganization of the file structure.
<br>
It is currently located at https://github.com/Swivelgames/Express/tree/alt/node

<br>

## Features Implemented 
[~] means it may not be implemented

- [x] Basic types (`int`, `bool`, `float`, `string`)
- [-] `array` type
    - [ ] `int[]`
    - [ ] `float[]`
    - [x] `bool[]`
    - [x] `string[]`
    - [ ] `var[]`
    - [ ] `object[]`
    - [ ] `<struct>[]`
    - [ ] `array[]`
    - [ ] `function[]`
    - [x] Type inference
- [-] `var` types
    - [x] Basic type encapsulation (`int`, `bool`, `float`, `string`)
    - [ ] `object` type encapsulation
    - [ ] `<struct>` type encapsulation
    - [~] `array` type encapsulation
        - `Not sure if I want a singular value to be able to hold multiple values`
    - [~] `function` type encapsulation
        - `A var holding a function doesn't make a lot of sense right now`
- [ ] `struct` type
    - [ ] Tags
- [ ] `function` type
- [ ] Type modifier keywords
    - [ ] `array` postfix
    - [ ] `unsigned` prefix
    - [ ] `constant` prefix
- [ ] Type modifiers
    - [ ] `s` postfix
    - [ ] `u` prefix
    - [ ] `c` prefix
- [-] Function usage
    - [-] Function declaration
        - [x] No args and no returns
        - [ ] Args without returns
        - [ ] Returns without args
        - [ ] Args and returns
    - [ ] Function call
- [x] Blocks
- [-] Access types
    - [x] `private`
    - [ ] `public`
    - [ ] `fileprivate`
<br>

## Stages Implemented
- [x] Lexer
- [x] Syntax
- [x] Semantics
- [x] C++ Transpiler with `clang-format`
- [x] Binary

<br>
<br>

## How To Contribute

### Proposals

Proposals are the first step to contributing your ideas and work into Express. The first step to contributing is submitting a `feature` proposal through a PR. Below contains the list of steps and criteria for submission and approval.
<br>
<br>

- `Features`

    _Features_ should be submitted using a markdown file (`title_of_your_feature.md`) describing the new abilities or features into an `appropriate_labeled` folder within the `proposals` root folder. Proposals need to use fenced code blocks with the `cs` highlighting tag and contain an Express program example.
<br>
<br>

- `Ideas:`
    <br>
    _Ideas_ about the languages direction, internals related to the architecture, or general core-langaage level imlementations should go into the `proposals/ideas` folder and should be formatted as a whitepaper with detail containing a beginning narative that will explain <u>where</u> and <u>how</u> the idea came from, atleast two supporting arguments that should answer the question of <u>why</u> your proposal should be considered, and one informal/abstract (i.e, pseudocode) use-case on <u>how</u> it would be used after implementation.
<br>
<br>

- `Implementations and Experiments`
    <br>
    _Implementations_ and _Experiments_ should be done in a separate branch off of the repo. When you are ready to submit your implementation, you should merge your code locally with the latest `master` branch pull and then submit a PR to `master`.
    <br>
    Your PR should also contain an Express test program in the `test/output/programs` folder, along with the respective `lex`, `syn`, `sem`, `cpp`, and `bin` files.
<br>
<br>

>_`Note`_: If your PR **does not** meet the criteria when it is submitted or cannot be merged it will be _automatically_ denied by the CI/CD pipeline with comments about failure. You will need to submit another one after fixing the problems.

<br>
<br>

## Example Program

Below is currently the most advanced program that can be written in Express. You can find the **full** *uncommented* version labeled `advanced.expr` under the `samples` directory.
<br>
You will find a few examples of the allowed flexibilities and optional verbosity that allow the language to be so _Expressive_.

<br>

```html
Let's Begin:
```

____

<br>

> Start off by declaring some variables:

```csharp
int powerlevel = 0x2327;
bool over9k = false;
float pi = 3.14159265359;
string arizona = "iced out boys";
```

<br>

> Beyond basic static types, Express also supports using dynamically typed variables as well. 
<br>
> It is important to note that these variable dynamically typed at _run-time_ and thus will incur a performance penalty in constrast to static variables.
<br>
> To see how the runtime that manages this in C++, open `var.cpp` from `test/output/cpp`.

```js
// start 'hi_my_type_is' off as a dynamically typed string variable
var hi_my_type_is = "what"

// change to an integer
hi_my_type_is = 666

// here's a bool
hi_my_type_is = false

// and finally, a float
hi_my_type_is = 2.71828
```

<br>

> Comments also take on a familiar syntax:

```vb
// Inline comment

/*
  Multiline
  comment
*/
```

<br>

> Both of these are `int` variables, with the latter showing a *type inference* based on evaluation of the `rvalue` expression:

```csharp
int zero = 0
    one := 1    // tabbed for visibility
```

<br>

> `Optional Verbosity`:
<br>
> Before moving on, let's explain a primary motivator in the Express language development. Above, you will observe that the usage of commas and semicolons as statement delimiters seems to be optional. This is because they are! Statement delimiters are acceptable if you prefer, but are in no way required if you'd rather not deal with that.
<br>
> In the underlying parser architecture, they serve a semantic purpose by marking the end of a statement parse. By default, the ending of the statement will be semantically inferred if not *expressly* specified, however, there are compiler `flags` and [ECMA-335](https://www.ecma-international.org/publications/standards/Ecma-335.htm) `attributes` to modify the default action to enforce strict punctuation as granularly (or entirely) as you prefer.
<br>
> In this regard, having the flexibility to allow the compiler to semantically infer the end of the statement, while also retaining the ability to manually signal when a statement should end, can be very relaxing.
<br>
> This allowed flexibility is known as *`optional verbosity`* in Express and is one of the key motivators in it's development.

<br>

> Now, I wouldn't do this, but as a testament towards the semantic reasoning within the parser; you can even write statements on the same line as each other:

```csharp
string ayy = "ayy" string waddup = "waddup" int timestamp = 1527799745
```

<br>

> Usage of the `set` operator is permitted even outside of `object` declarations.

<!-- 
    The difference from a standard declaration or assignment is that the `set` operator will enforce a *`non-destructive`* local declaration of a variable; meaning it will not crawl the scope tree to perform reassignment and will enforce that it is the  
    crawl _down_ the scope tree instead of _up_?
-->

```vb
anotherOne: "anotherOne"
```

<br>

> Below shows a type inferred `object` where most of its properties are also type inferred.
<br>
> *Note:* The following statements until the ending brace are all within the `testerino` variable

```vb
testerino := {
    id: "ba3d4793-cfae-48d1-ad51-47cbfd70f98a"
```

<br>

> Reference one of the above variables in a new property declaration

```csharp
    time: timestamp
```

<br>

> You can also use the `assignment` operator along with a type to crimp the type of the variable instead of leaving it up for interpretation by the compiler

```vb
    float price = 55.3592,
```

<br>

> The `inference` operator can also be used within objects. Although currently it doesn't do anything different than the `set` operator within an object, it may have a more impactful use later

<!-- 
    Could make this a "redeclaration" operator to specifically reinterpret the type as well
-->

```vb
    dank_meme := true
```

<br>

> **Any** unicode character is supported

```vb
    🔥420🔥    : "🅱️laze it"
    ⒹⒿKhalid : anotherOne
```

<br>

> An `array` composition using the above definitions of `zero` and `one` to derive a `type` and `length` inference

```csharp
    ten: [ one, zero, one, zero ]
```

<br>

> A few nested objects

```vb
    throw_more_dots: {
        throw_more_dots: {
            more_dots: {
                more_dots: {
                    ok: "stop dots",
                },
            },
        },
    }
```

<br>

> Ending of `testerino`

```vb
}
```

<br>

> Arrays can be declared as well; below is a `static` `string` type `array` using composition to infer a `length`.

```csharp
string[] stringArray = [ "hi", "my", "name", "is", "scott" ]
```

<br>

> Expanding on the above example, delineatiation of elements from one another using commas isn't required but should be used at your own descretion of verbosity.
<br>
The spacing also doesn't matter, but *readable* code does. <br>
Again - its all *semantics* `¯\_(ツ)_/¯`

```csharp
string[] here_comes =
[
  ayy,
      waddup
             "its",
                    "dat"
                          "boi"
                                ]
```

<br>

> Quick power level check before blasting off...

```vb
if powerlevel < 9001 {
   powerlevel = 9001
   over9k = true
}
```

<br>

> A simple for loop:

```vb
percent: 0
for progress := 0, progress < 100, progress++ {
   percent = progress
}
```

<br>

> A key-based iterator for loop (for..in):

```csharp
int i = 0;
for index in [ 1, 2, 4 ] {
    i = index;
}
```

<br>

> A value-based iterator for loop (for..of):

```csharp
houstonWeHaveLiftOff := false
countdown := [ 9, 7, 6, 5, 4, 3, 2, 1 ];
for step of countdown {
  // Get ready for take off
  houstonWeHaveLiftOff = false
}
houstonWeHaveLiftOff = true
```
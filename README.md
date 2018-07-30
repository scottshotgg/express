# express

<!--
    Put a demo proposal in for GraalVM and stuff
-->

Express is an extremely flexible language that supports both static types and a dynamic type (`var` &mdash; think `JavaScript` or `Python`, except better) with an _extremely_ lite, dynamically embedded runtime within the binary allowing for typing to be as _weak_ or _strong_ as you like. The runtime currently does dynamic typing and RAII lifetime management, however, in the future it will include a greenthread scheduler a la Go, atomization of operations, garbage collection, RTTI and reflection, SQLite3/GraphQL embedded databases, DOM instantiation and manipulation, and a few other ideas that are currently only conceptual. Most features will be optional and the programmer will be allowed to enable and disable at compile time and possibly runtime (if a JIT/AOT is supported).<br>The main influences in the languages design are: `C++`, `JavaScript`, `Go`, and (atleast _conceptually)_ `Rust`.

For binary production, programs are currently _transpiled_ to C++ and then LLVM is subsequently invoked (along with `clang-format`) to produce the corresponding binary. There will also be a C++ program produced at compile time, which can be included in the output via a flag. At this time, transpiling is, _time-wise_, sufficiently more efficient than outputting LLVM tokens or building an intermediary using SSA/3AC. Later on, this will most likely be changed in favor of direct LLVM token production when features either become too much of a burden to implement and maintain in C++ or the transpiler development lags too much to adequently support forwarding the development of the language.

Each stage of the compiler (lexer, syntax parser, semantics parser, and C++ transpiler) are currently all implemented in `Go` and may be converter to `Rust` later on, but a `JavaScript` implementation in Node is also being developed simultaneously and will later on be consolidated with this repo after a reorganization of the file structure.
<br>
It is currently located at https://github.com/Swivelgames/Express/tree/alt/node

<br>

## State of the Compiler

### `Stages Implemented:`

- [x] Lexer
- [x] Syntax
- [x] Semantics
- [x] C++ Transpiler with `clang-format`
- [x] Binary

<br>

### `Features Implemented:`

_[~] means the feature is still being decided on_<br>
_[=] means I am currently working on implementing that feature_<br><br>

- [x] Blocks
- [x] Basic types (`int`, `bool`, `float`, `string`)
- [x] `object` type
    - [ ] statements within objects (making objects and blocks identical structures)
- [ ] `struct` type
    - [ ] Tags
- [ ] `function` type
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
- [-] `var` type
    - [x] Basic type encapsulation (`int`, `bool`, `float`, `string`)
    - [ ] `object` type encapsulation
    - [ ] `<struct>` type encapsulation
        - _This is doing the same thing as casting your struct to an object_
    - [~] `array` type encapsulation
        - _Leaning against: not sure if I want to allow a single `var` to be able to contain multiple values_
    - [~] `function` type encapsulation
        - _A var holding a function doesn't make a lot of sense right now_
- [ ] Type modifiers - keywords
    - [ ] `array` postfix
    - [ ] `unsigned` prefix
    - [ ] `constant` prefix
- [ ] Type modifiers - shorts
    - [ ] `s` postfix
    - [ ] `u` prefix
    - [ ] `c` prefix
- [=] Function usage
    - [=] Function declaration
        - [=] no args and no returns
        - [=] args without returns
        - [=] returns without args
        - [x] args and returns
    - [ ] Function call
- [-] Access modifiers
    - [x] `private`
    - [ ] `public`
    - [ ] `fileprivate`

<br>
<br>

## How To Contribute

### Proposals

Proposal submission is **the** method to contribute your ideas and work into Express. The first step to submitting a proposal is a usually `feature` request through a PR. Below contains the list of steps and the corresponding criteria for submission and approval.
<br>
<br>

- `Features:`
   <br>
    _Features_ should be submitted in the form of a markdown file (`title_of_your_feature.md`) describing the proposed abilities into a folder within `proposals/features/<appropriately_labeled>` folder. Proposals need to use fenced code blocks with a combination of `cs`, `vb` and `js` highlighting tags and contain an Express program example of your feature.
<br>
<br>

- `Ideas:`
   <br>
    _Ideas_ about the languages direction, internals related to the architecture, or general core-langaage level imlementations requests should go into the `proposals/ideas` folder and should be formatted as a whitepaper with detail containing a beginning narative that will explain <u>where</u> and <u>how</u> the idea came from, atleast two supporting arguments that should answer the question of <u>why</u> your proposal should be considered, and one informal/abstract (i.e, pseudocode) use-case on <u>how</u> it would be used after implementation.
<br>
<br>

- `Implementations and Experiments:`
   <br>
    _Implementations_ and _Experiments_ should be done in a separate branch off of the repo after approval of a submitted `feature` request. When you are ready to submit your implementation, you should fetch the latest upstream `master` and merge your code in locally, confirm that tests still work, and then submit a PR into `master`.
   <br>
    Your PR should also contain an Express test program in the `test/output/programs` folder, along with the respective `lex`, `syn`, `sem`, `cpp`, and `bin` files that will be used to verify the tests.
<br>
<br>

>_`Note`_: If your PR **does not** meet the criteria when it is submitted, cannot be merged into the parent branch, or standard tests do not pass, it will be _automatically_ denied by the CI/CD pipeline with comments about failure. You will need to submit another PR after fixing the mentioned issues.

<br>
<br>

## Example Program

Below is currently the most advanced program that can be written in Express. You can find the *full* **uncommented** version under `advanced.expr` under the `test/programs` directory.
<br>
In the program below, you will find a few examples of the allowed flexibilities and optional verbosity that allow the language to be so _Expressive_.

<br>

```html
Let's Begin:
```

____

<br>

> Start off by declaring some variables:

```cs
int powerlevel = 0x2327;
bool over9k = false;
float pi = 3.14159265359;
string arizona = "iced out boys";
```

<br>

> In addition to the basic static types, Express also supports using dynamically typed variables as well.<br> It is important to note that these variables are dynamically typed at _run time_ and thus will incur a performance penalty in constrast to static variables correlating to the same shadow type.<br> For a closer look at how the runtime manages this in C++, see the source code in `lib/var.cpp`.<br> More documentation and better comments (_as if there is any, lol - really though :^)_) will be added later - <sup><sub>_i promise_</sub></sup>

```js
// start 'hi_my_type_is' off as a dynamically typed string variable
var hi_my_type_is = "what"

// now i'm an integer
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

```cs
int zero = 0
    one := 1    // tabbed for visibility
```

> _`Note`_: Type inference will _never_ produce a `var` or a `struct` type.<br> Logically speaking, the `var` type could be considered the ground state for any variable type and thus would always resolve as a possible type. Furthermore, if you are specifying to infer a variables type, it doesn't make much sense, functionally, to respond with a generic container.<br> On the other hand, `struct` is a a different issue. Structs and objects are very similar ideas, however, one is dynamic - `object`, and the other is not - `struct`. Thus, when assigning a `BLOCK` to a variable in Express, it will assume you do not want this to be a static `struct` type or else you would have declared it yourself at compile-time. However, this does not prevent you from declaring an `object` at compile-time or a `struct` at run-time using a type specification at declaration.

<br>

> `Optional Verbosity`:<br> Before moving on, let's explain a primary motivator in the Express language development. Above, you might have observed that the usage of commas and semicolons as statement delimiters seems to be optional, and indeed they are! Statement delimiters are not required, but are acceptable (`;` or `,`) if you'd prefer to be more verbose or are accustomed to C-style programming.<br>In the underlying parser architecture, they (`;` or `,`) serve a semantic purpose by marking the end of a statement parse which will manually command the compiler to dump the current parse. By default, the ending of the statement will be semantically inferred, however, there are compiler `flags` and [ECMA-335](https://www.ecma-international.org/publications/standards/Ecma-335.htm) `attributes` to modify the default action and enforce strict punctuation as granularly (or entirely) as you prefer.<br> In this regard, having the flexibility to allow the compiler to semantically infer the end of the statement, while also retaining the ability to manually signal when a statement should end, can be very satisfying and relaxing when programming.<br> This allowed flexibility is known as _`optional verbosity`_ in Express and is one of the key motivators in it's development.

<br>

> Now, I wouldn't do this, but as a testament towards the semantic reasoning within the parser; you can even write statements on the same line as each other:

```cs
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

> Below shows a type inferred `object` where most of its properties are also type inferred. The following statements until the ending brace are all conatined within the `testerino` variable.

```js
testerino := {
    id: "ba3d4793-cfae-48d1-ad51-47cbfd70f98a"
```

<br>

> Reference one of the above variables in a new property declaration

```cs
    time: timestamp
```

<br>

> You can also use the `assignment` operator along with a type to crimp the type of the variable instead of leaving it up for interpretation by the compiler

```cs
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

> An `array` composition using the above definitions of `zero` and `one` to derive a `type` and `length` inference:

```js
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

> Ending of the `testerino` object

```vb
}
```

<br>

> Arrays can be declared as well; below is a `static` `string` type `array` using composition to infer an array `length`:

```cs
string[] stringArray = [ "hi", "my", "name", "is", "scott" ]
```

<br>

> Expanding on the above example, delineatiation of elements from one another using commas follows the same logic as the aforementioned statement delimiters. It isn't required but should be used at your own descretion of verbosity.<br> The spacing also doesn't matter, but *readable* code does. <br> Again - its all *semantics* `¯\_(ツ)_/¯`

```cs
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

```cs
int i = 0;
for index in [ 1, 2, 4 ] {
    i = index;
}
```

<br>

> A value-based iterator for loop (for..of):

```cs
houstonWeHaveLiftOff := false
countdown := [ 9, 7, 6, 5, 4, 3, 2, 1 ];
for step of countdown {
  // Get ready for take off
  houstonWeHaveLiftOff = false
}
houstonWeHaveLiftOff = true
```

<br>

> A simple function:

```go
func lookAtme() {
    // function contents...

    something := "else"
}
```

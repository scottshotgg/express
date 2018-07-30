# TODO

## `General`

_[~] means the TODO may not happen_<br>
_[=] means I am currently working on completing that TODO_<br><br>

- [=] Fix arbitrary object statements
- [=] Implement _index_ `[]` and _access_ `.` operations
- [=] Implement all binary operations for basic types
- [=] Figure out std lib usage
  - [=] Import statement
  - [=] Auto import/lib adding
  - [=] Create `print`/`println` functions<br><br>
- [=] Make `var.cpp` only printout if a variable is asserted or if `DEBUG` is turned on
- [=] Comment `lib` - atleast `lib/var.cpp`
- [ ] Make a demo `feature` request, `idea`, and `misc`
- [ ] Look into making header files for `lib`
  - [ ] Look into making a single header file
- [ ] Setup circleCI to build the project
- [ ] Clean up:
  - [ ] `samples`
  - [ ] `documentation`
  - [ ] `documentation_old`
  - [ ] `lib/var.cpp`<br><br>
- [=] Commands for:
  - [ ] `lex [file]`
  - [ ] `parse [type=syn,sem] [file]`
  - [ ] `transpile [file] {translation_lang}`
  - [ ] `bin [file] {translation_lang}`
  - [ ] `compile [file]`
  - [ ] `run [file]`
  - [ ] `test` with sub commands:
    - [ ] `lex {file}`
    - [ ] `syn {file}`
    - [ ] `sem {file}`
    - [ ] `cpp {file}`
    - [ ] `bin {file}`
    - [ ] `all {file}`

<br>

## `Architecture Redesign`

- [=] Convert to cobra/viper for commands
- [ ] Move `TestAll` from `semantic_test.go` to it's own package
- [ ] Entirely restructure the data of the compiler (interfaces, structs, `token2`, etc)
  - [ ] Look into the collection/token method that Joseph implemented in the JS verion
  - [ ] Try to implement scatter-gather method of lexing/parsing
  - [ ] Fix the weird requirement for a newline at the end of an expr program, this has been fixed in `lex/lex_old.go`
- [ ] Develop tests for every part of the lexer, parsers, and transpiler
- [ ] Rearchitecture the parser into multiple packages, or subpackages; especially for the transpiler

<br><br>

## `Misc And Extra`

- [~] Attempt to make a formal grammar
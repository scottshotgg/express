# TODO

## `General`

_[~] means the TODO may not happen_<br>
_[=] means I am currently working on completing that TODO_<br><br>

- [=] Make an install script
- [=] Refactor build scripts
- [=] Add a run command
  - [=] Binary goes in temp dir
  - [=] Runs the binary after compiling the code
  - [=] Doesn't save the binary past run time
- [=] Add _proper_ logging back in with Uber's Zap logger
  - [=] Only log within the compiler if `EXPR_DEBUG` is `"true"`
    - [=] Make a list (starting documentation) on different environment flags
    - [=] Make separate `EXPR_DEBUG_XXX` for `LEX`, `SYN`, etc
    - [=] Fix arbitrary object statements
- [=] `stdlib` support/construction
  - [=] Investigate constructing a _generalized_ function to create the Express `stdlib` -> `libc`/`libstdc++` bindings
  - [ ] Basic/preliminary `libc` and `libstdc++` bindings
    - [ ] `open`
    - [ ] `read`
    - [ ] `write`
    - [ ] `exit`
    - [ ] AIO operatations: _https://www.gnu.org/software/libc/manual/html_node/Function-Index.html_
      - [ ] `aio_init`
      - [ ] `aio_cancel`
      - [ ] `aio_read`
      - [ ] `aio_write`
      - [ ] `aio_suspend`
      - [ ] `aio_return`
      - [ ] `aio_error`
      - [ ] `aio_fsync`
  - [=] `import` statement
  - [=] Auto import/lib adding
  - [x] Create `print`/`println` functions<br>_just make the functions translate to C++_<br><br>
- [ ] Implement _index_ `[]` and _access_ `.` operations
- [ ] Implement all binary operations for basic types
- [=] Fix Go funcs for building
  - [ ] Figure out how to do IDE level debugging in Go
- [x] Figure out `stdlib` usage
- [ ] Clean up `stdlib`
  - [ ] Only printout if `EXPR_DEBUG_TRANS` is `"true"`
  - [ ] Comment `stdlib` - atleast `lib/var.cpp`<br><br>
- [ ] Augment functions:
  - [ ] Make the return value declared at the start of the function def
  - [ ] no args - no returns
  - [ ] args - no returns
  - [ ] no args - returns
  - [ ] Multiple returns
    - _wait until we can get multi ident declarations (i.e, `int i, j = 0`)_
    - _Could just do this using objects... but then we would have to use var's everytime we wanted multiple returns..._
  - [ ] Default argument values in function def like Python
  - [ ] Named arguments in function call like Python
- [~] Look into making header files for `lib`
  - _Might not need to do these after the above `stdlib` work_
  - [~] Look into making a single header file
    _Would be nice, but might not be needed if we can just use `libc`/`libstdc++`_
- [ ] Investigate 'tracing' or some other way to look at how well the program runs
  - [ ] Testing vs C++ and Go equivalent programs
  - _analyze memory usage, pressure, management_
  - _cpu metrics during and not during an intensive stress test_
  - _safety_
  - _ease of use_
- [ ] Clean up:
  - [ ] `samples`
  - [ ] `documentation`
  - [ ] `documentation_old`
  - [ ] `lib/var.cpp`<br><br>
- [=] Commands for:
  - _Need to think about these commands_
  - [ ] `lex [file]`
  - [ ] `parse [type=syn,sem] [file]`
  - [ ] `transpile [file] {translation_lang}`
  - [ ] `bin [file] {translation_lang}`
  - [=] `run [file]`
  - [ ] `test` with sub commands:
    - [ ] `lex {file}`
    - [ ] `syn {file}`
    - [ ] `sem {file}`
    - [ ] `cpp {file}`
    - [ ] `bin {file}`
    - [ ] `all {file}`

<br>

## `Architecture and Redesign`

- [x] Convert to cobra/viper for commands
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
- [ ] Make a TextMate grammar
  - [ ] VS Code extension
    - [ ] Syntax highlighting
    - [ ] Snippets
    - [ ] Examples
- [ ] Make a demo `feature` request, `idea`, and `misc`
- [ ] Set up circleCI to build the project
- [ ] Set up official GitHub org / website
- [ ] Look into Graal VM
- [ ] Garbage collection?
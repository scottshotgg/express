# TODO

## General

- [ ] Make a demo `feature` request, `idea`, and `misc`
- [ ] Comment `lib` - atleast `lib/var.cpp`
- [ ] Look into making header files for `lib`
  - [ ] Look into making a single header file
- [ ] Setup circleCI to build the project
- [ ] Clean up:
  - [ ] `samples`
  - [ ] `documentation`
  - [ ] `documentation_old`
  - [ ] `lib/var.cpp`<br><br>
- [ ] Commands for:
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

## Architecture Redesign

- [ ] Convert to cobra/viper for commands
- [ ] Move `TestAll` from `semantic_test.go` to it's own package
- [ ] Entirely restructure the data of the compiler (interfaces, structs, `token2`, etc)
  - [ ] Look into the collection/token method that Joseph implemented in the JS verion
  - [ ] Try to implement scatter-gather method of lexing/parsing
  - [ ] Fix the weird requirement for a newline at the end of an expr program, this has been fixed in `lex/lex_old.go`
- [ ] Develop tests for every part of the lexer, parsers, and transpiler
- [ ] Rearchitecture the parser into multiple packages, or subpackages; especially for the transpiler

## Misc And Extra

- [~] Attempt to make a formal grammar
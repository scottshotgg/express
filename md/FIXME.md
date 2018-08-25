# FIXME

These will show up in the TODO utility:

- FIXME: add all errors to the `errors.go` file
- FIXME: move `TestAll` and other tests to `parse_test.go`
- FIXME: change all `token.Value{}` returns to `*token.Value{}` so that we can return nil instead
- FIXME: work on `token2`: will help structure compiler
- FIXME: get the variable/token format straight

- FIXME: make a textmate2 plugin
- FIXME: make a vscode plugin
- FIXME: investigate making header files for libraries
- FIXME: try to work on `bindc` and think of some other options / designs
- FIXME: make an example unqlite and sqlite usage within the core language
- FIXME: build a more general compiler architecture instead of a specific compiler

- TODO: read about how to actually generate object code/LLVM stuff
- TODO: fix linter errors

- CANT: fix gofunc testing: fix function string `cpp.go` being global
  - see below
- CANT: maybe try doing `t.Parallel()` and refactor `TestAll`
  - something is preventing the tests from being run in parallel, will have to look at it more at a later date
- DONE: remove all `fmt.Print*()`, replace with preliminary `logger.*()`
- DONE: add `logger.*()` with `zap.Error()`
- DONE: wrap all errors: might not need after using `zap.Error()`
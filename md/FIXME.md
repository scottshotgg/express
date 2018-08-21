# FIXME

These will show up in the TODO utility:

- FIXME: work on `token2`: will help structure compiler
- FIXME: try to work on `bindc` and think of some other options / designs
- FIXME: get the variable/token format straight
- FIXME: move `TestAll` and other tests to `parse_test.go`
- FIXME: fix gofunc testing: fix function string `cpp.go` being global
- FIXME: maybe try doing `t.Parallel()` and refactor `TestAll`
- FIXME: remove all `fmt.Print*()`, replace with preliminary `logger.*()`
- FIXME: add `logger.*()` with `zap.Error()`
- FIXME: wrap all errors: might not need after using `zap.Error()`
- FIXME: change all `token.Value{}` returns to `*token.Value{}` so that we can return nil
- FIXME: add all errors to the `errors.go` file
- FIXME: investigate making header files for libraries
- FIXME: make an example unqlite and sqlite usage within the core language
- FIXME: build a more general compiler architecture instead of a specific compiler
- FIXME: make a textmate2 plugin
- FIXME: make a vscode plugin
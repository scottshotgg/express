# TODO

- [ ] Convert to cobra/viper for commands<br><br>
- [ ] Move `TestAll` from `semantic_test.go` to it's own package<br><br>
- [ ] Rearchitecture the parser into multiple packages, or subpackages; especially for the transpiler<br><br>
- [ ] Commands for:
    - [ ] `lex [file]`<br><br>
    - [ ] `parse [type=syn,sem] [file]`<br><br>
    - [ ] `transpile [file] {translation_lang}`<br><br>
    - [ ] `bin [file] {translation_lang}`<br><br>
    - [ ] `compile [file]`<br><br>
    - [ ] `run [file]`<br><br>
    - [ ] `test` with sub commands:
      - [ ] `lex {file}`
      - [ ] `syn {file}`
      - [ ] `sem {file}`
      - [ ] `cpp {file}`
      - [ ] `bin {file}`
      - [ ] `all {file}`<br><br>
- [ ] Setup circleCI to build the project
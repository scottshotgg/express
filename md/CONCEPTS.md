# Compiler Concepts

## lexer

- generate tokens by collecting lexemes (letters)
  - how to deal with things such as "for" vs "foreach"
- need to be able to define _lexemes_ in JSON format (other formats to follow)
- need to be able to define _tokens_ in JSON format (other formats to follow)
- `F` `O` `R` ... `B` `O` `D` `Y` -> `FOR` ... `BODY`
- could split by whitespace to parallelize the lexing
- do not filter out comments like we currently do

<br>

## syntactic

- translate tokens into semantic structures (i.e, statments, loops, variable declarations, etc)
- `FOR` ... `BLOCK` -> `for` with a body property
- **maybe** add separators and filter out arbitrary whitespace so that the semantic parsing can be done semi parallelized
- pass the comments along

<br>

## semantic

- comments should live until here
- analyze semantic structures and generate a stack of scopes and variables from tokens
- **or** make it build the AST from here and then walk that AST in cpp
- apply rules of Express here (i.e, cannot call interface function with object, determining var types, correct operations, add import statments, link in functions, etc)

<br>

## cpp

- generate cpp code from stack of scopes and variables

<br>
<br>

## resources

- https://blog.gopheracademy.com/advent-2014/parsers-lexers/
- https://github.com/fatih/astrewrite
- https://github.com/cloudfoundry/go-pubsub

<br>
<br>
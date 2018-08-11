# How To

### A brief description of how I generated the Express stdlib from libc

In creating Express, it was apparent.... stdlib for Express
using regexpr

<br>

> First step (parsing, yada, yada, ...) is getting all of the information for libc
```
info libc | grep "" > info_libc
```

<br>

> To get all the functions names:
```
Function: .*? \(.*?\)\n
```
> Afterwards, just use VS Code to delete all of the `Function:` captures

<br>

> To get all the data types
 ```
-- Data Type:
 ```

### How I'm gonna do it:
  > 1. Write a frontend parser for C
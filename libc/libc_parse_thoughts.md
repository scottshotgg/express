### Thoughts on how to parse libc

Using `struct aiocb` defined on line 26512:

After finding all dashdash objects by regex:
Start processing all `Data Type: struct` objects

1. `struct` would tell us that we need to make an Express struct
2. 
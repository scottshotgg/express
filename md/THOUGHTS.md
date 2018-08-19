# Thoughts

### `structs:`
  - structs at **compile-time** are actual c++ structs with overloaded `operator<<` for printing
  - _however_ - structs at **run-time** are type-locked objects
  - this would also **force** you to declare objects and vars which would solve the thought issue I was having before about structs not being the default declaration instead of objects
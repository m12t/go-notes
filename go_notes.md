# Some Notes on the Go Programming Language
- all variables that are not explicity initialized are thusly implicitly initialized to the "zero value" for the variable's respective type, `0` for numeric types, `""` (empty string) for strings, arrays assign every element in the array to its default value, structs assign all fields to their default value, bool defaults to false, and `nil` is assigned for pointer, interface, map, channel, slice, function.
- because Go is garbage collected, local variabled created in a calling function can be "escaped" and get allocated to the heap instead of the stack, meaning using a pointer to a variable that was created in the called function is safe to access in the calling function.
-

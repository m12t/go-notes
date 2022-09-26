# Some Notes on the Go Programming Language

## Variables
- All variables that are not explicity initialized are implicitly initialized to the "zero value" for the variable's respective type, `0` for numeric types, `""` (empty string) for strings, arrays assign every element in the array to its default value, structs assign all fields to their default value, bool defaults to `false`, and `nil` is assigned for pointer, interface, map, channel, slice, function.
- Because Go is garbage collected, local variabled created in a calling function can be "escaped" and get allocated to the heap instead of the stack, meaning using a pointer to a variable that was created in the called function is safe to access in the calling function.
- All declared variabled must be used. However, there are times where syntax requires a variable, but the logic does not. An of this is a function that returns multiple values where you only want one value of the returned values. To satisfy the compiler, a _blank identifier_, `_` (underscore), is used on the unwanted variables.


## Switch Statements
- unlike C and many other languages, cases do not "fall through" in go. That is, if one case evaluates to true, it's block of code is executed and then execution resumes after the end of the switch block. `break` is not necessary as it is in languages where cases _can_ fall through.
- `switch { ... }` is called a "tagless switch" and is equivalent to `switch true { ... }`

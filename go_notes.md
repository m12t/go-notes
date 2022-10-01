# Some Notes on the Go Programming Language

## Variables
- All variables that are not explicity initialized are implicitly initialized to the "zero value" for the variable's respective type, `0` for numeric types, `""` (empty string) for strings, arrays assign every element in the array to its default value, structs assign all fields to their default value, bool defaults to `false`, and `nil` is assigned for pointer, interface, map, channel, slice, function.
- Because Go is garbage collected, local variabled created in a calling function can be "escaped" and get allocated to the heap instead of the stack, meaning using a pointer to a variable that was created in the called function is safe to access in the calling function.
    - Local variables live until they become *unreachable* at which point they are garbage collected.
    - It's worth noting that escaped variables can cause unwanted performance tolls because they require additional allocation on the heap.
- All declared variabled must be used. However, there are times where syntax requires a variable, but the logic does not. An of this is a function that returns multiple values where you only want one value of the returned values. To satisfy the compiler, a _blank identifier_, `_` (underscore), is used on the unwanted variables.
- Variables can be initialized to the values returned by a function.
- Variables can be created using the `new` function.
- Package-level variables never become unreachable, and therefore don't get garbage collected.
- To make a package-level variable exportable (that is, it can be accessed by code that pulls in the package in which it lives), it must start with a upper-case letter.
# Pointers
- Pointers are comparable
    - two pointers are equal if they both point to the same address
    - a test of `p != nil` will return `true` if the pointer `p` points to a variable
    - it is safe for a function to return a pointer to a local variable declared inside of a function, despite that function returning and being removed from the stack.
# Type declarations
- a type declaration is a new _named type_ that has the same underlying type as an existing type.
- Two named types that both use the same underlying type are _NOT_ the same.
    - eg. `type Celsius float64` and `type Fahrenheit float64` both use the same underlying type, `float64`, yet compring a variable of `Celsius` type with another of `Fahrenheit` or trying to perform an arithmetic operation with both of them will throw an error.
        - this can be a useful feature for preventing logical errors such as adding temperatures in two different scales.
    - Comparison operations (`==`, `<`, `>` etc.) can be used between a value of a named type with the same type or with the underlying type. However, two different named types with the same underlying type may still not be compared.
- for every type, `T`, (including named types like `Celsius` above) there is a corresponding conversion operation (*not a function*) `T(x)` that will explicity convert the value `x` to the new type, `T`. Note, in order to convert `x` to `T`, they must both use the same underlying type or be unnamed pointers to variables of the same underlying type

## Switch Statements
- unlike C and many other languages, cases do not "fall through" in go. That is, if one case evaluates to true, it's block of code is executed and then execution resumes after the end of the switch block. `break` is not necessary as it is in languages where cases _can_ fall through.
- `switch { ... }` is called a "tagless switch" and is equivalent to `switch true { ... }`

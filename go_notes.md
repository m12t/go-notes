# Some Notes on the Go Programming Language

## Variables (gopl 2.3)
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
# Type declarations (gopl 2.5)
- a type declaration is a new _named type_ that has the same underlying type as an existing type.
- Two named types that both use the same underlying type are _NOT_ the same.
    - eg. `type Celsius float64` and `type Fahrenheit float64` both use the same underlying type, `float64`, yet compring a variable of `Celsius` type with another of `Fahrenheit` or trying to perform an arithmetic operation with both of them will throw an error.
        - this can be a useful feature for preventing logical errors such as adding temperatures in two different scales.
    - Comparison operations (`==`, `<`, `>` etc.) can be used between a value of a named type with the same type or with the underlying type. However, two different named types with the same underlying type may still not be compared.
- for every type, `T`, (including named types like `Celsius` above) there is a corresponding conversion operation (*not a function*) `T(x)` that will explicity convert the value `x` to the new type, `T`. Note, in order to convert `x` to `T`, they must both use the same underlying type or be unnamed pointers to variables of the same underlying type

## Switch Statements
- unlike C and many other languages, cases do not "fall through" in go. That is, if one case evaluates to true, it's block of code is executed and then execution resumes after the end of the switch block. `break` is not necessary as it is in languages where cases _can_ fall through.
- `switch { ... }` is called a "tagless switch" and is equivalent to `switch true { ... }`

## Packages and Files (gopl 2.6)
- Packages in Go are similar to libraries or modules in other languages. The source code for a package lives in one or more `.go` files.
- Packages provide a modularity, encapsulation, separate compilation, and reuse.
- Each package serves as a separate name space for declarations. That is, names in a package can be repeated across packages because they must be qualified by the package as with `package.name`.
- Packages can also control which identifiers are visible outside of a package, known as being _exported_. To be exported, an identified must begin with an upper-case letter.
- Package _doc comment_ is a description that immediately precedes the package declaration and described the package as a whole.
    - Only one file in each package should have a package doc comment
    - The conventional description is something like: "Package {name} does {what it does}
    - If the doc comment becomes too large to fit at the beginning of a file, it will typically be moved to a file of its own, `doc.go` within the package.
- A package init function, simply `func init() { /* ... */ }`, is a function that automatically executes when the program starts and can be used to initialize complex variables.

## Scope:
- defined as the region of source code where the value of a declared name referes to the value of the declared name.
- this is _NOT_ the same as a variable's lifetime. Lifetime is a _run-time_ attribute and refers to the amount of time during execution that a variable is accessible. Scope is a _compile-time_ property refering to the region of program text.
- A _syntactic block_ is a series of statements enclosed in curly braces: `{ ... }`.
    - Variables defined within a syntactic block are not visible outside of that syntactic block, but _are_ visible in child blocks, or other syntactic blocks within the current block.
- A generalization of the block notion includes scope groupings that don't contain curly braces, but that behave the same way is called _lexical scope_.
- The lexical blocks in Go are:
    - All source code, called the _universe block_
    - Code for each package
    - For each file
    - For each `for`, `if`, and `switch` statement
    - For each `case` in a `switch` or `select` statement
    - For each syntactic block, `{}`
- When the compiler encounters a reference to a name, it searches from the most local lexical block outwards until it finds a definition. If none is found, it will throw an "undeclared name" error. This local-first search means that the innermost definition will supercede any previous definitions. In other words, the inner declaration is said to "shadow" or "hide" the outer one, making it inaccessible.
- Short variable declarations demand an awareness of scope:
    ```go
    var cwd string
    func init() {
        // since neigher `cwd` nor `err` are already declared in the
        // current lexical block (the `init` function), this will cause
        // the short variable declaration to assign them both as local variables
        // and the global `cwd` declared at the file-level scope will not be modified.
        cwd, err := os.Getwd() // compile error: unused: cwd
        if err != nil {
            log.Fatalf("os.Getwd failed: %v", err)
        }
    }
    ```
    - Perhaps worse, is if the value of `cwd` _is_ referenced locally, causing no compiler error about an unused variable. This is a tough bug to catch since it is silent and you would need to check the value of `cwd` to see that it wasn't modified as it should have been.

## Basic Data Types (3)
- there are 4 categories of types in Go:
1. Basic types
    - numbers, strings, booleans
2. aggregate types
    - arrays and structs
    - form complex data types of combining values of several basic types
3. reference types
    - pointers, slices, maps, functions, channels
    - refer to program variables or state _indirectly_. The effect of an operation is thus applied to all copies of the reference
4. interface types


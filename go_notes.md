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


### Strings
- Strings are immutable. This means that copying and a string of any length or taking a substring are cheap because they can safely point to the same location in memory.
- String literals are written with backticks (\`...\`) and no escape processing is performed other than elimination of carriage returns `\r`.
- ASCII (7 bits) was the original character set, but it couldn't support all languages and symbols we use today. Today, Unicode is used (`int32`, aka `rune` in Go). However, using 32 bits is wasteful for nearly all cases. Enter, UTF-8, a variable length encoding of Unicode.
- Converting an integer to a string as in `string(65)` will return the UTF-8 value of the integer interpreted as a rune. Eg. `fmt.Println(string(65))` will retun `A`, not `65`.


## Constants
- untyped constants have much greater precision than any of the basic types available.


## Arrays
- questions about index value pair arrays. How are these stored in memory? linked list? hashmap? array of alternating addresses? Is the index required to be an int? if so, maybe the array creation is just slower and the ordering is based on what is specified in the creation.
- The following example is interesting because it gives the best of both properties, arrays, and hashmaps in having fast iteration and fast indexing, while having lookup ability.
```go
type Currency int

const (
    USD Currency = iota
    EUR
    GBP
    RMB
)

symbol := [...]string{USD: "$", EUR: "E", GBP: "G", RMB: "R"}

fmt.Println(USD, "USD:", symbol[USD])

```


## Maps (4.3)
- A reference to a hashmap
- Can be created with `make` like `table := make(map[K]V)` where k is the key type and v is the value type
- key and value types don't need to match, but all keys must be of the same type and all values must be of the same type.
- keys must be comparable, so slices don't work as keys, and floats are a bad choice for this reason.
- map lookups are always safe, so calling `delete()` on an item not in a map is safe because the behavior of attempting to access a key not within a map is that the zero value of the type is returned.
- To know whether a value was in the map, there is a multiple return on access: `age, ok := ages["bob"]` if the key is not in the map, the value will be zero but the `ok` bool will be `false`
- however, storing to a `nil` map causes a panic. if `ages == nil`, `ages["Freddy"] = 100` will panic. You must allocate a map before being able to store to it.
- enumerating the key value pairs of a map is done with `range` like `for key, value := range myMap { fmt.Println(key, value) }`


## Slices (4.2)
- a lightweight data structure that gives access to some or all elements of an array, known as the slice's underlying array
- can be created with `s = make([]byte)`
- a slice has 3 components:
    1. A pointer to the first elemen in the array that is reachable to the slice
    2. A capacity, typically the number of elements from the pointer to the end of the array (why only typically? what other options)
    3. A length, the number of elements in the slice. *This cannot exceed capacity*
- stlicing beyong the capacity of the slice induces a panic. Slicing beyond the length of the underlying array extends the slice, allowing for it to be longer than the original.
- Unlike arrays, slices are *not* comparable. In fact, the only allowed slice comparison is with `nil` since this is the zero value of a slice. To test if a slice is empty, use `len(s) == 0`, not `s == nil` because the following slices are all empty (and have len(s) == 0) but one of them is not nil:
    ```go
        var s []int    // len(s) == 0, s == nil
        s = nil        // len(s) == 0, s == nil
        s = []int(nil) // len(s) == 0, s == nil
        s = []int{}    // len(s) == 0, s != nil
    ```
- appending to a slice gives amortized constant time since there is a growth factor used on the underlying array so that a new allocation is only needed periodically, with most calls to append simply extending the slice, not copying an array in memory. However, becasue of the unknown nature of when a growth is performed, it isn't safe to access an old slice after an append. Because of this, it's good practice to append to the same name, like `names = append(names, "Michael")`. This is not only for append, but for any function that may change the length or capacity of a slice, or make it point to a different underlying array.


## Maps (4.3)
- A reference to a hashmap
- Can be created with `make` like `table := make(map[K]V)` where k is the key type and v is the value type
- key and value types don't need to match, but all keys must be of the same type and all values must be of the same type.
- keys must be comparable, so slices don't work as keys, and floats are a bad choice for this reason.
- map lookups are always safe, so calling `delete()` on an item not in a map is safe because the behavior of attempting to access a key not within a map is that the zero value of the type is returned.
- To know whether a value was in the map, there is a multiple return on access: `age, ok := ages["bob"]` if the key is not in the map, the value will be zero but the `ok` bool will be `false`
- however, storing to a `nil` map causes a panic. if `ages == nil`, `ages["Freddy"] = 100` will panic. You must allocate a map before being able to store to it.
- enumerating the key value pairs of a map is done with `range` like `for key, value := range myMap { fmt.Println(key, value) }`


## Functions (5):
- functions are _first class_ values in Go. Functions have types and can be assigned to variables, passed as arguments to other functions, or returned from other functions.
- a _variadic_ function accepts a varying number of final inputs. The notation is the following:
    ```go

    func addNums(x []int, y...int) []int {
    }
    ```
- go has no default arguments... explore the behavior of `append()` accepting nil as in page 123 of tgpl
- a function can have named results, in which case they are variables initialized to the default value of their type.
    - When named results are used, a "bare return" can be employed. This is where the function's return statement doesn't explicity list out the variables to return, instead relying on the prescribed named values in the function declaration.
    - This  can make code more DRY and reduce errors where the returned values are modified and all return statements within the function must be updated.
    - However, this also comes at the cost of readability.
- The flow of a Go function typically follows a similar pattern:
    - After checking an error, failure is usually dealt with before success. If failure causes the function to return, the logic for success is not indented within an else block but follows at the outer level. Functions tend to exhibit a com- mon structure, with a series of initial checks to reject errors, followed by the substance of the function at the end, minimally indented.
- Functions are comparable to nil, but not comparable to other functions.
    - This is because of "hidden variables" where functions store state. An example of this is a function that declares a local variable and then returns an anonymous function that acts on that local variable. Since anonymous functions have access to the entire lexical environment, the anonymous function will act on the declared local variable. See tgpl 5.6 (gopl.io/ch5/squares) for an in-depth example.
- Iteration variable capture is when function values created within a loop "capture" a variable by using its memory location, not its value. This is a difficult bug to catch and often leads to strange patterns like assining a new variable to one of the same name to "pull-in" a local copy for usage, as in `dir := dir` within a loop. See tgpl 5.6.1
### Errors (5.4)
- Errors in Go are propogated using a multiple return.
- By convention, if a function can error, the error is returned as another value, typically the last value.
- If a function's error can only have one cause, it is typically retuned as a boolean, `ok`.
- If a function can error on more than one cause, the customary return is `err`, a variable of type `error`. If no error is encountered, `nil` is used to signify success.
### Defer (5.8)
- The `defer` keyword is a prefix to a function or method call and it means that execution of that call will happen once the parent function finishes exeecuting, be that by a return statement, falling off, or panicking.
- Any number of calls may be deferred. The order of execution of defered functions is the reverse order in which they were deferred. Think: stack data structure.
- A defer statement is typically used to ensure that resources are closed regardless of what happens inside the function that opened the resource.
    - It's customary to make the deferred call immediately after successful access of the resource.
- `defer` can also be used when accessing mutexes. Eg. the mutex is unlocked, a defer call to lock the mutex is made, then operation(s) is/are performed and the rest of the funciton can safely return or fall off.
### Panic (5.9)
- A panic is when a runtime error is encountered in a Go program, such as a `nil` pointer dereference, or an out-of-bounds array read.
- When Go panics, normal execution stops where the panic was generated, all defered function calls in that goroutine are executed, the program crashes with a log message.
- Panics don't necessarily have to be generated by the runtime. Panics can also be raised in code by calling `panic()`, accepting any value as an argument.
### Recover (5.10)
- In most cases, a panic should cause the program to crash and execution to stop. However, there are times where this is undesirable. `recover()` allows normal execution to resume after a `panic` and captures the value returned by panic.
- The function that panicked does not resume execution, rather, it returns normally and execution resumes with the calling function.
- For many reasons, `recover()` should be used sparingly. Especially if the cause of the panic is a package or code you don't maintain.


### Structs
- Structs can be "embedded" into other structs (4.4.3), giving access to the embedded struct's fields
- 


### Interface (7)
- Interfaces are a mapping of methods to concrete types. In other words, interfaces are a named collection of method signatures for a type.
- 

### channels
- send operations on _unbuffered_ channels do block.
- send operations on _buffered_ channels do NOT block, _unless_ the channel is full.


### concurrency
- A data race is UB when 2 or more goroutines access the same variable and at least one of those is a write.
- 3 ways to prevent data races:
    1. never update (duh)
    2. limit access to a single goroutine. Concurrency here is using channels to request access in a queue essentially. Access is still confined to a single goroutine.
    3. use a lock/mutex to limit access to one single goroutine at a time.
- how does the Go runtime handle multiple calls to mutex.Unlock?? Is it a queue (FIFO)??
- interesting pattern from Section 9.7 of TGPL where an unbuffered channel is created and passed into a function (actually it's packed into a struct and pushed onto a channel where it is later serviced by that channel's consumer), then the calling function immediately reads (blocking) from the channel it just created until the response it returned.


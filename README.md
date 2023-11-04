# Go-Lab
This is my repo to commit code daily to learn how to program in Go.

## Resources
- [go101](https://go101.org/)
- [Microsoft](https://learn.microsoft.com/en-us/training/paths/go-first-steps/)
- [Digital Ocean](https://www.digitalocean.com/community/tutorial-series/how-to-code-in-go)
- [Official Go](https://go.dev/doc/tutorial/)


### Day 1
- created hello.go that prints out a simple message "Hello, World!".
- Refactored hello.go module to use locally created example.com/greetings module.
- Refactored hello.go and greetings.go for error handling when no name is given.

### Day 2
- created var.go in /microsoft that declares and initializes variables and constants in Go.

In Go, you have four categories of data types:
- Basic types: numbers, strings, and booleans
- Aggregate types: arrays and structs
- Reference types: pointers, slices, maps, functions, and channels
- Interface types: interface

Integer numbers
keyword to define an integer type is int. Go also provides the int8, int16, int32, and int64 types, which are ints with a size of 8, 16, 32, or 64 bits, respectively. You'll need to cast explicitly when a cast is required.

A rune is simply an alias for int32 data type. It's used to represent a Unicode character (or a Unicode code point).

    rune := 'G'
    fmt.Println(rune)

Booleans
A boolean type has only two possible values: true and false.
    
    var do_you_cap bool = true

Strings
To initialize a string variable, you need to define its value within double quotation marks ("). Single quotation marks (') are used for single characters.

Escape characters
Use a backslash (\) before the character. For instance, here are the most common examples of using escape characters:

    \n for new lines
    \r for carriage returns
    \t for tabs
    \' for single quotation marks
    \" for double quotation marks
    \\ for backslashes

Default values
Here's a list of a few default values for the types we've explored so far:

    0 for int types (and all of its subtypes, like int64)
    +0.000000e+000 for float32 and float64 types
    false for bool types
    An empty value for string types
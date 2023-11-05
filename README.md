# Go-Lab
This is my repo to commit code daily to learn how to program in Go.

## Resources
- [go101](https://go101.org/)
- [Microsoft](https://learn.microsoft.com/en-us/training/paths/go-first-steps/)
- [Digital Ocean](https://www.digitalocean.com/community/tutorial-series/how-to-code-in-go)
- [Official Go](https://go.dev/doc/tutorial/)


# Day 1
- created hello.go that prints out a simple message "Hello, World!".
- Refactored hello.go module to use locally created example.com/greetings module.
- Refactored hello.go and greetings.go for error handling when no name is given.

# Day 2
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

Functions
You can only have one main() function in your program. If you're creating a Go package, you don't need to write a main() function.

Syntax for creating a function:

    func name(parameters) (results) {
        body-content
    }

Pointers
A pointer is a variable that contains the memory address of another variable. When you send a pointer to a function, you're not passing a value, you're passing a memory address. So every change you make to that variable affects the caller.

In Go, there are two operators for working with pointers:

    The & operator takes the address of the object that follows it.
    The * operator dereferences a pointer. It gives you access to the object at the address contained in the pointer.

# Day 3
## Control Flows
Syntax for compound if statements:

    func main() {
        if condition {
            fmt.Println("do something cool")
        } else another condition {
            fmt.Println("do something lame")
        } else {
            fmt.Println("i write in all lowercase cause i never cap")
        }
    }

You can also use switch statements to avoid chaining multiple if statements. Here is an example of multiple expressions:

    package main

    import "fmt"

    func location(city string) (string, string) {
        var region string
        var continent string
        switch city {
        case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
            region, continent = "India", "Asia"
        case "Lafayette", "Louisville", "Boulder":
            region, continent = "Colorado", "USA"
        case "Irvine", "Los Angeles", "San Diego":
            region, continent = "California", "USA"
        default:
            region, continent = "Unknown", "Unknown"
        }
        return region, continent
    }
    func main() {
        region, continent := location("Irvine")
        fmt.Printf("John works in %s, %s\n", region, continent)
    }

In some programming languages, you write a break keyword at the end of every case statement. But in Go, when the logic falls into one case, it exits the switch block unless you explicitly stop it. To make the logic fall through to the next immediate case, use the fallthrough keyword.

if for loops you can use break keyword to make an logic exit. You can also use continue to skip the current iteration of a loop.


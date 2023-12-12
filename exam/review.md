---
geometry: margin=1in
documentclass: article
fontsize: 12pt
---

# Semester Review

## Format

The exam will have two parts. The first part is on paper and consists of code with print statements. You will have to understand what the code does to determine what is printed. When you complete the first part, you will turn it in and receive part two. It will consist of a programming problem similar to the exercises on http://headfirstgo.com/. You will do this part on the computer and turn in your code on Teams.

## Topics

### Chapter 1: Syntax Basics

1. General template of a program
1. Basic syntax rules
1. Use of `fmt.Println()`
1. Use of escape codes `\n \t \" \\`
1. Interpreting error message
1. Using functions from other packages
1. How to write a `main()` function
1. Using function return values
1. Math operations and comparison operators
1. Types `int`, `float64`, `bool`, and `string`
1. Declaring variables with `var`
1. Zero values of variables
1. Short variable declarations with `:=`
1. Variable naming rules
1. Type conversions such as `float64`
1. Go commands `go build`, `go run`, `go fmt`, and `go version`

### Chapter 2: Conditionals and Loops

1. Calling methods
1. Multiple return values from a function or method
1. Conditionals with `if`, `else`, and `else if`
1. Checking for errors
1. Logging fatal errors with `log.Fatal()`
1. Converting strings to numbers using functions in the `strconv` package
1. Blocks and variable scope
1. Package names versus import paths
1. Generating a random number
1. Getting input from the keyboard
1. Counting loops
1. Using `continue` and `break`

### Chapter 3: Functions

1. Using `fmt.Printf()` and `fmt.Sprintf()`
1. Declaring functions with parameters and return values
1. Using pointers for changeable parameters

### Chapter 4: Packages

1. Sharing code between multiple programs by creating a package
1. Package naming conventions
1. Constants

### Chapter 5: Arrays

1. Declaring array variables
1. Using array literals
1. Accessing array elements in a loop
1. Looping through an array using `range`
1. Reading a text file into an array

### Chapter 6: Slices

1. Declaring slice variables
1. Using slice literals
1. Using the slice operator to get part of an array or slice
1. Understand how a slice gives a view to all or part of an underlying array
1. Understand how multiple slices can view the same array, and what happens when the array changes
1. Adding to a slice with the `append()` function
1. Processing command-line arguments using the `os.Args` slice
1. Using variadic functions

### Chapter 7: Maps

1. Declaring map variables
1. Using map literals
1. Working with unassigned keys via zero values and boolean second return values
1. Using maps in loops with `range`

### Chapter 8: Structs

1. Declaring struct variables
1. Assigning struct fields using the dot operator
1. Defining struct types
1. Using pointers with struct parameters
1. Using struct literals
1. Embedding structs within structs

### Chapter 9: Defined Types

1. Defining numeric types to prevent unit errors
1. Converting between types using functions
1. Defining methods with receiver parameters
1. Using pointer receiver parameters

### Chapter 10: Encapsulation and Embedding

Know how to protect fields in a struct from invalid values by doing the following:

1. Put the struct in a package
1. Make the fields unexported by using lower-case names
1. Write setter methods that do not allow invalid values
1. Write getter methods to allow viewing the field values

### Chapter 11: Interfaces

1. Declaring an interface
1. Know how types fulfill an interface by defining the methods of the interface
1. Know how to use type assertions to access the concrete type
1. Know the use of the `error` and `Stringer` interfaces
1. Know how to use the empty interface

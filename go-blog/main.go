package main

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"example/custom"
)

func main() {
	fmt.Println("Hello main")

	// Variables
	// This is also known as declaration without initialization:
	//
	// var foo string
	// Declaration with initialization:
	//
	// var foo string = "Go is awesome"
	// Multiple declarations:
	//
	// var foo, bar string = "Hello", "World"
	// OR
	// var (
	//
	//	foo string = "Hello"
	//	bar string  = "World"
	//
	// )
	// Type is omitted but will be inferred:
	//
	// var foo = "What's my type?"
	// Shorthand declaration, here we omit var keyword and type is always implicit. This is how we will see variables being declared most of the time. We also use the := for declaration plus assignment.
	//
	// foo := "Shorthand!"
	// Note: Shorthand only works inside function bodies.
	//
	// Constants
	// We can also declare constants with the const keyword. Which as the name suggests, are fixed values that cannot be reassigned.
	//
	// const constant = "This is a constant"
	// It is also important to note that, only constants can be assigned to other constants.
	//
	// const a = 10
	// const b = a // Works
	//
	// var a = 10
	// const b = a // a (variable of type int) is not constant (InvalidConstInit)
	//
	//
	// Data Types:
	//
	// string - var value string = "string"
	// bool - var value bool = false
	// int
	// Floating point numbers
	// The default type for floating point values is float64.
	//
	//
	// Zero Values
	// numbers - init and float64 - are assigned as 0
	// boolean - bool as false
	// string as an empty string - ""
	//
	//
	// String Formatting
	//
	// fmt.Print - does not format anything. it simple takes a string and prints it.
	// fmt.Println - it adds a new line at the end and also inserts space between the arguments.
	// fmt.Printf - Print Formatter -  allows us to format numbers, strings, booleans and much more
	// annotation verbs - %s, %v, %f, %.2f etc - they tell the function how to format the arguments. we can control things like width, types, and precision with these
	//
	//
	// Flow Control
	// If/Else
	// if x > 5 {} else {}
	//
	// compact if
	// func main() { if x := 10; x > 5 { fmt.Println("x is gt 5") } }

	// Switch expression
	// unlike other langs, break statement is auto added at the end of each case.
	num := 4
	switch num {
	case 1:
		fmt.Println(num)
	case 2:
		fmt.Println(num)
	default:
		fmt.Println("Switch statement")
	}
	// switch also supports shorthand declarations like - switch num := 1; num { case ... }
	// fallthrough keyword - used to trasfer control to the next case even though the current case might have matched.

	// Loops
	//
	// so in go, we only have one type of loop which is the for loop
	for i := range 10 {
		fmt.Println(i)
	}
	for i := range 2 {
		fmt.Println(i)
	}

	// Break and continue
	// Go also supports both break and continue statements for loop control

	// while loop in go
	i := 0
	for i < 10 {
		i += 1
	}

	// Forever loop
	// func main() { for {} } - we omit the loop condition, and it loops forever

	// Functions
	// func myFunction(p1 string) {fmt.Print}
	//
	// func Test(a, b, c string) {
	// 	// go functions
	// 	fmt.Println(a)
	// }
	//
	// Go also supports multiple returns
	// func Test(a string, b bool) (string, bool) {return a, b}
	//
	// func Test2(a string, b string, c bool) (string, string, bool) {
	// 	// go functions return Multiple values
	// 	fmt.Println(a)
	// 	return a, b, c
	// }
	//
	// Named returns
	// another cool feature is named returns, where return values can be
	// named and treated as their own variables
	//
	// func Test3(p1 string) (s string, i int) {
	// 	// go functions - Named returns
	// 	s = fmt.Sprintf("%s function", p1)
	// 	i = 10
	// 	return
	// 	// added return stmt without any arguments
	// 	// this is also known as naked return
	// }
	//
	//
	//
	//
	// Functions as values
	// in Go functions are first class and we can use them as values
	fn := func() {
		fmt.Println("Functions are first class")
	}
	fn()
	//
	// we can also smplify this by making fn an anonymous function
	func() {
		fmt.Println("anonymous function")
	}()

	// Closures
	// a closure is a function value that references variables
	// from outside its body.
	myFn1 := func() func(int) int {
		sum := 0
		return func(v int) int {
			sum += v

			return sum
		}
	}

	add1 := myFn1()
	fmt.Println(add1(10))

	//
	//
	// Variadic Functions
	// are functions that can take zero or multiple arguments using
	// the ... ellipses operator
	add2 := func(values ...int) int {
		sum := 0

		for _, v := range values {
			sum += v
		}

		return sum
	}

	sum1 := add2(1, 2, 3, 4, 5, 6, 7) // ...
	fmt.Println(sum1)

	//
	//
	//
	//
	// Init
	// in Go, init is a special lifecycle function that is executed before the main func
	// Similar to main, the init function doe snot take any args nor returns any value.
	//
	// func init() {
	// 	fmt.Println("executed before main!")
	// }
	//
	// Unlike main, there can be more than one init function in sigle or multiple files.
	//
	// The init func is optional and is particularly used for any global setup which might be
	// essential for our program, such as establishing a db connection, fetching config files
	// or setting env variables, etc
	//
	//

	//
	//
	//
	// Defer
	// the defer keyword, lets us postpone the execution of a function until
	// the surrounding function returns.
	//
	// Can we use multiple defer functions? Absolutely, this brings us to what is know as
	// defer stack, eg.
	defer fmt.Println("I am finished")
	defer fmt.Println("Are you?")
	fmt.Println("Doing some work...")
	// result -> Doing some work...
	// 				-> Are you?
	// 				-> I am finished
	//
	// As we can see, defer statements are stacked and executed in a LIFO manner.
	// so defer is incredibly useful and is commmonly used for doing cleanup or error handling.
	//
	//
	//
	//

	//
	//
	// Modules
	//
	// a Module is a collection of Go packages stored in a file tree with a
	// go.mod file at its root, provided the directory is outside $GOPATH/src
	//
	// GOPATH is a variable that defines the root of your workspace and it contains
	// the following folders:
	// 		src: contains Go source code organized in a hierarchy.
	// 		pkg: contains compiled package code.
	// 		bin: contains compiled binaries and executables
	//
	//
	// create a new module using go mod init command which creates a new module
	// and initializes the go.mod file that describes it.
	// 		go mod init example
	//
	// if you want to add a new dependency, we will use go get command
	// 		go get github.com/rs/zerolog
	//
	// as we can see a go.sum file was also created. this file contains the expected
	// hashes of the content of the new modules.
	//
	// we can list all the dependencies using go list command as follows:
	// 		go list -m all
	//
	// If the dependency is not used, we can simply remove it using
	// 		go mod tidy # command
	//

	//
	//
	//
	//
	//
	//Packages
	//
	// A package is nothing but a directory containg one or more Go source
	// files, or other Go packages.
	// This meaning every Go source file must belong to a package and
	// package declaration is done at top of every source file as follows
	// 		package <<package_name>>
	//
	// so far we have done everything inside of package main. By convention,
	// executable programs (by that I mean the ones with the main package)
	// are called commands, others are simply called packages.
	//
	// The main package should also contain a main() function which is a special
	// function that acts as the entry point of an executable program.
	//
	// imports and exports
	// Basically, any value (like a variable or function) can be exported and visible
	// for other packages if they have been defined with an upper case identifier.
	// values declared on code.go file
	fmt.Println(custom.Value)

	//
	//
	//
	//
	//
	//
	// External dependencies
	// In Go, we are not only limited to working with local packages,
	// we can also install external packages using go install command
	//
	log.Print(custom.Value)
	// Go doesn't have a particular "folder structure" convention, always try to
	// organize your packages in a simple and intuitive way.

	//
	//
	//
	// Workspaces
	//
	//
	// - Workspaces allows us to work with multiple modules simultaneously
	// without having to edit go.mod files for each module
	// - each module within a workspace is treated as a root module when resolving
	// dependencies.
	//
	// to understand this better, let's start by creating a hello module.
	// mkdir workspaces && cd workspaces
	// mkdir hello && cd hello
	// go mod init hello

	//
	//
	//
	//
	//
	//
	//
	//
	//
	//Useful Commands
	//
	//	go fmt - formats the source code and it's enforced by the language
	//	go vet - which reports likely mistakes in our package. so if I go ahead and
	//	make a mistake in the syntax, and then run go vet, it should notify me
	//	of the errors
	//	go env - which simply prints all the go environment variables
	//	go fix - finds go programs that use old apis and rewrites them to use newer ones.
	//	go generate - is usually used for code generation.
	//	go install - compiles and installs packages and dependencies
	//	go clean - is used for cleaning files that are generated by compilers
	//
	// Build
	// Building static binaries is one of the best features of Go which enables us to
	// ship our code efficiently.
	// we can do this using the go build command.
	//
	//
	//
	//
	//
	//
	//
	//
	// Pointers
	//
	// a pointer is a variable that is used to store the memory address of another variable
	// It can be used like this:
	// var x *T -> where T is the type such as int, string, bool etc
	var x int
	var p *int
	x = 20
	p = &x         // we use the & ampersand operator to refer to a vars memory address
	fmt.Println(p) // prints the address
	// nil is a predeclared identifier in Go that represents zero value for
	// pointers, interfaces, channels, maps, and slices
	// just like the uninitialized int has a zero value of 0, a bool has false, and so on
	//
	//
	// Dereferencing
	// We can also use the * asterisk operator to retrieve the values stored in the
	// variable that the pointer points to.
	fmt.Println(*p) // prints the value stored inside that address

	// we can also change the value through the pointer
	*p = 10
	fmt.Println("after change:", x)

	//
	//
	//
	//
	// Pointers as function args
	// Pointers can also be used as arguments for a function when we need to pass
	// some data by reference.
	myFn2 := func(p *int) {
		fmt.Println(*p)
	}
	myFn2(&x)

	//
	//
	//
	//
	// Pointer to a Pointer
	// Here's an interesting idea... can we create a pointer to a pointer? The answer
	// is yes, we can
	p1 := &p
	fmt.Println("P value", *p, "address", p)
	fmt.Println("p1 value", p1, "address", p)
	fmt.Println("Dereferenced Value", **p1)
	// Note: It is important to know that pointers in Go do not support pointer
	// arithmetic like in C or C++
	//
	// However, we can compare two pointers of the same type for equality using
	// a == operator

	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	// Structs
	//
	// a struct is a user-defined type that contains a collection of named fields.
	// Basically, it is used to group related data together to form a single unit.
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}
	var person1 Person
	fmt.Println("Person 1:", person1) // Person 1: {  0}
	// all the struct fields are initialized with their zero values. so the FirstName
	// and LastName are set to "" empty string and Age is set to 0.

	// we can also initialize it as "struct literal".
	p2 := Person{
		FirstName: "Karan",
		LastName:  "Pratap Singh",
		Age:       22,
	}
	fmt.Println("Person 2:", p2) // {Karan Pratap Singh 22}
	// Accessing Fields
	fmt.Println("FirstName:", p2.FirstName)

	//
	//
	//
	//We can also create a pointer to a struct as well
	ptr1 := &p2
	// when we access, we dont need to explicitly Dereference the pointer.
	// Both statements are eqal
	fmt.Println((*ptr1).FirstName)
	fmt.Println(ptr1.FirstName)

	// As a side note, two structs are equal if all their corresponding
	// fields are equal as well
	ptr2 := Person{"a", "b", 20}
	ptr3 := Person{"a", "b", 20}
	fmt.Println(ptr2 == ptr3) // true

	//
	//
	//
	//
	// Exported fields
	// if a struct field is declared with a lower case identifier, it will not
	// be exported and only be visible to the package it is defined in.
	type Person1 struct {
		FirstName, LastName string
		Age                 int
		zipCode             string
	}
	// so the zipCode field won't be exported. Also, the same goes for the Person1
	// struct, if we rename it as person1, it won't be exported as well.

	//
	//
	//
	//
	//
	// Embedding and composition
	//
	// As we discussed earlier, Go doesn't necessarily support inheritance, but we
	// can do something similar with Embedding
	type SuperHero struct {
		Person
		Alias string
	}
	// So, our new struct will have all the properties of the original struct.
	// and it should behave the same as our normal struct. However, this is
	// usually not recommended and in most cases, composition is preferred. so
	// rather than Embedding, we will just define it as a normal field
	type SuperHero1 struct {
		Person Person
		Alias  string
	}

	//
	//
	//
	//
	// Struct tags
	//
	// A struct tag is just a tag that allows us to attach metadata
	// information to the field which can be used for custom behavior using
	// the reflect package.
	//
	type Animal struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	// You will often find tags in encoding packages, such as XML, JSON,
	// YAML, ORMs and Configuration management.
}

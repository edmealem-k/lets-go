package main

import "fmt"

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
	for i := 0; i < 10; i++ {
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
	// so defer is incredibly useful and is commmonly used foor doing cleanup or error handling.
	//
	//
	//
	//
	//
	//
}

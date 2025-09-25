package main

import "fmt"

func main() {

	//This is an anonymous function its get call instantly
	func() {
		fmt.Println("anonymous function")
	}() //this () is calling the function and it will execute instantly.

	//function expression: As function is also a type it is treated as "First class type"
	//We Can also store function

	//now x is a function type
	x := func(a int, b int) int {
		return a + b
	}
	var y int = x(10, 20)
	fmt.Println(y)

	z := temp()

	fmt.Println("Returning a function", z())
}

// in return type we are givng "func() int" as we are returning a function
// so we have to give its method signature
func temp() func() int {

	//Here returning a function like a variable
	return func() int {
		return 43
	}
}

// a named function with an identifier
// func (r receive) identifier(p parameter(s)) (r return(s)) { code }

// an anonymous function
// func(p parameter(s)) (r return(s)) { code }

package main

import "fmt"

func main() {

	/*defer foo() schedules the foo() function to run after main finishes, just before it returns.*/
	defer foo()
	bar()
}

// func (r receiver) identifier(p parameter(s)) (r return(s)) { code }

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

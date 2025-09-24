package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

/*
An interface in Go specifies a set of method signatures.
Any type that implements all the methods in the interface is said to "satisfy" the interface.
Here, any type with a speak() method (with no parameters and no return value) satisfies the human interface.
Usage in your code: Both person and secretAgent types have a speak() method, so they satisfy the human interface and can be used wherever a human is expected.
*/

type human interface {
	speak()
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }
func (p person) speak() {
	fmt.Println("I am", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am secret Agent", sa.first)
}

//This function can now have both person and secret Agent as parameter.
func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "James",
	}

	p2 := secretAgent{
		person: person{
			first: "jimmy",
		},
		ltk: true,
	}

	//p1 as type of person and p2 as type of secret Agent but same syntax this is polymorphism.
	saySomething(p1)
	saySomething(p2)
}

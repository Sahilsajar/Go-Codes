package main

import "fmt"

/*Task is to understand how reiver method work with combinations of struct, interface  with value and pointer semantic*/

/*Type dog will have a name and two methods walk and run and it will make to object 1 with value variable and another pointer variable*/
type dog struct {
	first string
}

//it is a value type receiver
func (d dog) walk() {
	fmt.Println("My name is " + d.first + " and I am walking")
}

//It is a pointer type receiver
func (d *dog) run() {
	d.first = "Robert" //changing name to understand clearly
	fmt.Println("My name is " + d.first + " and I am running")
}

//first look struct then go for interface as method called in main
type youngDog interface {
	walk()
	run()
}

func youngRun(y youngDog) {
	y.walk()
	y.run()
}

func main() {
	/*d1 a value and d2 is pointer type and it will able to call both function walk and run*/
	fmt.Println("Struct: ")
	d1 := dog{"Henry"}
	d1.walk()
	d1.run()

	d2 := &dog{"Podrick"}
	d2.walk()
	d2.run()

	/*for interface d1 will not able to call run method */
	//youngRun(d1) this will throw an error as value semnatic object cannot pointer semantic method through an interface
	fmt.Println("Interface: ")
	youngRun(d2)
}

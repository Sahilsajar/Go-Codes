package main

import "fmt"

func main() {
	//general channel
	c := make(chan int)
	//send channel
	cs := make(chan<- int)
	//receive channel
	cr := make(<-chan int)

	fmt.Println(c)
	fmt.Println(cs)
	fmt.Println(cr)

	cs = c //general to specifc works
	//c=cs specific to general also does not work
	//cs=cr specific to specific not possible
}

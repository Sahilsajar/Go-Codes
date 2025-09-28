package main

import "fmt"

func main() {

	//Here by putting one we stroring the value in a buffer
	//c := make(chan int, 1)
	c := make(chan int, 2) //for two values

	c <- 42 //Putting value 42 in channel c
	c <- 43

	fmt.Println(<-c) //taking out value
	fmt.Println(<-c)
}

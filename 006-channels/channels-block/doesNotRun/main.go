package main

import "fmt"

/*The below code will not run as in channels we have to put input and take output at the same time*/
func main() {
	c := make(chan int)

	c <- 42 //Putting value 42 in channel c

	fmt.Println(<-c) //taking out value
}

//Throw error that all go routines are asleep-deadlock
/*Now any two way we can run it or use channels
i. make another go routine
ii. use buffer memory
*/

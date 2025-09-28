package main

import "fmt"

//making Go routine
func main() {
	c := make(chan int)
	/* With help  of another go routine which is running concurrently so 1 one routine it will put the value and in other routine it will take out the vlaue*/
	go func() {
		c <- 42
	}()

	fmt.Println(<-c) //taking out value
}

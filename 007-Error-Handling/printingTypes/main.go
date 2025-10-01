package main

import (
	"fmt"
	"log"
	"os"
)

/*
task: how to handle error with diffent methods
*/
func main() {

	/*this return two types one is file type and other is error unlike other languages goes does not follow try catch system
	 */
	_, er := os.Open("no-file.txt")

	if er != nil {
		fmt.Println(er) //first just print it.
		log.Println(er) // unlike fmt this log print is used to also print in files
		//log.Fatal(er) this indicates fatal error means it will close the program and also does not run defer function
		//log.Panic(er) this closes current routine and then execute defer functions. panic(er) is just same as this.
	}
}

package main

import (
	"fmt"
	"log"
	"os"
)

/*task: to print error in txt file using logsetout and log.Println method*/

func main() {
	f, err := os.Create("log.txt")

	if err != nil {
		fmt.Println("Error in creating")
	}
	defer f.Close()

	//If error occured this will print the print ln text in f file
	log.SetOutput(f)

	_, er := os.Open("no-file.txt")

	if er != nil {
		log.Println("Error happend ", er)
	}
}

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

/*
	the task is to use Write function to write in both file as well as buffer

also byte and os library both have write method and write method also have an interface writer.
So also check polymorphism.
*/
type person struct {
	first []byte
}

// read about io package: https://pkg.go.dev/io#Writer
func (p person) writeOut(w io.Writer) error {

	_, err := w.Write(p.first)

	return err
}

func main() {

	p := person{
		first: []byte("Sahil"),
	}

	//os package from doc: https://pkg.go.dev/os
	f, err := os.Create("output.txt")

	if err != nil {
		//Log is same as fmt but it also print time stamp with output
		//I think fatal is method to print just error
		log.Fatalf("Got error")
	}
	defer f.Close()

	//Buffer is just a temporary memory to store data in go we can directly access that with buffer package.

	// read about bytes package note bytes is package and byte is a primitive data type.
	var b bytes.Buffer

	// why & because write method of buffer takes pointer
	//Polymorphism because buffer and os both have write method and its also a interface
	p.writeOut(&b)
	p.writeOut(f)

	fmt.Println(b.String())

}

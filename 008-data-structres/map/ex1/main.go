/*
Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of
TYPE []string which stores their favorite things. Store three records in your map. Print out all of
the values, along with their index position in the slice.
*/
package main

import "fmt"

func main() {
	mp := make(map[string][]string)

	for i := 0; i < 2; i++ {
		var key string
		fmt.Println("Enter last and first name as key")
		fmt.Scanln(&key)

		var slc []string
		for {
			var value string
			fmt.Println("Enter values: ")
			/*
				fmt.Scanln(&value) reads the whole line, so pressing Enter gives an empty string.
				fmt.Scan(&value) waits for a non-whitespace input
			*/
			fmt.Scanln(&value)
			if value == "" {
				break
			}
			slc = append(slc, value)
		}
		mp[key] = slc
	}

	fmt.Println(mp)
}

/*
->fmt.Scan reads input up to the first whitespace and leaves the newline character (\n) in the input buffer.
fmt.Scanln reads input until a newline and consumes the newline character.
If you use fmt.Scan followed by fmt.Scanln, the leftover newline from fmt.Scan causes fmt.Scanln to execute immediately, often reading an empty string.
*/

/*
Use a for range loop to loop over these words, then count how many times each word occurs.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter String")
	var s string
	fmt.Scanln(&s)

	//As input is in comma separated that is why converting the input into slice
	var arr []string
	arr = strings.Split(s, ",")
	mp := make(map[string]int)

	/*
			for i := range len(arr) {
		    mp[arr[i]]++
			}
			In Go, range doesn’t work directly on an integer like len(arr).
			You can only use range over arrays, slices, strings, maps, or channels.
	*/
	//By index
	// for i := range arr {
	// 	mp[arr[i]]++
	// }

	//By word
	for _, word := range arr {
		mp[word]++
	}

	fmt.Println(mp)
}

/*IMP
🧩 Why it fails for large input

fmt.Scanln() reads only until the first space or newline — it’s meant for short, single-word or short-line input.
So if you paste a long string (like thousands of comma-separated words), it stops early or truncates the input.

✅ Solution — Use bufio.NewReader

bufio.Reader can read entire lines from stdin safely, even very long ones.
*/

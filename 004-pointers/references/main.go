package main

import "fmt"

/*
Slices, Maps, Channels: Reference Types
Slices, maps, and channels are reference types in Go. that means when we pass it It doesnt make a new copy all changes are done on that only
Passing them to functions allows modification of the underlying data.
*/

func mapDelta(m map[string]int, s string) {
	m[s] = 50
}

func sliceDelta(sl []int) {
	sl[0] = 99
}

func main() {

	var s = []int{1, 2, 3, 4}

	fmt.Println("slice before", s)
	sliceDelta(s)
	fmt.Println("slice after", s)

	var mp = make(map[string]int)
	mp["sas"] = 35
	fmt.Println("Map before", mp)
	mapDelta(mp, "sas")
	fmt.Println("Map after", mp)

}

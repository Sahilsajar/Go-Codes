package main

import "fmt"

/* task example of callback function */
/*callback function: Is simply just passing of a funciton
 */
func main() {
	x := doMath(20, 10, add)
	var y int = doMath(20, 10, sub)

	fmt.Println(x)
	fmt.Println(y)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

//passing function like add and sub to doMath so passing function are call back function
func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

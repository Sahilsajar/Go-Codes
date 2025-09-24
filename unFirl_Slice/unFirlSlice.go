package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	/*However, when you pass a slice (arr), Go expects you to "unpack" the slice using the ... operator.*/
	fmt.Println("The sum is ", sum(arr...))
}

//Variadic Function
func sum(ii ...int) int {

	var sum int
	for _, val := range ii {
		sum += val
	}

	return sum
}

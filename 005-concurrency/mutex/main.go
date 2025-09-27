package main

import (
	"fmt"
	"runtime"
	"sync"
)

//check online for what is race around condition
//The below condition is race around condition as the counter value should be 100 at the end

func main() {
	var count int = 0
	fmt.Println("Num of Routines \t\t", runtime.NumGoroutine())

	var wg sync.WaitGroup

	wg.Add(100)
	defer wg.Wait()

	for i := 0; i < 100; i++ {
		go func() {
			var x = count
			x++
			count = x
			fmt.Println("Num of Routines \t\t", runtime.NumGoroutine())
			wg.Done()
		}()
	}

	fmt.Println("Count: ", count)
}

//Mutex is the solution for race around condtion

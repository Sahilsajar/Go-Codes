package main

import (
	"fmt"
	"runtime"
	"sync"
)

//check online for what is race around condition
//The below condition is race around condition as the counter value should be 100 at the end and with help of mutex lock it will always come 100

func main() {
	var count int = 0
	fmt.Println("Num of Routines \t\t", runtime.NumGoroutine())

	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			var x = count
			x++
			count = x
			mu.Unlock()
			runtime.Gosched()
			fmt.Println("Num of Routines \t\t", runtime.NumGoroutine())
			wg.Done()
		}()
	}

	wg.Wait() //put wait before  print so let all the concurrent code finished first
	fmt.Println("Count: ", count)
}

//Mutex is the solution for race around condtion

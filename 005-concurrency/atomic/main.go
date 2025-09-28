package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//check online for what is race around condition
//The below condition is race around condition as the counter value should be 100 at the end and instead of mutex we can also use atomic package of sync

func main() {
	var count int64 = 0
	fmt.Println("Num of Routines \t\t", runtime.NumGoroutine())

	var wg sync.WaitGroup

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&count, 1) //atomicatically handles race condtions
			runtime.Gosched()
			fmt.Println("Num of Routines \t\t", runtime.NumGoroutine())
			wg.Done()
		}()
	}

	wg.Wait() //put wait before  print so let all the concurrent code finished first
	fmt.Println("Count: ", count)
}

//Mutex is the solution for race around condtion

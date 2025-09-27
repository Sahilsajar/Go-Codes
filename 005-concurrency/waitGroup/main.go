package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*Task: write concurrent code */
//wg is package scope
var wg sync.WaitGroup

func main() {

	/*Runtime package : https://pkg.go.dev/runtime*/
	fmt.Println("OS: \t\t", runtime.GOOS)
	fmt.Println("Architecture \t\t", runtime.GOARCH)
	fmt.Println("Num of Cpu \t\t", runtime.NumCPU())
	/*Here routines are like threads in java
	If we dont like concurrent code then total routines will be 1
	Then as we increase the routine(thread) then it will increse*/
	fmt.Println("Num of Routines \t\t", runtime.NumGoroutine())

	/*with just adding go it foo function is running in concurrent that is why u cannot predict the output flow and as in node js you know that if we dont add wait the main function will close before printing the foo function*/
	// go foo()
	// boo()

	/*WaitGroup: https://pkg.go.dev/sync#WaitGroup*/

	//This wg.add(1) task is added and wait until it is completed
	//The completion mark is wg.Done()
	wg.Add(2) //to concurrent is going wait for 2 wg.done()
	defer wg.Wait()
	go foo()
	go boo()
	fmt.Println("Num of Routines \t\t", runtime.NumGoroutine())

}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo")
	}
	//now wait function will completed until its hit wg.done
	defer wg.Done()
}

func boo() {
	for i := 0; i < 10; i++ {
		fmt.Println("boo")
	}
	defer wg.Done()
}

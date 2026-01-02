package main

import (
	"fmt"
	"sync"
	"time"
)

// 5 philosopher
// 5 fork

type philosopher struct {
	name      string
	leftFork  int
	rightFork int
}

var eatTime time.Duration = 1 * time.Second
var thinkTime time.Duration = 2 * time.Second
var turnsToEat int = 3

//Philospher with two forks can eat

func main() {

	forks := []int{0, 1, 2, 3, 4}
	var wg sync.WaitGroup
	var seated sync.WaitGroup

	philosophers := []philosopher{{
		name:      "a",
		leftFork:  4,
		rightFork: 0,
	}, {
		name:      "b",
		leftFork:  0,
		rightFork: 1,
	}, {
		name:      "c",
		leftFork:  1,
		rightFork: 2,
	}, {
		name:      "d",
		leftFork:  2,
		rightFork: 3,
	}, {
		name:      "e",
		leftFork:  3,
		rightFork: 4,
	}}

	//we need map of locks so we can put lock in each fork
	forksMap := make(map[int]*sync.Mutex)

	for i := 0; i < len(forks); i++ {
		forksMap[i] = &sync.Mutex{}
	}

	wg.Add(len(philosophers))
	defer wg.Wait()

	seated.Add(len(philosophers))

	//Philosophers are cannot eat until they all are seated

	for i := 0; i < len(philosophers); i++ {
		go dine(&philosophers[i], forks, forksMap, &wg, &seated)
	}

}

func dine(p *philosopher, forks []int, forksMap map[int]*sync.Mutex, wg *sync.WaitGroup, seated *sync.WaitGroup) {
	defer wg.Done()
	seated.Done()

	//Waiting for philospher to be seated
	seated.Wait()
	//All professor are seated now we can start eating.

	for i := 0; i < turnsToEat; i++ {
		//Locking the fork before eating.
		//Lock can only be applied if it is unlock otherwise it will till it is unlock.
		//if lock put directly it will create a deadlock in condtion if all professor took there left fork first

		if p.leftFork > p.rightFork {
			forksMap[p.rightFork].Lock()
			fmt.Printf("\t %s have right fork \n", p.name)
			forksMap[p.leftFork].Lock()
			fmt.Printf("\t %s have left fork \n", p.name)
		} else {
			forksMap[p.leftFork].Lock()
			fmt.Printf("\t %s have left fork \n", p.name)
			forksMap[p.rightFork].Lock()
			fmt.Printf("\t %s have right fork \n", p.name)
		}

		fmt.Printf("\t %s is eating \n", p.name)
		time.Sleep(eatTime)

		time.Sleep(thinkTime)
		fmt.Printf("\t%s done eating \n", p.name)
		forksMap[p.leftFork].Unlock()
		forksMap[p.rightFork].Unlock()
	}

	fmt.Printf(" %s is full\n", p.name)
}

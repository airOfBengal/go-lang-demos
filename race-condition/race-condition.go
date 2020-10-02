package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0
var wg sync.WaitGroup

func increment() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		x++
		fmt.Println("incremented: ", x)
		time.Sleep(time.Millisecond * 100)
	}
}

func decrement() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		x--
		fmt.Println("decremented: ", x)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	wg.Add(2)

	// Race condition: It is defined as the condition when programs behavior depends on non-deterministic ordering.
	// Explanation: In this program race condition happens when one goroutine changes value of x but before printing this value,
	// another goroutine alters the value of x and unexpected output occurs.

	go increment()
	go decrement()
	wg.Wait()
}

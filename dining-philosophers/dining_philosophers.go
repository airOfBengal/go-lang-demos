package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	lChopStick, rChopStick *ChopStick
}

func (p Philosopher) eat(number int) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		p.lChopStick.Lock()
		p.rChopStick.Lock()

		fmt.Println("starting to eat ", number)
		time.Sleep(time.Second)

		p.rChopStick.Unlock()
		p.lChopStick.Unlock()

		fmt.Println("finishing eating ", number)
		time.Sleep(time.Second)
	}
}

func main() {
	chopSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = new(ChopStick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 1; i <= 5; i++ {
		philosophers[i-1] = &Philosopher{chopSticks[i-1], chopSticks[i%5]}
	}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go philosophers[i-1].eat(i)
	}

	wg.Wait()
}

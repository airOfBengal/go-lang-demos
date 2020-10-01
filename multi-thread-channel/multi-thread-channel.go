package main

import "fmt"

func multiply(a, b int, c chan int) {
	c <- a * b
}

func main() {
	c := make(chan int)

	go multiply(1, 2, c)
	go multiply(3, 4, c)
	go multiply(<-c, <-c, c)

	fmt.Println(<-c)
}

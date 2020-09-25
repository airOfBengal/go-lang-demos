package main

import "fmt"

func main() {
	var f float64

	fmt.Printf("Enter a floating point number: ")
	_, _ = fmt.Scan(&f)
	fmt.Printf("Truncated number is: %d\n", int(f))
}

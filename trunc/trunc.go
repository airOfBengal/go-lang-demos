package main

import "fmt"

func main() {
	var f float32

	fmt.Printf("Enter a floating point number: ")
	_, _ = fmt.Scan(&f)
	fmt.Printf("Truncated number is: %.f\n", f)
}

package main

import "fmt"

func GenDisplaceFn(a float64, v float64, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v*t + s
	}
}

func main() {
	var a, v, s float64
	fmt.Printf("Enter acceleration: ")
	_, _ = fmt.Scanf("%f", &a)

	fmt.Printf("Enter initial velocity: ")
	_, _ = fmt.Scanf("%f", &v)

	fmt.Printf("Enter initial displacement: ")
	_, _ = fmt.Scanf("%f", &s)

	fn := GenDisplaceFn(a, v, s)

	var t float64
	fmt.Printf("Enter time value: ")
	_, _ = fmt.Scanf("%f", &t)

	result := fn(t)

	fmt.Printf("Displacement after %f second(s): %f\n", t, result)
}

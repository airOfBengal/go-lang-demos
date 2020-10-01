package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var a, v0, s0, t float64

	prompt("acceleration", &a)
	prompt("initial velocity", &v0)
	prompt("initial displacement", &s0)
	prompt("time", &t)

	fn := GetDisplaceFn(a, v0, s0)

	fmt.Printf("Displacement is: %.2f", fn(t))
}

func prompt(parameter string, dest *float64) {
	fmt.Printf("Enter '%s': ", parameter)
	if _, err := fmt.Scanf("%f", dest); err != nil {
		log.Fatalf("Error, you typed not an float...\n")
	}
}

func GetDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		fmt.Printf("%f\n", (0.5*a*math.Pow(t, 2) + v0*t + s0))
		return (1/2)*a*math.Pow(t, 2) + v0*t + s0
	}
}

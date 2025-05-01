package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	z := 1.0
	for i := range 5 {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("%d = %v\n", i, z)
	}
	return z
}

func main() {

	sqrt := Sqrt(2)

	fmt.Println(sqrt)

	if sqrt == math.Sqrt2 {
		println("OK")
	} else {
		println("K0")
	}
}

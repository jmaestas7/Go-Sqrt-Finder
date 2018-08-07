package main

import (
	"flag"
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, string) {
	z := 1.0
	if x == 0 {
		return 0, "no solution possible, X is 0"
	}
	for {
		z -= (z*z - x) / (2 * z)
		if z == math.Sqrt(x) {
			return z, "is the solution found"
		}
	}
}

func main() {
	num := flag.Int("number", 2, "number to calculate into the squareroot equation")
	flag.Parse()

	fmt.Println(Sqrt(float64(*num)))
	fmt.Println("exit")
}

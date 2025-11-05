package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x / 2
	y := 0.0
	for z-y > 0.001 {
		y = z
		z -= (z*z - x) / (2 * z)
	}
	return z

}

func main() {
	fmt.Println(Sqrt(1))
}

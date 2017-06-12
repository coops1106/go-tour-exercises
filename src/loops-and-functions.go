package main

import (
	"fmt"
	"math"
)

const aSmallDelta = 0.00001

func Sqrt(x float64) (z float64) {
	z = x
	for i:=0;i<9;i++ {
		n := z - ((math.Pow(z, 2)-x) / (2*z))
		if (z - n) < aSmallDelta {
			break
		} else {
			z = n
		}
	}
	return
}

func main() {
	v := 10.0
	fmt.Println("Approx sqrt :", Sqrt(v))
	fmt.Println("Math sqrt :", math.Sqrt(v))
}

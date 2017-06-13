package main

import (
	"fmt"
	"math"
)

const aSmallDelta = 0.00001

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x
	for i:=0;i<9;i++ {
		n := z - ((math.Pow(z, 2)-x) / (2*z))
		if (z - n) < aSmallDelta {
			break
		}
		z = n
	}
	return z, nil
}

func main() {
	v := 10.0
	fmt.Println(Sqrt(v))
	fmt.Println(Sqrt(v * -1))
	fmt.Printf("Math sqrt: %g", math.Sqrt(v))
}

package main

import (
	"flag"
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	z := 1.0

	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("%v回目：", i)
		fmt.Println(z)
		if math.Sqrt(x) == z {
			return z, nil
		}
	}
	return z, nil
}

func (e ErrNegativeSqrt) Error() string {
	v := fmt.Sprint(float64(e))
	return fmt.Sprintf("cannot Sqrt negative number: %v", v)
}

func main() {
	f := flag.Float64("num", 2, "Set Sqrt args")
	flag.Parse()
	fmt.Println(Sqrt(*f))
	fmt.Println("----")
	fmt.Println(Sqrt(-*f))
}

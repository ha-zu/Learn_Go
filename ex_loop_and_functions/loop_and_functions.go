package main

import (
	"flag"
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("%v回目：%v\n", i, z)
		if math.Sqrt(x) == z {
			return z
		}
	}
	return z
}

func main() {
	f := flag.Float64("num", 2, "Set Sqrt args")
	flag.Parse()
	fmt.Println("own sqrt:", Sqrt(*f))
	fmt.Println("math pkg sqrt:", math.Sqrt(*f))
}

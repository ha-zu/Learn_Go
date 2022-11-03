package main

import "golang.org/x/tour/pic"

/*
Picメソッドのdx, dyはconstで指定してあるので、引数を渡すことは不可能
https://cs.opensource.google/go/x/tour/+/refs/tags/v0.1.0:pic/pic.go;l=27
*/

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)
	for y := range ret {
		ret[y] = make([]uint8, dx)
		for x := range ret[y] {
			ret[y][x] = uint8(x ^ y)
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}

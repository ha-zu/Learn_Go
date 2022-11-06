package main

import (
	"io"
	"os"
	"strings"
)

const (
	a     byte = byte('a')
	m     byte = byte('m')
	n     byte = byte('n')
	z     byte = byte('z')
	A     byte = byte('A')
	M     byte = byte('M')
	N     byte = byte('N')
	Z     byte = byte('Z')
	rot13 byte = 13
)

type rot13Reader struct {
	r io.Reader
}

//https://ja.wikipedia.org/wiki/ROT13#memfrob()%E9%96%A2%E6%95%B0
func (r rot13Reader) Read(bytes []byte) (num int, e error) {

	num, err := r.r.Read(bytes)
	if err != nil {
		return num, err
	}

	for i, val := range bytes {
		switch {
		//if a to m add 13
		case (a <= val && m >= val) || (A <= val && M >= val):
			bytes[i] += rot13
		//if n to z sub 13
		case (n <= val && z >= val) || (N <= val && Z >= val):
			bytes[i] -= rot13
		default:
			bytes[i] = val
		}
	}

	return num, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

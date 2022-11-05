package main

import (
	"golang.org/x/tour/reader"
	"io"
)

type MyReader struct{}

//confirm implemented io.Reader
var _ io.Reader = MyReader{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
//https://pkg.go.dev/io#Reader
//args: make([]byte, 1024, 2048)
func (mr MyReader) Read(b []byte) (n int, e error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}
func main() {
	//args: type Reader
	//https://cs.opensource.google/go/x/tour/+/refs/tags/v0.1.0:reader/validate.go;l=13
	reader.Validate(MyReader{})
}

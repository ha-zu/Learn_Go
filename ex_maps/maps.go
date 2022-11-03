package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	ret := make(map[string]int)
	wl := strings.Fields(s)
	for _, w := range wl {
		ret[w]++
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}

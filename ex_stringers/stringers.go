package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
//fmtのStringer interfaceのStringが実装された場合、優先されるよう
//pkgのPrintingのNo5で記載されている
//https://pkg.go.dev/fmt#hdr-Printing
//https://cs.opensource.google/go/go/+/refs/tags/go1.19.2:src/fmt/print.go;l=62
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

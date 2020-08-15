package main

import "fmt"

func main() {
	type NoKtp string
	type NoTlp int32
	type Merried bool

	var Noktpkku NoKtp = "001238123886889"
	var noTlpku NoTlp = 123
	var Status Merried = true

	fmt.Println(Noktpkku)
	fmt.Println(Status)
	fmt.Println(noTlpku)
}

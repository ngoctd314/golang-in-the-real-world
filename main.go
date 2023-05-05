package main

import (
	"fmt"
)

const (
	X float32 = 3.14
	Y         // here must be one identifier
	Z         // here must be one identifier

	A, B = "Go", "language"
	C, _
)

func main() {
	fmt.Println(X, Y, Z, A, B, C)
}

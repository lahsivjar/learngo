package main

import (
	"fmt"
)

// Global const
const s = "constant"

func main() {
	// Local const (NOTE: type inference)
	const a = 1000

	// For constant operations are performed with arbitrary precision
	const d = a / 3;

	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(d)
}

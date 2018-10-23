package main

import "fmt"

func main() {

	// No brackets, yay~
	if true {
		fmt.Println("Told ya")
	} else {
		fmt.Println("So lonely here")
	}

	// To many if-else, let's try switch-case
	x := 10
	switch x {
	case 0:
	case 3:
	case 5:
	case 10:
		fmt.Println("I am the chosen one")
	case 11:
		// No fallthrough
		fmt.Println("Missed by one :cry_face:")
	default:
		// Optional
	}
}

package main

import "fmt"

func main() {
	// Normal for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// While alternative (since Go has only for loop)
	for {
		fmt.Println("Hello there")
		break;
	}
}

package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}

	// Find the sum of the array using range
	sum := 0
	for _, val := range nums {
		sum += val
	}
	fmt.Printf("Sum is %d\n", sum)

	// Let's see range on maps next
	kv := map[string]int{"one": 1, "two": 2}
	fmt.Println(kv)

	for key, val := range kv {
		fmt.Printf("Key: %s, Value: %d\n", key, val)
	}

	// Iterate over key only
	for key := range kv {
		fmt.Printf("Key: %s\n", key)
	}

	// Iterate over string
	for i, c := range "superman" {
		fmt.Printf("Char at index %d is %c\n", i, c)
	}
}

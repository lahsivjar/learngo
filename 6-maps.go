package main

import "fmt"

func main() {
	// Declaring and assigning
	tempMap := map[string]int{"one": 1, "two": 2}
	tempMap["three"] = 3

	fmt.Println(tempMap)

	// Getting value from map
	val, state := tempMap["one"]
	fmt.Println("Value: ", val, "state: ", state)

	delete(tempMap, "one")
	// NOTE: `state` is required to check if the value is present in map or not
	val, state = tempMap["one"]
	fmt.Println("Value: ", val, "state: ", state)

	fmt.Println(tempMap)

	// Declaring map with map as values
	tempMap2 := make(map[string]map[int]string)

	tempMap2["one-two"] = map[int]string{1: "one", 2: "two"}

	tempMap2["three-four"] = make(map[int]string)
	tempMap2["three-four"][3] = "three"
	tempMap2["three-four"][4] = "four"

	fmt.Println(tempMap2)

	// Iterating over array
	// Since in Go every variable has to be used if there is a case
	// where we don't need the value of the variable we can use `_`
	for _, val := range tempMap2 {
		for innerKey, innerVal := range val {
			fmt.Println("Inner key: ", innerKey, "inner val: ", innerVal)
		}
	}
}

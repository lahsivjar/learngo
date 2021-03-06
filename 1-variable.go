package main

import "fmt"

func main() {
	// Declare primitive (declaration, type inference and assignment all at once)
	str := "Hello Go~"

	// fmt is standard go library package, used for printing stuff
	fmt.Println(str)

	// In go all variables must be used, to override use `_`
	// unusedVar := "This is an error if unusedVar is not used"
	_ = "This is okay"

	// Declaration
	var str2 string
	str2 = "Here we go again"
	fmt.Println(str2)

	// Another way to declare and assign
	var str3 = "Another hello"
	fmt.Println(str3)

	// Multiple declaration (NOTE: both `a` and `b` are declared as int type)
	var a, b int
	b = 5
	a = 10
	fmt.Println(a)
	fmt.Println(b)
}

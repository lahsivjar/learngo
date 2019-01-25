package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	lens := len(s)
	for i := 0; i <= lens/2; i++ {
		if s[i] != s[lens-i-1] {
			return false
		}
	}
	return true
}

func main() {
	strs := []string{"not_a_palindrome", "mom", "Detartrated", "sxxs"}

	for _, s := range strs {
		fmt.Printf("%s: %t\n", s, isPalindrome(strings.ToLower(s)))
	}
}

package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {

	normalizedStr := strings.ReplaceAll(strings.ToLower(s), " ", "")
	length := len(normalizedStr)

	// Membandingkan karakter dari depan dan belakang
	for i := 0; i < length/2; i++ {
		if normalizedStr[i] != normalizedStr[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var input string

	fmt.Print("Masukkan kata atau kalimat: ")
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Println("Ya, ini adalah palindrom.")
	} else {
		fmt.Println("Tidak, ini bukan palindrom.")
	}
}

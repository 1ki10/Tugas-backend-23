package main

import "fmt"

func isPalindrome(s string) bool {
	length := len(s)

	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var input string

	fmt.Print("Masukkan kata: ")
	fmt.Scanln(&input)

	// Memeriksa apakah input adalah palindrom
	if isPalindrome(input) {
		fmt.Println("Ya, ini adalah palindrom.")
	} else {
		fmt.Println("Tidak, ini bukan palindrom.")
	}
}

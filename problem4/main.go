package main

import "fmt"

// fungsi untuk mengenkripsi teks dengan penggeseran sebanyak 10 huruf
func encrypt(text string) string {
	
	encryptedText := ""
	
	for _, char := range text {
		
		if char >= 'A' && char <= 'Z' {
			newPos := (char-'A'+10)%26 + 'A'
			encryptedText += string(newPos)
		} else {
			encryptedText += string(char)
		}
	}
	return encryptedText
}

func main() {
	// Contoh teks yang ingin dienkripsi
	input := "GOLANG"
	
	fmt.Println("Original: ", input)
	fmt.Println("Encrypted: ", encrypt(input))
}
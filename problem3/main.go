package main

import "fmt"

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var number int

	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&number)

	// Memeriksa apakah bilangan tersebut adalah bilangan prima
	if isPrime(number) {
		fmt.Println(number, "adalah bilangan prima.")
	} else {
		fmt.Println(number, "bukan bilangan prima.")
	}
}

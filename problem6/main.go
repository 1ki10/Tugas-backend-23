package main

import "fmt"

func main()  {
	// full prima 2 3 5 6 13 23
	var input = 27
	var isPrima = true
	var isFullPrima = true
	for  i := 2; i < input; i++ {
		if input%i == 0 {
			isPrima = true
			isFullPrima = false
		}
	}
	for input > 0 && isPrima {
		var proses = input % 10
		if proses < 2 {
			isFullPrima = false
		}
		for  i := 2; i < proses; i++ {
			if proses%i == 0 {
				fmt.Println(input, proses)
				isFullPrima = false
			}
		}
		input /= 10
	}

	if isFullPrima {
		fmt.Println("Full Prima")
	} else {
		fmt.Println("Bukan Full Prima")
	}
}
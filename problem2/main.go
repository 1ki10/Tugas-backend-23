package main

import "fmt"

func DrawXYZ(N int) string {
	result := ""
	total := N * N

	for i := 1; i <= total; i++ {
		// Menentukan karakter yang akan ditambahkan
		if i%3 == 0 {
			result += "X"
		} else if i%2 == 0 {
			result += "Z"
		} else {
			result += "Y"
		}

		// Menambahkan baris baru di akhir setiap baris
		if i%N == 0 {
			result += "\n"
		}
	}

	return result
}

func main() {
	fmt.Print(DrawXYZ(3))
	// Output:
	// YZX
	// XYZ
	// YZX

	fmt.Print(DrawXYZ(5))
	// Output:
	// YZXZY
	// XYZXZ
	// YXYZX
	// ZYXYZ
	// XZYXY
}

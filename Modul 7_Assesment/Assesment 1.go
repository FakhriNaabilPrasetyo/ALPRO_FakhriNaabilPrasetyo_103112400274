package main

import (
	"fmt"
)

// Function to sum numbers from x to y
func sumRange(x int, y int) int {
	// Initialize sum variable
	sum := 0

	// Loop from x to y and accumulate the sum
	for i := x; i <= y; i++ {
		sum += i
	}

	return sum
}

func main() {
	// Input values for x and y
	var x, y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	// Check if x is less than or equal to y
	if x <= y {
		result := sumRange(x, y)
		fmt.Printf("Jumlah dari %d sampai %d adalah: %d\n", x, y, result)
	} else {
		fmt.Println("Nilai x harus lebih kecil atau sama dengan y.")
	}
}

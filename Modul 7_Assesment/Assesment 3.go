package main

import "fmt"

func main() {
	var qirat int
	fmt.Print("Masukkan jumlah uang dalam satuan qirat: ")
	fmt.Scanln(&qirat)

	// Konversi qirat ke fals
	fals := qirat / 6
	sisaQirat := qirat % 6

	// Konversi fals ke dirham
	dirham := fals / 10
	sisaFals := fals % 10

	// Konversi dirham ke dinar
	dinar := dirham / 10
	sisaDirham := dirham % 10

	fmt.Printf("Konversi uang:\n")
	fmt.Printf("Dinar: %d\n", dinar)
	fmt.Printf("Dirham: %d\n", sisaDirham)
	fmt.Printf("Fals: %d\n", sisaFals)
	fmt.Printf("Qirat: %d\n", sisaQirat)
}

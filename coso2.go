package main

import "fmt"

func main() {
	var alas, tinggi, luas float64
	fmt.Print("Masukan Alas: ")
	fmt.Scan(&alas)
	fmt.Print("tinggi: ")
	fmt.Scan(&tinggi)
	luas = (alas * tinggi / 2)
	fmt.Print(luas)
}
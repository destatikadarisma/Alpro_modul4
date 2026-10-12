package main

import "fmt"

func main() {
	var totalawal, diskon, totalakhir float64
	fmt.Scan(&totalawal)
	fmt.Scan(&diskon)
	totalakhir = totalawal - (totalawal * diskon / 100)
	fmt.Printf("%.0f\n", totalakhir)
}

//program hargadiskon
//kamus
// 	totalawal, diskon, totalakhir : real
//algoritma
// 	input(totalawal)
// 	input(diskon)
// 	totalakhir <- totalawal - (totalawal * diskon / 100)
// 	output(totalakhir)
//endprogram

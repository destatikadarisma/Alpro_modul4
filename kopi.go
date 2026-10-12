package main

import "fmt"

func main() {
	const harga_kopi = 20000
	const pajak = 0.11

	var jumlahkopi int
	var totalsebelumpajak, totalpajak, totalbayar float64

	fmt.Scan(&jumlahkopi)

	totalsebelumpajak = float64(jumlahkopi) * harga_kopi
	totalpajak = totalsebelumpajak * pajak
	totalbayar = totalsebelumpajak + totalpajak

	fmt.Printf("%.0f\n", totalbayar)
}

//program hitungbayar
//kamus
//	const harga_kopi : real = 20000
//	const pajak : real = 0.11
//	jumlahkopi : integer
//	totalsebelumpajak, totalpajak, totalbayar : real
//algoritma
//	input(jumlahkopi)
//	totalsebelumpajak <- jumlahkopi * harga_kopi
//	totalpajak <- totalsebelumpajak * pajak
//	totalbayar <- totalsebelumpajak + totalpajak
//	output(totalbayar)
//endprogram

package main

import "fmt"

func main() {
	var h, batas, tekanan float64
	const p float64 = 1025.0 //"const" angka tetap, tidak bisa di ubah
	const g float64 = 9.81

	fmt.Scan(&h, &batas)
	tekanan = p * g * h
	fmt.Println(tekanan)
	fmt.Println(tekanan <= batas)
}

//program penyelam
// 	h, batas, tekanan : real
// 	constanta p : real <- 1025.0
//	constanta g : real <- 9.81
//algoritma
// 	tekanan <- p * g * h
// 	output (tekanan)
// 	output (tekanan <= batas)
//endprogram

package main

import "fmt"

func main() {
	var bilangan, d1, d2, d3 int
	fmt.Scan(&bilangan)

	d1 = bilangan / 100
	d2 = (bilangan % 100) / 10
	d3 = bilangan % 100 % 10
	fmt.Println(d1 <= d2 && d2 <= d3)
}

//program urutandigit
// 	bilangan, d1, d2, d3 : integer
//algoritma
// 	d1 <- bilangan / 100
// 	d2 <- (bilangan % 100) / 10
// 	d3 <- bilangan % 100 % 10
// 	output (d1 <= d2 && d2 <= d3)
//endprogram
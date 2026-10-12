package main

import ("fmt"
"math")

func main() {
	var x1, y1, x2,y2, x3, y3 float64
	var ab, bc, ca,terpanjang float64

	fmt.Scan(&x1, &y1)
	fmt.Scan(&x2, &y2)
	fmt.Scan(&x3, &y3)

	ab = math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
	bc = math.Sqrt((x3-x2)*(x3-x2) + (y3-y2)*(y3-y2))
	ca = math.Sqrt((x1-x3)*(x1-x3) + (y1-y3)*(y1-y3))

	terpanjang = ab
	if bc > terpanjang {
		terpanjang = bc
	}
	if ca > terpanjang {
		terpanjang = ca
	}

	fmt.Printf("%.2f\n", terpanjang)
}

package main

import "fmt"

func main() {
	var c1, c2 rune
	var issymbol1, issymbol2 bool

	fmt.Scanf("%c%c", &c1, &c2)
	issymbol1 = !((c1 >= 'A' && c1 <= 'Z') || (c1 >= 'a' && c1 <= 'z') || (c1 >= '0' && c1 <= '9'))
	issymbol2 = !((c2 >= 'A' && c2 <= 'Z') || (c2 >= 'a' && c2 <= 'z') || (c2 >= '0' && c2 <= '9'))

	fmt.Println(issymbol1, issymbol2)

}

//cek simbol
//kamus
//	c1, c2 : karakter
//	issymbol1, issymbol2 : boolean
//algoritma
//	input(c1, c2)
//	issymbol1 <- not((c1 >= 'A' and c1 <= 'Z') or (c1 >= 'a' and c1 <= 'z') or (c1 >= '0' and c1 <= '9'))
//	issymbol2 <- not((c2 >= 'A' and c2 <= 'Z') or (c2 >= 'a' and c2 <= 'z') or (c2 >= '0' and c2 <= '9'))
//	output(issymbol1, issymbol2)
//endprogram
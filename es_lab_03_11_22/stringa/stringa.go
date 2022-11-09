package main

import (
	"fmt"
)

func main() {
	/*
		Scrivere un programma stringa.go che legge da standard input una stringa e la analizza carattere per
		carattere: stampa "+" quando il carattere considerato è maggiore del precedente, stampa "-"
		quando è minore e stampa "=" quando è uguale.
	*/

	var s string
	var curr, prec rune

	fmt.Scan(&s)

	prec = rune(s[0])

	for i := 1; i < len(s); i++ {
		curr = rune(s[i])
		switch {
		case curr > prec:
			fmt.Print("+")
		case curr == prec:
			fmt.Print("=")
		case curr < prec:
			fmt.Print("+")
		}
		prec = curr
	}

}

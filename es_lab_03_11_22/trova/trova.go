package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		Scrivere un programma trova.go che legge da standard input un carattere (runa) e una
		stringa e stampa la posizione del carattere nella stringa (la prima volta che appare),
		o -1 se il carattere non c'è.

	*/

	var s string
	var char rune
	var index int

	fmt.Scanf("%c", &char)
	fmt.Scan(&s)

	index = strings.IndexRune(s, char)

	if index != -1 {
		fmt.Println("il carattere", string(char), "è in posizione", index)
	} else {
		fmt.Println(index)
	}

}

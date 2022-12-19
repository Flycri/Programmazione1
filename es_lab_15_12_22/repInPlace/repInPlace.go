package main

import (
	"fmt"
	"os"
)

/*
	Scrivere un programma `repInPlace.go` che definisca una funzione:
	func repInPlace(stringa []rune, old rune, new rune)
	che modifichi la `stringa` passata sostituendo ogni occorrenza della runa `old` con la runa `new`

	Aggiungere inoltre un main che accetti tre parametri a linea di comando, rispettivamente:
	- stringa da modificare
	- carattere sostituendo
	- carattere di rimpiazzo
*/

func repInPlace(stringa []rune, old rune, new rune) {
	for k, r := range stringa {
		if r == old {
			stringa[k] = new
		}
	}
}

func main() {
	s := []rune(os.Args[1])
	old := rune(os.Args[2][0])
	new := rune(os.Args[3][0])
	repInPlace(s, old, new)
	fmt.Println(string(s))
}

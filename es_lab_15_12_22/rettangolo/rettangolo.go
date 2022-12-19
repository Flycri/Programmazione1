package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
	Scrivere un programma rettangolo.go, dotato di:
	una struttura Rettangolo, con campi base e altezza, in quest'ordine, di tipo int
	un metodo String() per Rettangolo che restituisce il disegno del rettangolo, pieno e disegnato con '.'.
	una funzione main() che, dati due numeri naturali come
	argomenti sulla linea di comando, emetta nel flusso d'uscita il disegno del rettangolo che ha quei valori per la base
	e l'altezza (utilizzando il metodo String).
*/

type Rettangolo struct {
	base    int
	altezza int
}

func (r Rettangolo) String() string {
	var s string
	if r.altezza == 0 || r.base == 0 {
		fmt.Println("Rettangolo con una delle due misure parti a 0")
	} else {
		for i := 0; i < r.altezza; i++ {
			for j := 0; j < r.base; j++ {
				s += "."
			}
			s += "\n"
		}
	}
	return s
}

func main() {

	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])

	r := Rettangolo{base: x, altezza: y}

	s := r.String()

	fmt.Print(s)

}

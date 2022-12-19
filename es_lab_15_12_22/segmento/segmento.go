package main

import (
	"fmt"
)

/*
	Scrivere un programma segmento.go, dotato di:

	una struttura Segmento, con campi

	estremi byte
	interno byte
	orizzontale bool (false sta per verticale)
	lunghezza int

	dichiarati in quest'ordine

	un metodo String() per Segmento che restituisce una rappresentazione grafica del segmento. Ad esempio un segmento
	seg := Segmento{'*', '-', true, 5}
	è rappresentato graficamente dalla stringa *---*

	un metodo allunga(n int) per *Segmento che incrementa di n la lunghezza del segmento (o la accorcia se n è negativo)

	una funzione main() che legge da standard input:

	il carattere per gli estremi
	il carattere interno
	true o false per orizzontale o vertivale
	la lunghezza del segmento

	ed emetta nel flusso d'uscita la rappresentazione grafica del segmento con quelle caratteristiche.
	Poi allunghi quello stesso segmento di 5 e di nuovo ne emetta nel flusso d'uscita la rappresentazione grafica.
*/

type Segmento struct {
	estremi     byte
	interno     byte
	orizzontale bool
	lunghezza   int
}

func (s Segmento) String() string {
	var output string
	output += string(s.estremi)
	if !s.orizzontale {
		output += "\n"
	}
	for i := 0; i < s.lunghezza-2; i++ {
		output += string(s.interno)
		if !s.orizzontale {
			output += "\n"
		}
	}
	output += string(s.estremi)
	return output
}

func (s *Segmento) Allunga(n int) {
	s.lunghezza += n
}

func main() {

	var segmento Segmento

	fmt.Println("Inserire carattere per gli estremi")
	fmt.Scanf("%c\n", &segmento.estremi)
	fmt.Println("Inserire carattere per l'interno")
	fmt.Scanf("%c\n", &segmento.interno)
	fmt.Println("Inserire true per orizzontale, false per verticale")
	fmt.Scanf("%t\n", &segmento.orizzontale)
	fmt.Println("Inserire la lunghezza del segmento")
	fmt.Scanf("%d\n", &segmento.lunghezza)

	fmt.Println(segmento.String())
	segmento.Allunga(5)
	fmt.Println(segmento.String())

}

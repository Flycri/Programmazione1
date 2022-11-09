package main

import "fmt"

func main() {
	/*
		Scrivere un programma lunghezza_tot.go che legge da standard input un int totLen e una sequenza
		di stringhe (una per riga) sommandone le lunghezze fino a raggiungere (o superare) totLen.
		Raggiunto totLen, il programma stampa la somma delle lunghezze e la concatenazione delle
		stringhe lette.
	*/

	var s, totString string
	var totLen int

	fmt.Scan(&totLen)
	fmt.Scan(&s)
	totString = s

	for len(totString) < totLen {
		fmt.Scan(&s)
		totString = totString + s
	}

	fmt.Println("somma delle lunghezze delle stringhe:", len(totString), "\nStringa concatenata totale: ", totString)

}

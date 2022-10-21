package main

import "fmt"

func main() {
	/*
		Scrivere un programma data_valida1.go che legge due interi gg e mm rappresentanti giorno e mese dell’anno,
		e verifica che sia una data valida. In questa versione si assuma che tutti i mesi abbiano 31 giorni.
		Il programma stampa "data valida" / "data non valida".
	*/

	var gg, m int

	fmt.Scan(&gg)
	fmt.Scan(&m)

	if ((gg >= 1) && (gg <= 31)) && ((m >= 1) && (m <= 12)) {
		fmt.Println("data valida")
	} else {
		fmt.Println(("data non valida"))
	}
}

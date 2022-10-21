package main

import "fmt"

func main() {
	/*
		Scrivere una seconda versione data_valida2.go che legge due interi gg e mm rappresentanti giorno e mese dell’anno, e verifica che sia una data valida.
		In questa seconda versione si tenga conto del fatto che solo i mesi 1, 3, 5, 7, 8, 10, 12 hanno 31 giorni,
		che i mesi 4, 6, 9, 11 ne hanno 30, e si assuma che febbraio ne abbia sempre 28.
		Il programma stampa "data valida" / "data non valida".
	*/

	var gg, m int

	fmt.Scan(&gg)
	fmt.Scan(&m)

	if (m >= 1) && (m <= 12) {
		if ((m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10 || m == 12) && ((gg >= 1) && (gg <= 31))) || ((m == 4 || m == 6 || m == 9 || m == 11) && ((gg >= 1) && (gg <= 30))) || ((m == 2) && ((gg >= 1) && (gg <= 28))) {
			fmt.Println("data valida")
		} else {
			fmt.Println("data non valida")
		}
	} else {
		fmt.Println("data non valida")
	}

}

package main

import "fmt"

func main() {
	/*
		Scrivere un programma guasti.go che riceve in ingresso tre interi che segnalano eventuali guasti di un impianto.
		Per ciascun codice, se è diverso da 0, il programma deve stampare il guasto segnalato secondo la tabella che segue:(vedere .md lab02)
	*/

	var e1, e2, e3 int

	fmt.Scan(&e1)
	fmt.Scan(&e2)
	fmt.Scan(&e3)

	if e1 == 1 {
		fmt.Println("caricamento acqua")
	} else if e1 == 2 {
		fmt.Println("scarico acqua")
	} else if e1 == 3 {
		fmt.Println("sistema di riscaldamento")
	}

	if e2 == 1 {
		fmt.Println("caricamento acqua")
	} else if e2 == 2 {
		fmt.Println("scarico acqua")
	} else if e2 == 3 {
		fmt.Println("sistema di riscaldamento")
	}

	if e3 == 1 {
		fmt.Println("caricamento acqua")
	} else if e3 == 2 {
		fmt.Println("scarico acqua")
	} else if e3 == 3 {
		fmt.Println("sistema di riscaldamento")
	}
}

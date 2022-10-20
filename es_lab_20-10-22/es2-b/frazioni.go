package main

import "fmt"

func main() {
	/*
	Scrivere un programma frazioni.go che verifica se due frazioni num1/den1 e num2/den2 sono equivalenti
	e stampa "equivalenti" o "non equivalenti", a seconda del caso.
	Usare il tipo int per i dati in input.
	Trovate una soluzione che non faccia uso di float64.
	*/
	var num1,num2,den1,den2 int
	fmt.Scan(&num1)
	fmt.Scan(&den1)
	fmt.Scan(&num2)
	fmt.Scan(&den2)
	if ((num1*den2)==(den1*num2)) {
		fmt.Println("le due frazioni sono equivalenti")
	} else {
		fmt.Println("le due frazioni non sono equivalenti")
	}
}
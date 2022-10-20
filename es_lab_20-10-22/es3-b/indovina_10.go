package main

import "fmt"

func main() {
	/*
	Scrivere un programma indovina_10.go che fissa nel programma un numero intero (a vostra scelta)
	tra 1 e 10 da indovinare, legge un intero da standard input e:

	se il numero in input è fuori dall’intervallo [1-10], stampa “Valore non valido”;
	se il numero è quello fissato, stampa “Hai indovinato!”;
	altrimenti stampa “Non hai indovinato”.
	*/

	var x int
	const NUM=7

	fmt.Scan(&x)
	if ((x>=1)&&(x<=10)) {
		if (x==NUM) {
			fmt.Println("Hai indovinato!")
		} else {
			fmt.Println("Non hai indovinato")
		}
	} else {
		fmt.Println("Valore non valido")
	}

}

package main

import (
	"fmt"
)

func main() {
	/*
	Scrivere un programma tariffe.go che chiede all’utente l’età (int) e se è studente (bool)
	e stampa il costo del biglietto di ingresso secondo la tabella che segue.
	Provare a risolvere l'esercizio utilizzando solo 4 (e non di più) istruzioni
	di assegnamento a una variabile tariffa e una sola istruzione di stampa finale per la stampa del costo.
	*/
	var age int
	var isStud bool

	fmt.Scan(&age)
	fmt.Scan(&isStud)

	if (age>=0)&&(age<9) {
		fmt.Println("gratis")
	}
	if ((age>=9)&&(age<18))||(age>=18)&&(isStud){
		fmt.Println("5")
	}
	if ((age>=18)&&(age<26))||(age>=65)&& !isStud {
		fmt.Println("7.5")
	}
	if ((age>=26)&&(age<65))&& !isStud {
		fmt.Println("10")
	}
}
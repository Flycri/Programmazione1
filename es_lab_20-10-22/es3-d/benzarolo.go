package main

import (
	"fmt"
)

func main() {
	/*
	Scrivere un programma benzarolo.go che calcola il prezzo di un pieno,
	dati litri (float64) e tipo di carburante (int, 0=benzina, 1=diesel, 2=alcol, 3=cherosene)
	secondo la seguente tabella dei prezzi: (vedere lab)
	*/
	var litri float64
	var carb int
	const BENZINA,DIESEL,ALCOL,CHEROSENE = 1.78,1.98,1.2,1.1

	fmt.Scan(&litri)
	fmt.Scan(&carb)
	if (carb==0) {
		fmt.Println(litri*BENZINA)
	}
	if (carb==1) {
		fmt.Println(litri*DIESEL)
	}
	if (carb==2) {
		fmt.Println(litri*ALCOL)
	}
	if (carb==3) {
		fmt.Println(litri*CHEROSENE)
	}
}

package main

import (
	"fmt"
)

func main() {
	/*
	Scrivere un programma fasi_giornata.go che che legga un orario (intero) e stampa la fase della giornata
	corrispondente, secondo la tabella che segue, o "orario non valido" se
	l'intero è fuori dall'intervallo 0-23
	*/
	var h int

	fmt.Scan(&h)

	if (h<0)||(h>23){
		fmt.Println("Orario non valido")
	} else {
		if (h>=0)&&(h<=6) {
			fmt.Println("notte")
		} else if (h>=7)&&(h<=13) {
			fmt.Println("mattino")
		} else if (h>=14)&&(h<=18) {
			fmt.Println("pomeriggio")
		} else if (h>=19)&&(h<=23) {
			fmt.Println("sera")
		}
	}
}
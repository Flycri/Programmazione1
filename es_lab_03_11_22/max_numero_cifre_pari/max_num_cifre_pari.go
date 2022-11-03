package main

import "fmt"

func main() {
	/*
		Scrivere un programma max_num_cifre_pari.go che, data una sequenza di numeri 
		(da leggere come stringhe), stampa il numero di cifre pari contenute nel numero che 
		ne contiene più di tutti.
	*/

	var s string
	var maxpari, npari int

	for {
		fmt.Scan(&s)
		for _,r:=range s {
			fmt.Println(rune(r)-48)
			if (int(rune(r)-48)%2==0) {
				npari++
			}
		}
		if npari>maxpari {
			maxpari=npari
			npari = 0
		}
		fmt.Println("numero massimo di cifre pari:",maxpari)
	}

	/* 
		NB: funziona SOLO perchè il codice unicode associato alle cifre è pari o 
		dispari anch'esso in maniera
		tale che, ad esempio, '1' (dispari) -> 49 (dispari) FIXED: cambiato in maniera + elegante
	*/
}
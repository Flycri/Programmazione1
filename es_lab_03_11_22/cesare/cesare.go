package main

import "fmt"

func main() {
	/*
		Scrivere un programma cesare.go che legge da standard input un valore intero non negativo
		k (la chiave di cifratura) e una sequenza di lettere minuscole consecutive (sulla stessa riga
		e senza spazi) terminate da invio. Il programma stampa la sequenza letta cifrata secondo il
		cifrario di Cesare, usando come chiave k (quella fornita dall'utente):
		ogni lettera del testo in chiaro è sostituita nel testo cifrato dalla
		lettera che si trova k posizioni dopo nell'alfabeto, ritornando alla lettera a dopo la zeta.
	*/

	var k int
	var s string
	var newchar rune

	fmt.Scan(&k)
	fmt.Scan(&s)

	for _, r := range s {
		newchar = rune(r)
		for i := 0; i < k; i++ {
			newchar++
			if newchar > 'z' {
				newchar = 'a'
			}
		}
		fmt.Print(string(newchar))
		newchar = 0
	}

}

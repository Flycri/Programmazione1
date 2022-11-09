package main

import "fmt"

func main() {
	/*
		Scrivere un programma max_char.go che legge da standard input una sequenza di 5 caratteri ASCII
		(byte) e stampa il maggiore in ordine lessicografico (cioè con il codice ASCII più alto).
	*/

	var char, max rune
	const NUM = 5

	for i := 0; i < NUM+3; i++ {
		fmt.Scanf("%c", &char) //problema degli invio: chiedere a lab
		if char > max {
			max = char
			fmt.Println("max=", max)
		}
	}

	fmt.Println("valore massimo lessicografico: ", string(max))

}

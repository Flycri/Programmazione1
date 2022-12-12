package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	/*
		Scrivere un programma posizioni_parole.go che legge una sequenza di stringhe da standard input e produce su standard output, per ogni stringa,
		la lista delle posizioni in cui essa compare nella sequenza (partendo dalla posizione 0).
		Nota: per terminare l’input da tastiera, premere invio e la combinazione di tasti Ctrl D, che corrisponde a
		EOF (end of file) per lo standard input.
		In caso di dubbi su come gestire la fine dell’input nel programma, consultare la documentazione della funzione Scan,
		funzione che, oltre a salvare i valori letti, restituisce dei valori.
	*/

	word_map := make(map[string][]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	i := 0

	for scanner.Scan() {
		word_map[scanner.Text()] = append(word_map[scanner.Text()], i)
		i++
	}

	fmt.Println(word_map)

}

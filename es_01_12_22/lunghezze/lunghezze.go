package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func aggiornaConteggio(m map[int]int, riga string) {
	var word_slice []string
	word_slice = strings.Split(riga, " ")
	for _, v := range word_slice {
		m[len(v)]++
	}
}

func main() {

	/*
		Scrivere un programma lunghezze.go che legge riga per riga un testo da standard input (potete usare la ridirezione),
		terminato da EOF, e stampa quante parole ci sono di lunghezza 1, quante di lunghezza 2, ecc.
		Il programma è dotato di una funzione

		func aggiornaConteggio(m map[int]int, riga string)


		che aggiorna la mappa dei conteggi e che deve essere usata dal main.
	*/

	scanner := bufio.NewScanner(os.Stdin)

	count_map := make(map[int]int)

	for scanner.Scan() {
		aggiornaConteggio(count_map, scanner.Text())
	}

	fmt.Println(count_map)

}

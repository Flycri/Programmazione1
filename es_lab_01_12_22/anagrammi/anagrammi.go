package main

import (
	"fmt"
	"os"
	"strings"
)

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	rune_map := make(map[rune]int)
	for _, r := range s1 {
		rune_map[r]++
	}
	for _, r := range s2 {
		x := strings.Count(s2, string(r))
		if rune_map[r] != x {
			return false
		}
	}
	return true
}

func main() {

	/*
		Scrivere un programma anagrammi.go che legge due stringhe da linea di comando e valuta se le due stringhe sono una l'anagramma
		dell'altra (la seconda stringa è formata da una permutazione dei caratteri della prima)
		In particolare il programma è dotato di:


		una funzione

		func isAnagram(s1, s2 string) bool

		che restituisce true se le due stringhe sono una l'anagramma dell'altra, false altrimenti


		una funzione

		func main()

		che legge due parole p1 e p2 da linea di comando e stampa uno dei due messaggi:
		"p1 e p2 sono anagrammi" o "p1 e p2 non sono anagrammi"

		NB: che caratteristiche ha un anagramma?
		C'è in Go una struttura di dati (array, struct, slice, map) che
		si presta a rappresentare i dati di una stringa s1 che servono
		ad individuare se è un anagramma di un'altra stringa s2?
		Se sì, quale e coma la uso? Se no, come imposto la soluzione?
		In assenza di esattamente due parametri sulla linea di comando il programma stampa: input errato
		nomefile: anagrammi.go
	*/

	if len(os.Args) != 3 {
		fmt.Println("input errato")
		os.Exit(1)
	}

	p1 := os.Args[1]
	p2 := os.Args[2]

	if isAnagram(p1, p2) {
		fmt.Printf("%s e %s sono anagrammi", p1, p2)
	} else {
		fmt.Printf("%s e %s non sono anagrammi", p1, p2)
	}

}

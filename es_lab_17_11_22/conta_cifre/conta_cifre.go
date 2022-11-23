package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func contaCifre(s string, numCifre *[10]int) (haCifre bool) {
	var digit int
	for _, r := range s {
		if unicode.IsDigit(r) {
			digit, _ = strconv.Atoi(string(r))
			(*numCifre)[digit]++
			haCifre = true
		}
	}
	return
}

func printArray(array []int) {
	for k, val := range array {
		fmt.Printf("numero %d: %d\n", k, val)
	}
}

func main() {

	/*
		Scrivete un programma conta_cifre.go dotato di:

		- una funzione `main` che legge una sequenza di stringhe da standard input fino a incontrare la parola "stop", analizza una stringa alla volta utilizzando la funzione `contaCifre` (vedi sotto) e alla fine stampa:
			- quante stringhe contengono almeno una cifra
			- per ogni cifra (0, 1, ..., 9), il numero di volte che quella cifra appare nella sequenza di stringhe
		- una funzione
			contaCifre(s string, numCifre *[10]int) (haCifre bool)
		che, data una stringa, aggiorna il conteggio delle cifre incontrate e restituisce true se la stringa contiene almeno una cifra, false altrimenti.

		Nota. Le stringhe possono contenere caratteri di qualsiasi tipo (cifre, lettere, simboli, punteggiatura, lettere accentate, ecc.).

		Il programma NON deve fare uso di variabili globali.

		Domanda: che prototipo (signature) deve avere la funzione `contaCifre`?
		- ha parametri? se sì, quanti, di che tipi e con che finalità?
		- restituisce valori? se sì, quanti, di che tipi e con che finalità?
		Soffermatevi su questi punti per progettare la funzione prima di scriverla.

		Esempio di esecuzione:

		$ go run conta_cifre.go
		c140 c140
		c0m3 t1 ch14m1?
		bye bye
		stop
		5 stringhe con cifre
		[0 1 2 3 4 5 6 7 8 9]
		[3 5 0 1 3 0 0 0 0 0]

	*/
	var input string
	var digitArray [10]int
	var numDigitStrings int
	var flag bool
	stdArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Scan(&input)
	for input != "stop" {
		flag = contaCifre(input, &digitArray)
		if flag {
			numDigitStrings++
		}
		fmt.Scan(&input)
	}
	fmt.Printf("%d stringhe con cifre\n", numDigitStrings)
	fmt.Println(stdArray)
	fmt.Println(digitArray)
	//printArray(digitArray[:])
	//fmt.Printf("numero di stringhe con cifre: %d ", numDigitStrings)

}

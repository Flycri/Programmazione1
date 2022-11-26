package main

import (
	"bufio"
	"fmt"
	"os"
)

const LEN_ALFABETO = 26

func aggiornaConteggi(s string, contaMinu *[LEN_ALFABETO]int) {
	if (contaMinu) == nil {
		return
	}
	for _, r := range s {
		if (r >= 'a') && (r <= 'z') {
			contaMinu[int(r-'a')]++
		}
	}
}

func printArray(array []int) {
	for i := 0; i < len(array); i++ {
		fmt.Println(string('a'+rune(i)), array[i])
	}
}

func main() {

	/*
		Scrivere un programma contaLettere.go che legge un testo (da stdin) e stampa quante volte (anche 0) appare nel testo
		ciascuna lettera minuscola dell'alfabeto ('a'-'z').
		Il programma è dotato di una costante
		const LEN_ALFABETO = 26
		e di una funzione
		func aggiornaConteggi(s string, contaMinu *[LEN_ALFABETO]int)
		che, data una stringa s, aggiorna i conteggi delle lettere minuscole di s.
		Si consiglia di usare la ridirezione dell'input per provare il programma.
	*/

	var alphabet [LEN_ALFABETO]int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		aggiornaConteggi(scanner.Text(), &alphabet)
	}

	printArray(alphabet[:])
}

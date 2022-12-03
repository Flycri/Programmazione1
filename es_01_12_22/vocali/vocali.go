package main

import (
	"bufio"
	"fmt"
	"os"
)

func contaVocali(s string, vocali map[rune]int) {
	for _, r := range s {
		if (r == 'a') || (r == 'A') || (r == 'e') || (r == 'E') || (r == 'i') || (r == 'I') || (r == 'o') || (r == 'O') || (r == 'u') || (r == 'U') {
			vocali[r]++
		}
	}
}

func main() {

	/*
		Scrivere un programma vocali.go che analizza un testo e conta le occorrenze delle vocali (sia minuscole che maiuscole,
		ma non le accentate) nel testo e stampa, ma solo per le vocali presenti nel testo, il numero di volte che le vocali stesse
		sono presenti nel testo.
		In particolare il programma è dotato di:

		una funzione

		func contaVocali(s string, vocali map[rune]int)

		per contare le occorrenze delle diverse vocali (sia minuscole che maiuscole - vedi es sotto) in tutte le stringhe che
		le vengono passate. La funzione, data una stringa s e una mappa vocali, aggiorna opportunamente la mappa vocali
		aggiungendo eventuali vocali non ancora presenti / incrementandone i valori.
		Nota: per individuare le vocali e aggiornare la variabile vocali usate uno switch con un solo case o un if,
		sempre con un solo caso.

		una funzione

		func main()

		che legge una riga di testo da standard input e produce una mappa tra vocali presenti nel testo e il numero delle loro
		occorrenze nel testo, e la stampa.
	*/

	vow_map := make(map[rune]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		contaVocali(scanner.Text(), vow_map)
	}

	for k, v := range vow_map {
		fmt.Printf("%s %d\n", string(k), v)
	}

}

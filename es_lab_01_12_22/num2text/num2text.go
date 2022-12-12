package main

import (
	"fmt"
	"os"
)

func printNumber(num string, num_map map[rune]string) {
	for i := 0; i < len(num)-1; i++ {
		fmt.Printf("%s - ", num_map[rune(num[i])])
	}
	fmt.Printf("%s", num_map[rune(num[len(num)-1])])
}

func main() {

	/*
		Scrivere un programma num2text.go per convertire un numero intero non negativo nella sequenza delle parole corrispondenti alle sue cifre.
		Il programma legge un intero non negativo da standard input, per ogni nuova (non incontrata finora) cifra del numero chiede
		il nome corrispondente (e alimenta un dizionario), e infine stampa la sequenza delle parole corrispondenti alle sue cifre.
	*/

	num_map := make(map[rune]string)

	for _, r := range os.Args[1] {
		if num_map[r] == "" {
			var s string
			fmt.Println("parola per ", string(r), "?")
			fmt.Scan(&s)
			num_map[r] = s
		}
	}

	printNumber(os.Args[1], num_map)

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
	Scrivere un programma max.go, dotato di:

	una funzione ricorsiva recursiveMax(list []int) int che restituisca il massimo tra i valori di list
	una funzione main() che legga da standard input (ctrl D per terminare) una lista di numeri interi
	(che posono essere positivi, negativi, nulli) ed emetta nel flusso d'uscita il massimo tra i numeri letti.

	Il massimo di una sequenza vuota non è definito, quindi assumiamo (non è richiesto che vengano fatti controlli)
	che la sequenza abbia sempre almeno un numero.
	Può essere comodo definire una funzione di supporto


	func greater(m, n int) int
	che restituisce il maggiore tra m e n.
*/

func greater(m, n int) int {
	if m > n {
		return m
	} else {
		return n
	}
}

func recursiveMax(list []int) int {
	if len(list) == 1 {
		return list[0]
	} else {
		return greater(list[0], recursiveMax(list[1:]))
	}
}

func main() {

	var int_list []int

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		int_list = append(int_list, x)
	}

	fmt.Println(recursiveMax(int_list))

}

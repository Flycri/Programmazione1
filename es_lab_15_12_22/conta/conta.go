package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
	Scrivere un programma conta.go, dotato di:

	una funzione ricorsiva recursiveCount(el int, list []int) int che restituisca il numero di volte che el appare tra i valori di list
	una funzione main() che legga da standard input (ctrl D per terminare) una lista di numeri interi (che posono essere positivi,
	negativi, nulli) ed emetta nel flusso d'uscita quante volte il primo numero letto è ripetuto nella sequenza dei successivi numeri
	letti.
*/

func recursiveCount(el int, list []int) int {
	if len(list) == 1 {
		if el == list[0] {
			return 1
		} else {
			return 0
		}
	} else {
		if el == list[0] {
			return 1 + recursiveCount(el, list[1:])
		} else {
			return 0 + recursiveCount(el, list[1:])
		}
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

	fmt.Println(recursiveCount(int_list[0], int_list))

}

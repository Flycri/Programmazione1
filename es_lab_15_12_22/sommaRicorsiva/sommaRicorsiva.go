package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
	Scrivere un programma sommaRicorsiva.go, dotato di:

	una funzione ricorsiva func recursiveSum(list []int) int
	che calcola la somma degli interi della slice list passata come parametro,
	una funzione main() che legga da standard input (ctrl D per terminare) una lista di numeri interi
	ed emetta nel flusso d'uscita la somma dei numeri letti.
*/

func recursiveSum(list []int) int {
	if len(list) == 0 {
		return 0
	} else {
		return list[0] + recursiveSum(list[1:])
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

	fmt.Println(recursiveSum(int_list))

}

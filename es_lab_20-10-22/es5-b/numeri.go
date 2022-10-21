package main

import "fmt"

func main() {
	/*
		Scrivere un programma numeri.go che riceve in ingresso un numero intero e stampa "divisibile per 10"
		se il numero è divisibile per 10, "positivo" se il numero è positivo, "positivo divisibile per 10"
		se è sia divisibile per 10 che positivo, niente altrimenti.
	*/

	var x int

	fmt.Scan(&x)

	if x%10 == 0 && x > 0 {
		fmt.Println("positivo divisibile per 10")
	} else if x > 0 {
		fmt.Println("positivo")
	} else if x%10 == 0 {
		fmt.Println("divisibile per 10")
	}

}

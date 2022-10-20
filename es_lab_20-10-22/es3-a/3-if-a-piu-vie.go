package main

import "fmt"

func main() {
	/*
	 Scrivere un programma che prende in input un intero positivo o negativo e stampa a schermo
	 se è positivo, nullo o negativo
	*/

	var num int

	fmt.Print("un int: ")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("positivo")
	} else if num == 0 {
		fmt.Println("nullo")
	} else { // < 0
		fmt.Println("negativo")
	}
}

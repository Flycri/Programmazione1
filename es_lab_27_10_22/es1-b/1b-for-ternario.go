package main

import "fmt"

func main() {
    /*
    Scrivere un programma che per 5 volte chieda in input un intero e calcoli la somma di
	tutti i numeri inseriti
    */
    const K = 5
	var n int
	sum := 0
	for i := 1; i <= K; i++ {
		fmt.Print("un numero: ")
		fmt.Scan(&n)
		sum = sum + n
	}
	fmt.Println(sum)
}

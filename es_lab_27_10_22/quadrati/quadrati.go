package main

import "fmt"

func main() {
	/*
	Scrivere un programma quadrati.go che legge da standard input un intero n positivo e calcola,
	utilizzando solo variabili di tipo int, il massimo quadrato (k^2) <= n.
	Per calcolare il quadrato di un numero n usate n*n.
	*/

	var n,k int

	fmt.Scan(&n)

	for k=1;k*k<n;k++ {
		fmt.Println(k*k)
	}
	
	k--

	fmt.Println(k*k)

}
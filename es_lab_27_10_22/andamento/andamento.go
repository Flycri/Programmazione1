package main

import "fmt"

func main() {
	/*
		Scrivere un programma andamento.go che legge da tastiera una serie (di almeno un elemento) di numeri interi > -1
		e stampa "+" ogni volta che il nuovo valore è maggiore o uguale al precedente e "-" altrimenti.
		Si ferma quando il numero in input è -1 e stampa la somma di tutti i numeri letti.
	*/

	var x,oldx,sum int

	fmt.Scan(&x)

	for x!=-1 {
		if x>=oldx {
			fmt.Print("+")
		} else {
			fmt.Print("-")
		}
		sum+=x
		oldx=x
		fmt.Scan(&x)
	}

	fmt.Println(sum)

}
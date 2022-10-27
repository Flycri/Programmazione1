package main

import "fmt"

func main() {
	/*
		Scrivere un programma es05.go che legge un numero n e stampa i numeri da 1 a n.
	*/

	var n int 

	fmt.Scan(&n)

	for i:=1;i<=n;i++ {
		fmt.Println(i)
	}

}
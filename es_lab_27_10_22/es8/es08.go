package main

import "fmt"

func main() {
	/*
		Scrivere un programma es08.go che legge un numero n e stampa
		la tabellina di n (cioè n*1, n*2, ..., n*10)
	*/

	var n int

	fmt.Scan(&n)

	for i:=1;i<=10;i++ {
		fmt.Println(n*i)
	}

}
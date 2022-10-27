package main

import "fmt"

func main() {
	/*
		Scrivere un programma es02.go che legge K = 5 numeri e di ciascuno stampa il doppio.
	*/
	const K = 5
	var n int
	for i := 1; i <= K; i++ {
		fmt.Scan(&n)
		fmt.Println("il doppio di",n,"è",n*2)
	}

}
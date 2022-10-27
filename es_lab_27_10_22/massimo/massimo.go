package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		Scrivere un programma massimo.go che genera 10 numeri interi casuali tra 10 e 30,
		li stampa, e stampa il massimo valore generato.
	*/

	const K=30
	var x,max int

	rand.Seed(time.Now().UnixNano())

	for i:=0;i<10;i++ {
		x = rand.Intn(K+1)
		fmt.Print(x," ")
		if x>max {
			max=x
		}
	}

	fmt.Println(max)

}
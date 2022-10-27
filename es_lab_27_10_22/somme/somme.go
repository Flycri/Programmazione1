package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		Scrivere un programma somme.go che genera 10 numeri interi casuali tra 0 e 10,
		li stampa, e stampa la somma dei valori generati.
	*/

	const K=10
	var x,sum int

	rand.Seed(time.Now().UnixNano())

	for i:=0;i<K;i++ {
		x = rand.Intn(K+1)
		fmt.Print(x," ")
		sum+=x
	}

	fmt.Println(sum)

}
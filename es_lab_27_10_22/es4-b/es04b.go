package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		Scrivere un programma es04b.go che genera K = 10
		numeri casuali in [0,10], conta quanti sono pari e stampa questo risultato.
	*/
	
	const K=10
	var x,sum int

	rand.Seed(time.Now().UnixNano())

	for i:=0;i<K;i++ {
		x = rand.Intn(K+1)
		fmt.Print(x," ")
		if (x%2==0) {
			sum++
		}
	}

	fmt.Println("\n ci sono",sum,"numeri pari")


}
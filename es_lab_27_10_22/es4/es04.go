package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		Scrivere un programma es04.go che genera K = 10
		numeri casuali in [0,10] e li stampa su una riga, separati da spazi. rand.Intn(max)
	*/
	
	const K=10

	rand.Seed(time.Now().UnixNano())

	for i:=0;i<K;i++ {
		fmt.Print(rand.Intn(K)," ")
	}

}
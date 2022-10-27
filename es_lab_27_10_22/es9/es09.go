package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		Scrivere un programma es09.go che legge un numero n e stampa i numeri tra 1 e n che
		sono dei quadrati.
	*/

	var n int 
	var root float64

	fmt.Scan(&n)

	for i:=1;i<=n;i++ {
		root = math.Sqrt(float64(i))
		if (root/float64(int(root)))==1 {
			fmt.Println(i)
		}
	}

}
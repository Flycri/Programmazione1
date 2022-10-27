package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)
func main() {
	/*
		Scrivere un programma vicino.go che, data una serie di 5 valori interi tra 0 e 20,
		trova il valore piu' vicino a TARGET, dove TARGET è una costante definita nel programma.
	*/

	const TARGET,MAX = 10,20
	var x, minDiff, minDiffTemp int 
	var isNeg bool

	rand.Seed(time.Now().UnixNano())

	minDiff = MAX
	for i:=0;i<5;i++ {
		x = rand.Intn(MAX+1)
		fmt.Print(x," ")
		if (TARGET-x)<0 {
			isNeg = true
		} else {
			isNeg = false
		}
		minDiffTemp = int(math.Abs(float64(TARGET-x)))
		if (minDiffTemp<minDiff) {
			minDiff = minDiffTemp
		}
	}

	if isNeg {
		fmt.Println("il valore più vicino è",TARGET + minDiff)
	} else {
		fmt.Println("il valore più vicino è",TARGET - minDiff)
	}
}
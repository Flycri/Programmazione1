package main

import (
	"fmt"
	"math"
)

func main() {
	/*
	Scrivere un programma intorno.go che legge due float64 x e y e stampa "vicino all'origine"
	se il punto di coordinate (x,y) è abbastanza vicino (a meno di 10^-5) all'origine (0,0),
	mentre stampa "non vicino all'origine" in caso contrario.
	Impostate una costante usando la notazione scientifica di Go, ad esempio:  const  EPSILON = 1e-5
	*/
	var x,y float64
	const EPSILON=1e-5

	fmt.Scan(&x)
	fmt.Scan(&y)
	if (math.Sqrt((math.Pow(x,2)+math.Pow(y,2)))<EPSILON) {
		fmt.Println("vicino all'origine")
	} else {
		fmt.Println("non vicino all'origine")
	}
}
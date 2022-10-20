package main

import (
	"fmt"
	"math"
)

func main() {
	/*
	Scrivere un programma punto_retta.go che legge da standard input le cooordinate (x1,y1) di un punto P
	e i coefficienti m e q di una retta y = mx + q; controlla se il punto P sta sopra, sotto o appartiene
	alla retta y data, e stampa rispettivamente "sopra", "sotto", "sulla retta".
	Si dichiarino i dati in ingresso come float64.
	Nota Bene, trattando con numeri float, cosideriamo appartenenti alla retta
	i punti che distano dalla retta meno di 10^-6. In Go 10^-6 può essere scritto come 1e-06.
	*/
	const EPSILON=1e-6
	var x,y,m,q,d float64

	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scan(&m)
	fmt.Scan(&q)

	d = (y-(m*x+q))/(math.Sqrt(1+math.Pow(m,2)))
	if (math.Abs(d)<=EPSILON) {
		fmt.Println("sulla retta")
	} else {
		if (d>0) {
			fmt.Println("sopra")
		}
		if (d<0) {
			fmt.Println("sotto")
		}
	}

}

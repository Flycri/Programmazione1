package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		Scrivere un programma rettangolo.go che legge da tastiera le coordinate (int) x,y di tre punti P1, P2, P3 e stabilisce se P3 è interno,
		perimetrale o esterno al rettangolo che ha P1 e P2 come vertici su una diagonale (vedi figura).
	*/

	var x1, x2, x3, y1, y2, y3 int

	fmt.Scan(&x1)
	fmt.Scan(&y1)
	fmt.Scan(&x2)
	fmt.Scan(&y2)
	fmt.Scan(&x3)
	fmt.Scan(&y3)

	if ((x3 >= int(math.Min(float64(x1), float64(x2)))) && (x3 <= int(math.Max(float64(x1), float64(x2))))) && ((y3 >= int(math.Min(float64(y1), float64(y2)))) && (x3 <= int(math.Max(float64(y1), float64(y2))))) {
		fmt.Println("interno")
	} else {
		fmt.Println("esterno")
	}
}

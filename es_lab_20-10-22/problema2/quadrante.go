package main

import "fmt"

func main() {
	/*
		Scrivere un programma quadrante.go che legge da tastiera le coordinate x,y (float64)
		di un punto P nel piano cartesiano e stampa in che quadrante (I, II, III, IV) si trova P.
		Per i punti che cadono sull’asse delle ordinate e/o delle ascisse, assegnate il quadrante positivo (+).
	*/

	var x, y float64

	fmt.Scan(&x)
	fmt.Scan(&y)

	if x >= 0 {
		if y >= 0 {
			fmt.Println("I quadrante")
		} else {
			fmt.Println("IV quadrante")
		}
	} else {
		if y >= 0 {
			fmt.Println("II quadrante")
		} else {
			fmt.Println("III quadrante")
		}
	}
}

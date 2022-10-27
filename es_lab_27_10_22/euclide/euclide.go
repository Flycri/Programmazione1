package main

import "fmt"

func main() {
	/*
		
	Scrivere un programma euclide.go che legge da standard input due interi a e b, con a >= b,
	e calcola il MCD tra i due numeri con l'algoritmo di Euclide.
	Algoritmo di Euclide:
	Dati due numeri naturali a e b:

		controlla se b è zero.
		se lo è, a è il MCD.
		se non lo è, assegna ad r il resto della divisione a / b, poni a = b e b = r e ripeti da 1.
	*/

	var a,b,r int

	for {
		fmt.Println("numero 1 (maggiore)")
		fmt.Scan(&a)
		fmt.Println("numero 2 (minore)")
		fmt.Scan(&b)
		if (a>=b) {
			break;
		}
	}

	for b!=0 {
		r = a%b
		a = b
		b = r
	}

	fmt.Println(a,"è il MCD")

}
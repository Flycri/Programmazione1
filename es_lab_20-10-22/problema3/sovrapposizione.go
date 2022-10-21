package main

import "fmt"

func main() {
	/*
		Scrivere un programma sovrapposizione.go che legge da tastiera il giorno [1-31], l'ora di inizio [0-24]
		e l'ora di fine [0-24] di due appuntamenti e stabilisce se si sovrappongono (anche parzialmente) o no.
	*/

	var g1, g2, sh1, sh2, eh1, eh2 int

	fmt.Scan(&g1)
	fmt.Scan(&sh1)
	fmt.Scan(&eh1)
	fmt.Scan(&g2)
	fmt.Scan(&sh2)
	fmt.Scan(&eh2)

	if (g1 == g2) && (((sh2 >= sh1) && (sh2 <= eh1)) || ((eh2 >= sh1) && (eh2 <= eh1))) {
		fmt.Println("si sovrappongono")
	} else {
		fmt.Println("non si sovrappongono")
	}
}

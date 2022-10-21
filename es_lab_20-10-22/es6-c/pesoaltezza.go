package main

import "fmt"

func main() {
	/*
		Scrivere un programma pesoaltezza.go che riceve in ingresso due numeri (float64)
		rappresentanti peso in kg e altezza in cm di una persona e stampa un responso in funzione della tabella:(vedere doc)
	*/

	var weight, height int

	fmt.Scan(&weight)
	fmt.Scan(&height)

	if (((weight >= 10) && (weight <= 50)) && ((height >= 100) && (height <= 150))) || (((weight >= 50) && (weight <= 100)) && ((height > 150) && (height <= 200))) {
		fmt.Println("normopeso")
	} else if (((weight > 50) && (weight <= 100)) && ((height >= 100) && (height <= 150))) || ((weight > 100) && ((height >= 150) && (height <= 200))) {
		fmt.Println("sovrappeso")
	} else if ((weight >= 10) && (weight <= 50)) && ((height > 150) && (height <= 200)) {
		fmt.Println("sottopeso")
	} else if (weight > 100) && ((height >= 100) && (height < 150)) {
		fmt.Println("molto sovrappeso")
	}

}

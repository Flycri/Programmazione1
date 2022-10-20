package main

import "fmt"

func main() {
	/*
	 Scrivere un programma che preso in input un valore intero rappresentante un voto, dica se un voto è valido (compreso 
	fra 0 e 30 compresi).
	*/

	var voto int

	fmt.Print("voto: ")
	fmt.Scan(&voto)

	if voto < 0 || voto > 30 {
		fmt.Println("voto non valido")
	}
}

package main

import "fmt"

func main() {
	/* 	 
	Scrivere un programma che dato un intero controlli che sia compreso fra 0 e 100 (interrompendo il programma
	in caso negativo) e poi stampi a schermo:
	A se è [90,100]
	B se è [80,90)
	C se è [70,80)
	D se è [60,70)
	F se è [0,60)
	*/

	var voto int
	fmt.Print("voto? ")
	fmt.Scan(&voto)

	if voto < 0 || voto > 100 {
		fmt.Println("voto non valido")
		return //interrompe l'esecuzione del programma
	}

	if voto >= 90 {
		fmt.Println("A")
	} else if voto >= 80 {
		fmt.Println("B")
	} else if voto >= 70 {
		fmt.Println("C")
	} else if voto >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}

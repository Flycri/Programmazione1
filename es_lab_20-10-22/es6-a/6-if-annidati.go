package main

import "fmt"

func main() {
	/*
		Scrivere un programma che stampi a schermo la formula per calcolare l'area od il perimetro di un triangolo o di un rettangolo,
		facendo inserire in input all'utente due interi, il primo rappresentante la forma (rettangolo=1, triangolo=2) e
		il secondo il tipo di formula (perimetro=1, area=2).

	*/
	const (
		RETTANGOLO = 1
		TRIANGOLO  = 2
		PERIMETRO  = 1
		AREA       = 2
	)
	var figura, operazione int

	fmt.Print("rettangolo (1) o triangolo (2)? ")
	fmt.Scan(&figura)
	fmt.Print("perimetro (1) o area (2)? ")
	fmt.Scan(&operazione)

	if figura == RETTANGOLO {
		if operazione == PERIMETRO {
			fmt.Println("formula = (base + altezza) * 2")
		} else { //operazione == AREA
			fmt.Println("formula = base * altezza")
		}
	} else { //figura == TRIANGOLO
		if operazione == PERIMETRO {
			fmt.Println("formula = lato1 + lato2 + lato3 ")
		} else { //operazione == AREA
			fmt.Println("formula = (base * altezza)/2")
		}
	}
}

package main

import "fmt"

func main() {
	/*
		Definiamo "gradino" una sequenza di (uno o più) interi non negativi uguali seguita da
		un'altra sequenza di (uno o più) interi più grandi di 1 (es. 1 1 2 2 2).
		Scrivere un programma gradino.go che, data in input una sequenza di interi tali che ogni
		intero è >= del precedente, stampa la lunghezza (il numero di interi) del gradino più lungo
		(Si noti che i gradini si sovrappongono). Il programma termina quando legge un numero minore
		di quello appena letto.
	*/

	var x, curr, prec, precNum, currNum, maxTemp int

	fmt.Scan(&x)
	prec = x
	currNum = 1

	for {
		fmt.Scan(&x)
		curr = x
		if curr < prec {
			break
		}
		if prec == curr {
			currNum++
		} else {
			precNum = currNum
			currNum = 1
		}
		if (precNum + currNum) > maxTemp {
			maxTemp = precNum + currNum
		}
		prec = curr
	}

	fmt.Println(maxTemp)

}

package main

import "fmt"

func main() {
	/*
		Scrivere un programma num_sequenze.go che legge da standard input una sequenza di uni (1)
		e zeri (0) (terminata da un 2), che inizia e finisce con 1, e stampa il numero di
		sottosequenze di zeri.
	*/

	var x, prec, curr, sum0 int
	var err bool

	fmt.Scan(&x)
	if x == 1 {
		prec = x
		for x != 2 {
			fmt.Scan(&x)
			if x != 0 && x != 1 && x != 2 {
				err = true
				break
			}
			curr = x
			if (curr != 0) && prec == 0 {
				sum0++
			}
			prec = curr
		}
		if err {
			fmt.Println("cifra inserita non 0,1 o 2!")
		} else {
			fmt.Println(sum0)
		}

	} else {
		fmt.Println("Il primo numero della sequenza non è 1!")
	}
}

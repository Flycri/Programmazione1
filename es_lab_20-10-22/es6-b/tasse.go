package main

import "fmt"

func main() {
	/*
		Scrivere un programma tasse.go che chiede reddito (int) e se coniugato (bool)
		e stampa le tasse da pagare secondo la seguente tabella.
		Usate costanti per le aliquote (10% e 25%) e per gli scaglioni (32000 e 64000).
	*/

	const ALIQ1, ALIQ2, SCA1, SCA2 = 0.1, 0.24, 32000, 64000

	var reddito int
	var isConiugato bool

	fmt.Scan(&reddito)
	fmt.Scan(&isConiugato)

	if isConiugato {
		if reddito >= 0 && reddito <= SCA2 {
			fmt.Println(float64(reddito) * ALIQ1)
		} else {
			fmt.Println(float64(reddito) * ALIQ2)
		}
	} else {
		if reddito >= 0 && reddito <= SCA1 {
			fmt.Println(float64(reddito) * ALIQ1)
		} else {
			fmt.Println(float64(reddito) * ALIQ2)
		}
	}
}

package main

import (
	"fmt"
	"math"
)
func main() {
	/*
		Scrivere un programma bin2ten.go che converte un intero composto di soli 0 e 1 nel valore
		corrispondente al numero binario rappresentato (es. 101 -> 5).
		Nel caso il numero contenga altre cifre, il programma stampa un messaggio di errore.
	*/

	var x,i,j,sum int
	var e bool

	fmt.Scan(&x)



	for i=1;x/i>0;i*=10 {
		if (x/i%10!=1)&&(x/i%10!=0) {
			e = true
			break;
		}
		sum+=x/i%10 * int(math.Pow(2,float64(j)))
		j++
	}

	if e {
	fmt.Println("errore: numero non contiene solo 0 e 1")
	} else {
		fmt.Println(x,"è",sum,"in decimale")
	}

}
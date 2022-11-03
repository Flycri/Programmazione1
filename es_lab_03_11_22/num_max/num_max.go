package main

import "fmt"

func main() {
	/*
		Scrivere un programma num_max.go che legge una sequenza di 10 interi positivi 
		e stampa il massimo intero letto e quante volte tale massimo compare nella sequenza.
	*/

	var x,maxtemp,nmax int
	const MAX = 10

	for i:=0;i<MAX;i++ {
		fmt.Scan(&x)
		if (x>maxtemp) {
			nmax=1
			maxtemp = x
		} else if (x==maxtemp) {
			nmax++
		}
	}

	fmt.Println("il massimo è",maxtemp,"e compare",nmax,"volte")
}
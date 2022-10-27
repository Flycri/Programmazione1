package main

import "fmt"

func main() {
	/*
	Scrivere un programma differenze.go che legge una serie di valori float64 da tastiera e stampa
	le differenze, cioè la differenza tra il secondo e il primo, tra il terzo e il secondo, e così via.
	Il programma termina quando riceve in input il valore 0.
	*/

	var x,xtemp float64

	fmt.Scan(&x)

	for x!=0 {
		xtemp = x
		fmt.Scan(&x)
		if (x==0) {
			break;
		}
		fmt.Println("differenza:",x-xtemp)
	}

}
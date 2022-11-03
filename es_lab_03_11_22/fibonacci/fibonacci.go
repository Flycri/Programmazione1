package main

import "fmt"

func main() {
	/*
		Scrivere un programma fibonacci.go che legge un intero positivo n e stampa i numeri di fibonacci 
		dal primo all'n-esimo, rappresentandoli come righe di asterischi, ciascuna lunga quanto 
		il numero da rappresentare.
	*/
	
	var n,x1,x2,xtemp int

	fmt.Scan(&n)

	x1=1
	x2=1

	fmt.Println("*")
	for i:=0;i<n-1;i++ {
		for j:=0;j<x2;j++ {
			fmt.Print("*")
		}
		xtemp=x2
		x2=x2+x1
		x1=xtemp
		fmt.Print("\n")
	}

}
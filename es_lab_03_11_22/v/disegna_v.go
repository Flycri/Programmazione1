package main

import "fmt"

func main() {
	/*
		Scrivere un programma disegna_v.go che legge un intero 
		positivo n e stampa una v di altezza n di asterischi.
	*/
	var n int
	var s string

	s="*"

	fmt.Scan(&n)

	for i:=n;i>0;i-- {
		fmt.Print(string(s))
		for j:=0;j<(i*2)-1;j++{
			fmt.Print(" ")
		}
		fmt.Print("*")
		s=" " + s
		fmt.Print("\n")
	}
	fmt.Print(s+"\n")
}
package main

import "fmt"

func main() {
	/*
		Scrivere un programma disegna_slash.go che legge un intero positivo n 
		e stampa un backslash (\) di altezza n composto da asterischi.
	*/

	var n int
	var s string

	s="*"

	fmt.Scan(&n)

	for i:=0;i<n;i++ {
		fmt.Println(string(s))
		s=" " + s
	}

}
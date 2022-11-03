package main

import "fmt"

func main() {
	/*
		Scrivere un programma crescente.go che legge da standard input una stringa di 
		sole lettere minuscole e la stampa inserendo un '-' ogni volta che una lettera viene 
		prima in ordine alfabetico della lettera precedente.
	*/

	var s string

	fmt.Scan(&s)

	for i:=0;i<len(s)-1;i++ {
		fmt.Print(string(s[i]))
		if (s[i]>s[i+1]) {
			fmt.Print("-")
		}
	}
	fmt.Print(string(s[len(s)-1]))
}
package main

import "fmt"

func main() {
	/*
	Scrivere un programma Go maggiore.go che legga due interi, li salvi in due var min e max nell'ordine
	in cui li legge; se non sono in ordine crescente, li sistemi in modo che
	min contenga il minore e max il maggiore; infine stampi il contenuto di max.
	*/
	var min, max int
	fmt.Scan(&min)
	fmt.Scan(&max)
	if (min>max) {
		min,max = max,min
	}
	fmt.Println(max)
}
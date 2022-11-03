package main

import (
	"fmt"
	"unicode"
)
func main() {
	/*
		Scrivere un programma minu_maiu.go che legge da standard input una stringa e stabilisce se la 
		stringa contiene solo minuscole, solo maiuscole o sia minuscole che maiuscole, 
		quindi stampa un messaggio opportuno (es. "solo minuscole","solo minuscole", "sia minuscole 
		che maiuscole").
	*/

	var s string
	var nlower,nupper int

	fmt.Scan(&s)

	for _,r:=range s {
		if unicode.IsLower(r) {
			nlower++
		}
		if unicode.IsUpper(r) {
			nupper++
		}
	}

	if nlower==len(s) {
		fmt.Println("solo minuscole")
	} else if nupper==len(s) {
		fmt.Println("solo maiuscole")
	} else {
		fmt.Println("sia minuscole che maiuscole")
	}
}
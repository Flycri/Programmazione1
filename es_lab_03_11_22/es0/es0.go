package main

import "fmt"

func main() {
	/*
		Scrivere un programma es0.go che fa le seguenti cose:

			legge un byte
			lo stampa
			stampa il byte precedente, il byte stesso, e quello successivo in ordine lessicografico (ASCII).
				Ad esempio, con 'd' come input, deve stampare: c d e

			stabilisce se è una lettera tra A e L, o altro (stampa "A-L" o "altro")
			poi legge una stringa e la stampa in verticale, una runa per riga.
	*/

	var char rune
	var s string

	fmt.Scanf("%c",&char)
	fmt.Println(string(char))
	fmt.Println(string(char-1),"\t",string(char),"\t",string(char+1))

	if (char>='A'&&char<='L') {
		fmt.Println("A-L")
	} else {
		fmt.Println("altro")
	}

	fmt.Scan(&s)

	for _,r:=range s {
		fmt.Println(string(r))
	}

}
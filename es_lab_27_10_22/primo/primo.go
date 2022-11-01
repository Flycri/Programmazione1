package main

import "fmt"

func main() {
	/*
		Scrivere un programma primo.go che, dato un numero intero su standard input, determina se il numero è primo.
	*/

	var x int
	var isPrimo bool

	isPrimo=true

	fmt.Scan(&x)

	for i:=2;i<=x/i;i++ {
		if (x%i==0) {
			isPrimo=false
			break;
		}
	}

	if (!isPrimo) {
		fmt.Println("il numero",x,"non è primo")
	} else {
		fmt.Println("il numero",x,"è primo")
	}

}
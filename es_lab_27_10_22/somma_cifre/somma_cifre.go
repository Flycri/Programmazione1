package main

import "fmt"

func main() {
	/*
		Scrivere un programma somma_cifre.go che calcola la somma delle cifre di un numero intero positivo fornito da standard input.
	*/

	var x,sum int

	fmt.Scan(&x)

	for i:=1;x/i>0;i*=10 {
		sum+=(x/i)%10
	}

	fmt.Println(sum)

}
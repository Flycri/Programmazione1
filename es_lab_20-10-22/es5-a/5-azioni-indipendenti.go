package main

import "fmt"

func main() {
	/*
		Scrivere un programma che preso in input un intero stampa a schermo Fizz se è divisibile per 3 e Buzz se è divisibile per 5
	*/

	var num int
	fmt.Print("numero? ")
	fmt.Scan(&num)

	if num%3 == 0 {
		fmt.Println("Fizz")
	}
	if num%5 == 0 {
		fmt.Println("Buzz")
	}
}

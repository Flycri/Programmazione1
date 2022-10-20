package main

import "fmt"

func main() {
	/*
	Scrivere un programma Go mattino.go che legga un orario (intero) e stampi
	"mattino" se l'orario è compreso tra le 6 (incluso) e le 13 (escluso).
	*/
	var ora int
	fmt.Scan(&ora)
	if (ora>=6)||(ora<=13) {
		fmt.Println("mattino")
	}
}
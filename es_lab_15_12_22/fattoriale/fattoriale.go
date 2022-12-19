package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
	Scrivere un programma fattoriale.go, dotato di:

	una funzione ricorsiva func fattoriale(n int) int
	che calcola il fattoriale del numero n passato come parametro,
	una funzione main() che, dato un numero naturale come
	argomento sulla linea di comando, emetta nel flusso d'uscita il fattoriale del numero dato.
*/

func fattoriale(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * fattoriale(n-1)
	}
}

func main() {

	x, _ := strconv.Atoi(os.Args[1])
	fmt.Println(fattoriale(x))

}

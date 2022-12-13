package main

import (
	"fmt"
	"os"
	"sort"
)

/*
	Scrivere un programma appello.go che legge da linea di comando una sequenza di nomi e li stampa in ordine alfabetico.
*/

func main() {

	sort.Strings(os.Args[1:])

	fmt.Println(os.Args[1:])

}

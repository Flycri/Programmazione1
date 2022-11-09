package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		Scrivere un programma ultima_pioggia.go che legge da standard input una sequenza di valori
		interi (terminata da EOF, da tastiera prodotto con ctrl-D) che indicano i mm di pioggia
		caduti (0 se non ha piovuto) ogni giorno in una sequenza successiva di giorni e stampa
		(l'indice del)l'ultimo giorno in cui ha piovuto.
	*/

	var totmm, i, lastDay int
	var err error

	i = 1
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		totmm, err = strconv.Atoi(line)
		if err != nil {
			break
		}
		if totmm != 0 {
			lastDay = i
		}
		i++
	}

	fmt.Println("L'ultimo giorno che ha piovuto era il giorno", lastDay)

}

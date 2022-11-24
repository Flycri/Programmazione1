package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	/*
	   Scrivere un programma ordini.go che legge da standard input una serie di stringhe che descrivono ordini nel
	   formato prezzo#quantità#sconto e stampa il totale finale da pagare.
	   Prezzo, quantità e sconto sono float; prezzo indica il prezzo unitario del prodotto, quantità indica la
	   quantità acquistata e sconto è lo sconto applicato per quel prodotto, espresso come float tra 0 e 1
	   (ad esempio 0.2 indica uno sconto del 20%). Il programma termina la lettura quando incontra EOF
	   (ctrl D su riga nuova).
	*/
	var tot float64

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		s := scanner.Text()
		temp_slice_s := strings.Split(s, "#")
		var temp_slice_f []float64
		for _, v := range temp_slice_s {
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			temp_slice_f = append(temp_slice_f, f)
		}
		tot += (temp_slice_f[0] * temp_slice_f[1] * (1 - temp_slice_f[2]))
	}
	fmt.Println(tot)

}

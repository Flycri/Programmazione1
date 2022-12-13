package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

/*
	Scrivere un programma temperature.go che legge delle temperature (int) da tastiera e termina quando il valore letto è 999.
	Il programma deve stampare:

	la media
	la mediana
	il numero di temperature sotto la media delle temperature stesse.
	le tre temperature più basse (se ci sono almeno 3 temperature)
	le tre temperature più alte (se ci sono almeno 3 temperature)

	Nota 1: la mediana di un insieme di dati  è data, nel caso ci sia un numero dispari di dati, dal dato centrale dei dati ordinati per valore (ad es. crescente), altrimenti dalla media dei due dati centrali.
	Nota 2: Utilizzare la funzione Ints(a []int) del pacchetto sort per ordinare in modo crescente le temperature.
*/

func main() {

	var temperature_slice []int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() == "999" {
			break
		}
		new_temp, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("errore nella conversione, dato scartato")
		} else {
			temperature_slice = append(temperature_slice, new_temp)
		}
	}

	sort.Ints(temperature_slice)

	var tot int
	for _, val := range temperature_slice {
		tot += val
	}
	media := float64(tot) / float64(len(temperature_slice))
	fmt.Println("la media è:", media)

	if len(temperature_slice)%2 != 0 {
		fmt.Println("la mediana è:", temperature_slice[len(temperature_slice)/2])
	} else {
		fmt.Println("la mediana è:", (float64(temperature_slice[len(temperature_slice)/2])+float64(temperature_slice[len(temperature_slice)/2-1]))/2)
	}

	for k, val := range temperature_slice {
		if val > int(media) {
			fmt.Printf("ci sono %d valori minori della media\n", k)
			break
		}
	}

	if len(temperature_slice) >= 3 {
		fmt.Println("3 temperature + basse", temperature_slice[:3])
		fmt.Println("3 temperature + alte", temperature_slice[len(temperature_slice)-3:])
	} else {
		fmt.Println("non ci sono almeno 3 valori")
	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Capoluogo struct {
	Nome       string
	Sigla      string
	Regione    string
	Superficie int
}

func isCapoluogo(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func main() {

	/*
		Scrivere un programma regioni.go che che legge da standard input (usate la ridirezione dal file capoluoghi.csv fornito, lo trovate qui) una serie di
		informazioni sui capoluoghi organizzate su righe, nel seguente formato:
		Nome,Sigla,Regione,Popolazione,Superficie,Densita,Altitudine
		e memorizza nome e regione in modo che sia possibile ottenere la stampa della lista dei capoluoghi delle regioni i cui
		nomi sono stati forniti da linea di comando.

	*/

	scanner := bufio.NewScanner(os.Stdin)
	first := true
	city_map := make(map[string]Capoluogo)

	for scanner.Scan() {
		if first {
			first = false
			continue
		}

		data_slice := strings.Split(scanner.Text(), ",")
		sup, _ := strconv.Atoi(data_slice[4])
		new_city := Capoluogo{Nome: data_slice[0], Sigla: data_slice[1], Regione: data_slice[2], Superficie: sup}

		if isCapoluogo(new_city.Nome) {
			city_map[new_city.Regione] = new_city
		}

	}

	for i := 1; i < len(os.Args); i++ {
		fmt.Println(city_map[os.Args[i]].Nome)
	}

}

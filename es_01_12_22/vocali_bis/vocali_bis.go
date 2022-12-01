package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func contaVocali(s string, vocali map[rune]int) {
	for _, r := range s {
		if (r == 'a') || (r == 'A') || (r == 'e') || (r == 'E') || (r == 'i') || (r == 'I') || (r == 'o') || (r == 'O') || (r == 'u') || (r == 'U') {
			vocali[r]++
		}
	}
}

func main() {

	/*
		Come si potrebbero stampare le vocali con le loro occorrenze in ordine alfabetico?
		Scrivere una seconda versione vocali_bis.go che produce l'output in ordine:
	*/

	var vow_slice []string

	vow_map := make(map[rune]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		contaVocali(scanner.Text(), vow_map)
	}

	for k, _ := range vow_map {
		vow_slice = append(vow_slice, string(k))
	}

	sort.Strings(vow_slice)

	for _, s := range vow_slice {
		for _, r := range s {
			fmt.Printf("%s %d\n", s, vow_map[r])
		}
	}

}

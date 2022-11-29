package main

import (
	"fmt"
	"os"
)

func anagramma(s string) map[string]bool {

	var res map[string]bool

	if len(s) == 0 {
		return map[string]bool{"": true}
	}

	res = make(map[string]bool)

	for i := 0; i < len(s); i++ { // es "cane"
		primaLettera := string(s[i]) //prima lettera dell'anagramma es "n"
		resto := s[:i] + s[i+1:]     //resto della parola es "ca" + "e"
		z := anagramma(resto)
		for k, _ := range z {
			res[primaLettera+k] = true
		}
	}

	return res

}

func main() {

	/*
		Data una stringa senza caratteri ripetuti, produrre una slice con tutti i suoi anagrammi
	*/

	s := os.Args[1]

	res := anagramma(s)

	for k, _ := range res {
		fmt.Printf("%s ", k)
	}

}

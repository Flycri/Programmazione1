package main

import (
	"fmt"
)

func main() {
	/*
		Leggere un testo (una sequenza di "parole", cioè di stringhe separate da spazi) da standard 
		input e stampare una parola per riga.
	*/

	var s, s_tot string

	s_tot = ""

	for {
		_,err := fmt.Scan(&s)
		if err!=nil {
			fmt.Println(err)
			break
		}
		s_tot+=s+"\n"
	}

	fmt.Println(s_tot)

}
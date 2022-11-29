package main

import (
	"fmt"
	"os"
	"strconv"
)

func f(n int) []string {

	var res []string

	if n == 0 {
		return []string{""} //caso base
	}

	s := f(n - 1)

	for _, x := range s {
		res = append(res, x+"0", x+"1")
	}

	return res

}

func main() {

	/*
		Funzione che dato n intero restituisce una slice di stringhe con tutte le possibili stringhe fatte da 0 e 1 di lunghezza n
	*/

	n, _ := strconv.Atoi(os.Args[1])

	fmt.Println(f(n))

}

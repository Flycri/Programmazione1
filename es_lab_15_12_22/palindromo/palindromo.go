package main

import (
	"fmt"
	"os"
)

/*
	Scrivere un programma palindromo.go, dotato di:

	una funzione ricorsiva isPalindrome(s string) bool che stabilisca se una stringa è palindroma
	una funzione main() che, data una stringa come argomento sulla linea di comando, emetta nel flusso
	d'uscita il messaggio "s è palindroma", se s
	è palindroma, e "s non è palindroma" altrimenti.
*/

func isPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	} else {
		if s[0] == s[len(s)-1] {
			return isPalindrome(s[1 : len(s)-1])
		} else {
			return false
		}
	}
}

func main() {

	if isPalindrome(os.Args[1]) {
		fmt.Println(os.Args[1], "è palindroma")
	} else {
		fmt.Println(os.Args[1], "non è palindroma")
	}

}

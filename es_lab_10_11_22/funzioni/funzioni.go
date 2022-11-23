package main

import (
	"fmt"
	"unicode"
)

/*
hasUpper(s string) bool
La funzione riceve una stringa come parametro e restituisce true se la stringa contiene
almeno una lettera latina maiuscola (tra 'A' e 'Z'), false altrimenti.
*/
func hasUpper(s string) bool {
	for i := 0; i < len(s); i++ {
		if unicode.IsUpper(rune(s[i])) {
			return true
		}
	}
	return false
}

/*
firstUpper(s string) int
La funzione riceve una stringa come parametro e restituisce la posizione della prima lettera
latina maiuscola (tra 'A' e 'Z'), se la stringa ne contiene almeno una, -1 altrimenti.
*/
func firstUpper(s string) int {
	for i := 0; i < len(s); i++ {
		if unicode.IsUpper(rune(s[i])) {
			return i
		}
	}
	return -1
}

/*
lastUpper(s string) int
La funzione riceve una stringa come parametro e restituisce la posizione dell'ultima lettera
latina maiuscola (tra 'A' e 'Z'), se la stringa ne contiene almeno una, -1 altrimenti.
*/
func lastUpper(s string) int {
	for i := len(s) - 1; i > 0; i-- {
		if unicode.IsUpper(rune(s[i])) {
			return i
		}
	}
	return -1
}

/*
countDigitsLettersOthers(s string) int int int
La funzione riceve una stringa come parametro, conta quante cifre, quante lettere e quanti altri
caratteri (né cifre né lettere) contiene e restituisce i tre risultati in questo ordine.
*/
func countDigitsLettersOthers(s string) (int, int, int) {
	var sumDigits, sumLetters, sumOthers int
	for i := 0; i < len(s); i++ {
		switch {
		case unicode.IsDigit(rune(s[i])):
			sumDigits++
		case unicode.IsLetter(rune(s[i])):
			sumDigits++
		default:
			sumOthers++
		}
	}
	return sumDigits, sumLetters, sumOthers
}

/*
isPalindrome(s string) bool
La funzione riceve una stringa come parametro e restituisce true se la stringa è palindroma,
e false altrimenti. Una parola è palindroma se può essere letta  sia da sinistra a destra che
da destra a sinistra. Ad esempio "osso" e "ingegni" sono palindrome, ma anche "" (la stringa vuota) e
"t" (le stringhe di un solo carattere).
*/
func isPalindrome(s string) bool {
	var reverse string
	for i := 0; i < len(s); i++ {
		reverse = string(s[i]) + reverse
	}
	if s == reverse {
		return true
	} else {
		return false
	}
}

func main() {

	var input string

	fmt.Scan(&input)

	if hasUpper(input) {
		fmt.Println("Ha maiuscole")
	} else {
		fmt.Println("Non ha maiuscole")
	}

	fmt.Println("Prima maiuscola in posizione", firstUpper(input))
	fmt.Println("Ultima maiuscola in posizione", lastUpper(input))
	digits, letters, others := countDigitsLettersOthers(input)
	fmt.Printf("Contiene %d Cifre, %d Lettere e %d altri caratteri", digits, letters, others)
	if isPalindrome(input) {
		fmt.Println("Palindroma")
	} else {
		fmt.Println("Non palindroma")
	}
	//fmt.Printf("Contiene %d Cifre, %d Lettere e %d altri caratteri",Wrap(countDigitsLettersOthers(input)...))

}

package main

/*
	Creare una dir stack in cui spostare la versione 3
	del programma stack.go, di cui dovete commentare
	la funzione main().
	Creare nella stessa dir (stack) un programma RPN.go
	e (con il comando go mod init RPN) un file go.mod con il seguente contenuto:
	module RPN
	go 1.17
	Per compilare il tutto dovrete dare (nella dir stack) il comando
	$ go build

	Il programma deve realizzare una calcolatrice in
	notazione polacca inversa (RPN), o notazione postfissa
	(https://it.wikipedia.org/wiki/Notazione_polacca_inversa),
	notazione in cui prima si inseriscono gli operandi e
	poi gli operatori: ad esempio 3 4 + 5 * invece che
	(3 + 4) * 5, e 3 4 5 * +  invece di 3 + 4 * 5.

	La calcolatrice deve disporre delle quattro operazioni
	aritmetiche (+ - * /) e operare su valori float64.
	Il funzionamento e` il seguente:
	ripetere:

	leggere un input dall'utente

	se e` un numero, lo mette in testa allo stack (una push)
	se e` un operatore,

	preleva gli operandi dallo stack (due pop)
	esegue l'operazione corrispondente all'operatore (+ - * /)
	salva il risultato nello stack (push)


	condizione di uscita: inserimento di "q" (quit)


	condizioni di errore: mancanza di operandi
	Nel caso l'input non sia ne' un numero ne' un operatore ne'
	"q", il programma lo ignora (non occorre che segnali errore).

	Per verificare se l'input e` un numero, utilizzare la funzione
	val, err := strconv.ParseFloat(input, 64)
	Nel caso err sia nil, l'input e un numero che verra salvato
	in val.
*/

import (
	"fmt"
	"strconv"
)

func main() {

	var input string
	var first *Element

	fmt.Println("NEXT? (+, -, *, /, q o un numero)")
	fmt.Scan(&input)

	for input != "q" {
		if input == "+" || input == "-" || input == "*" || input == "/" {
			if first == nil || first.next == nil {
				fmt.Println("not enough data")
			} else {
				var v1 float64
				var v2 float64
				first, v1 = popFunction(first)
				first, v2 = popFunction(first)
				switch input {
				case "+":
					first = pushFunction(first, (v1 + v2))
				case "-":
					first = pushFunction(first, (v1 - v2))
				case "*":
					first = pushFunction(first, (v1 * v2))
				case "/":
					first = pushFunction(first, (v1 / v2))
				}
			}
		} else {
			val, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("valore non è operando o cifra")
			} else {
				first = pushFunction(first, val)
			}
		}
		printStack(first)
		fmt.Println("NEXT? (+, -, *, /, q o un numero)")
		fmt.Scan(&input)

	}

}

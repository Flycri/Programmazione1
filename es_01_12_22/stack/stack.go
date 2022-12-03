package main

import (
	"fmt"
)

func viewMenu() {
	fmt.Printf("")
}

func main() {

	/*
		Creare un programma stack.go che gestisca uno stack (pila)
		generico (numero arbitrario di posizioni) di valori float
		Lo stack e` una struttura dati con le seguenti caratteristiche:

		mantiene nella prima posizione (detta testa) il valore piu' recente
		e` possibile accedere sempre e solo alla sua testa

		Le operazioni disponibili sullo stack sono:

		push: aggiunge un valore in testa allo stack
		pop: rimuove il valore in testa allo stack e lo restituisce
		top (o peek):  restituisce il valore in testa allo stack senza rimuoverlo
		empty: restituisce vero se lo stack e` vuoto, falso altrimenti

		Per ciascuna operazione creare una funzione, usando
		nomi maiuscoli (Push, Pop, ...).
		Implementare poi una funzione main() che chieda ripetutamente
		all'utente quale operazione vuole richiedere (push/pop/top/empty/quit),
		nel caso di push chieda anche il valore da aggiungere,
		stampi ogni volta il risultato e lo stack, e termini con quit.
	*/

	type Element struct {
		prev *Element
		val  float64
		succ *Element
	}
	var instruction string

	fmt.Printf("Operazione (PUSH/POP/TOP/EMPTY/QUIT)?")
	fmt.Scan(&instruction)

	for instruction != "quit" {

	}

}

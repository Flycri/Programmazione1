package main

import (
	"fmt"
)

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
	val  float64
	next *Element
}

func printStack(first *Element) {
	var elementptr *Element
	elementptr = first
	fmt.Printf("----------\n")
	for elementptr != nil {
		fmt.Printf("%f\n", elementptr.val)
		elementptr = elementptr.next
	}
	fmt.Printf("----------\n")
}

func pushFunction(first *Element, value float64) (newFirst *Element) {
	new_element := new(Element)
	*new_element = Element{val: value, next: first}
	return new_element
}

func popFunction(first *Element) (newFirst *Element, value float64) {
	if first != nil {
		return first.next, first.val
	} else {
		return nil, -1
	}
}

func topFunction(first *Element) {
	if first != nil {
		fmt.Printf("valore in cima alla pila:%f\n", first.val)
	}
}

func emptyFunction(first *Element) {
	var b bool
	if first == nil {
		b = true
	} else {
		b = false
	}

	fmt.Println("la pila è vuota? ", b)
}

/* func main() {

	var instruction string
	var first *Element

	fmt.Println("Operazione (PUSH/POP/TOP/EMPTY/QUIT)?")
	fmt.Scan(&instruction)

	for instruction != "QUIT" {
		switch instruction {
		case "PUSH":
			first = pushFunction(first)
			printStack(first)
		case "POP":
			first, _ = popFunction(first)
			printStack(first)
		case "TOP":
			topFunction(first)
			printStack(first)
		case "EMPTY":
			emptyFunction(first)
			printStack(first)
		default:
			fmt.Println("operazione sbagliata, sceglierne una corretta")
		}
		fmt.Println("Operazione (PUSH/POP/TOP/EMPTY/QUIT)?")
		fmt.Scan(&instruction)
	}

} */

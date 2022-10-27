package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
	Scrivere un programma indovina_numero.go che chiede all'utente di indovinare un numero intero random x
	tra 1 e MAX, (dove MAX è una costante definita nel programma) e ripete la richiesta fino a che l'utente
	indovina, oppure fino a un massimo di MAX/2 tentativi.
	Il programma stampa il numero di tentativi che sono stati necessari per indovinare oppure
	il messaggio "hai perso, il numero era x". Se il numero digitato dall'utente è fuori
	dall'intervallo [1,MAX], il tentativo non viene considerato e viene visualizzato il messaggio 
	"fuori intervallo!", senza interrompere l'esecuzione.
	Utilizzare la funzione rand.Intn del package "math/rand" per fissare il numero da indovinare.
	*/

	const MAX=10
	var x,n,i int
	var isRight bool
	rand.Seed(time.Now().UnixNano())
	x = rand.Intn(MAX)+1

	for i=1;(i<=MAX/2);i++ {
		fmt.Scan(&n)
		if n>MAX||n<1 {
			fmt.Println("fuori intervallo!")
		}
		if (n==x){
			isRight = true
			break;
		}
	}

	if (isRight) {
		fmt.Println("numero di tentativi:",i)
	} else {
		fmt.Println("hai perso il numero era",x)
	}
}
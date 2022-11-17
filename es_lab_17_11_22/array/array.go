package main


import (
	"fmt"
)

const DIM = 5

func reverse (aptr *[DIM]int) {
	for i:=0;i<len(*aptr)/2;i++ {
		(*aptr)[i],(*aptr)[len(*aptr)-i-1] = (*aptr)[len(*aptr)-i-1], (*aptr)[i]
	}
}

func switchFirstLast (aptr *[DIM]int) {
	(*aptr)[0],(*aptr)[len(*aptr)-1] = (*aptr)[len(*aptr)-1], (*aptr)[0]
}

func main() {
	/*
Scrivere un programma array.go dotato di:

- una costante DIM per la dimensione dell'array, dichiarata a livello di package

- una funzione main che inizializza a piacere un array di int di dimensione DIM 
(ad esempio con numeri 0, 1, 2, ...) e testa le due seguenti funzioni che lavorano 
**direttamente sull'array stesso**, senza quindi dichiarare e usare altri array. 
Il programma stampa l'array iniziale, l'array dopo aver invocato reverse e l'array 
dopo aver invocato switchFirstLast.

- una funzione reverse che inverte l'ordine dei valori in un array di dimensione DIM, mettendo il primo 
in fondo, il secondo in penultima posizione e così via, nell'array stesso

- una funzione switchFirstLast che scambia il primo con l'ultimo dei valori in un array di 
dimensione DIM nell'array stesso


Esempio di esecuzione

$ go run array.go

array: [0 1 2 3 4]
reverse: [4 3 2 1 0]
switchFirstLast: [0 3 2 1 4]

*/

array := [DIM]int{1,2,3,4,5}

fmt.Println(array)
reverse(&array)
fmt.Println(array)
switchFirstLast(&array)
fmt.Println(array)

}
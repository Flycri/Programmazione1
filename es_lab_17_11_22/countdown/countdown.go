package main

import (
	
)

func main() {

	/*
		Scrivere un programma timer.go dotato di:

		- una struttura `Clock` con tre campi `hour`, `min`, `sec` di tipo `int`, dichiarati in 
		quest'ordine

		- una funzione `countdown` (a cui passare un puntatore ad un Clock) che fa scalare 
		l'orario di un secondo, invocando opportunamente `updateMin` (vedi sotto) quando sono da 
		modificare anche i minuti

		- una funzione `updateMin` (idem) che fa scalare l'orario di un minuto, invocando opportunamente 
		`updateHour` (vedi sotto) quando sono da modificare anche le ore

		- una funzione `updateHour` (idem) che fa scalare l'orario di un'ora

		Stabilite voi la segnatura adeguata per le funzioni qui sopra.

		La funzione `main` chiede l'orario di partenza nel formato ore minuti secondi e fa partire 
		il countdown, stampando l'orario a ogni secondo, fino a raggiungere 0 0 0.

		Nota. Il programma deve creare un solo Clock e aggiornarne via via i valori dei campi, 
		non deve creare un nuovo Clock a ogni secondo.

		**Suggerimento**: usare l'istruzione `time.Sleep(time.Duration(1) * time.Second)` 
		(potete copiarla così come è) per far passare (circa) un secondo prima di stampare ogni 
		nuovo orario.

	*/

}



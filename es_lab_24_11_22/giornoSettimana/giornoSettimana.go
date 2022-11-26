package main

import (
	"fmt"
	"os"
)

var DAYS [7]string = [7]string{"lun", "mart", "merc", "giov", "ven", "sab", "dom"}

func returnDay(n, day_code int) string {
	x := (n % 7) + day_code

	if x > 6 {
		x -= 7
	}

	return DAYS[x]
}

func main() {

	/*
		Scrivere un programma giornoSettimana.go, dotato di un array giorniDellaSettimana
		[7]string{"lun", "mart", "merc", "giov", "ven", "sab", "dom"}
		che legge da linea di comando il nome del giorno della settimana del 1° gennaio di un dato anno. Se il giorno non è
		fra "lun", "mart", "merc", "giov", "ven", "sab", "dom", il programma avvisa e termina. Altrimenti poi accetta dall'utente
		dei numeri interi (tra 1 e 365), corrispondenti a giorni di quell'anno, e per ciascuno stampa il giorno della settimana
		corrispondente. Il programma termina quando l'utente digita -1.
	*/

	var day string
	var found bool
	var num, day_code int

	fmt.Println("Inserire il giorno della settimana del primo dell'anno")
	fmt.Scan(&day)

	for i := 0; i < len(DAYS); i++ {
		if day == DAYS[i] {
			found = true
			day_code = i
			break
		}
	}

	if !found {
		fmt.Println("valore giorno settimana errato, termine programma")
		os.Exit(1)
	}

	for {
		fmt.Println("inserire il numero di giorni a partire dal primo gennaio")
		fmt.Scan(&num)
		if num == -1 {
			os.Exit(0)
		}
		fmt.Println(returnDay(num, day_code))
	}

}

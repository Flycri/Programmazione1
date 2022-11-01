package main

import "fmt"

func main() {
	/*
		Scrivere un programma capitale.go che legge da stardard input tre valori float64 che rappresentano 
		un capitale iniziale (es. 1000.50 euro), un tasso di interesse annuale (es. 1.3%) e un obiettivo (es. 2000 euro), 
		e calcola quanti anni occorrono per arrivare a (o superare) l'obiettivo.
	*/

	var ini,yearInt,goal,sum float64
	var totYears int

	fmt.Scan(&ini,&yearInt,&goal)

	yearInt/=100
	sum = ini
	for sum<goal {
		sum+=yearInt*sum
		totYears++
	}

	fmt.Println(totYears,sum)
}
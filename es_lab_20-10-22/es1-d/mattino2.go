package main

import "fmt"

func main() {
	/*
	crivere un programma Go mattino2.go che legga un orario (due numeri interi, ora e minuto)
	e stampi "mattino" se l'orario è compreso tra le 5:30 (incluso) e le 12:30 (escluso)
	*/
	var h, min int
	fmt.Scan(&h)
	fmt.Scan(&min)
	if (h>=6)||(h<=11) {
		fmt.Println("mattino")
	} else {
		if h==5 {
			if min>=30 {
				fmt.Println("mattino")
			}
		}
		if h==12 {
			if min<30 {
				fmt.Println("mattino")
			}
		}
	}

	/* versione "furba"
	if ((h*100+min)>=530)&&((h*100+min)<1230) {
		fmt.Println("mattino")
	}
	*/
}
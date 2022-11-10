package main

import (
	"fmt"
)

/*
	Vogliamo scrivere una funzione 
	puntiCarta(s string) int 
	che accetta una stringa di 3 caratteri che rappresenta una mano di tre carte e restituisce il 
	punteggio complessivo relativo per il gioco della briscola. Ad esempio per la mano "53J" 
	restituisce 12 (10 della carta 3 + 2 del fante). Se la stringa contiene un simbolo che non corrisponde 
	a nessuna carta, la funzione restituisce -1.
*/

func punti(s string) int {
	var tot int
	for _,r:=range s {
		switch r {
		case 'a','A':
			tot+=11
		case '3':
			tot+=10
		case 'k','K':
			tot+=4
		case 'q','Q':
			tot+=3
		case 'j','J':
			tot+=2
		case '7','6','5','4','2':
			break;
		default:
			return -1
		}
	}
	return tot
}

func main() {
	/*
		Scrivi la funzione puntiCarta come da specifiche. Scrivi un main per invocare e testare la funzione.
		Il programma completo deve leggere da standard input una stringa e controllare che sia di tre 
		caratteri. Poi stampa "mano <mano>: <tot> punti" (dove <mano> è la stringa di 3 caratteri che 
		rappresenta una mano di tre carte), usando la funzione punti per determinare tot. Chiama il 
		programma briscola.go e caricalo su upload.
	*/
	
	var s string

	for {
		fmt.Scan(&s)
		if len(s)==3 {
			break;
		}
	}

	fmt.Printf("Mano %s: %d tot punti",s,punti(s))

}
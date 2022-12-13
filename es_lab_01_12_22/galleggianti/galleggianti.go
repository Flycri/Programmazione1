package main //TODO: FARE VERSIONE CON MATRICE, NON SO SE QUESTO è IL COMPORTAMENTO "CORRETTO"

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
	Scrivete un programma galleggianti.go che legga da standard input due interi r, c, seguiti da una matrice di r righe e c
	colonne contenente lettere maiuscole e asterischi, e che stampi in output la matrice che si ottiene da quella in input
	mandando verso il basso le lettere e facendo "galleggiare" gli asterischi, cioè mandandoli verso l'alto.
	Fornire in input una riga della matrice alla volta, e sfruttare la funzione strings.Split(s, " ") (vedi documentazione) per
	popolare le righe della matrice.
	Si assuma che ogni riga di input contenga esattamente c caratteri, separati da spazi (non sono richiesti controlli in tal senso).
*/

func PrintMatrix(matrix []string, c int) {
	for i := 0; i < len(matrix); i += c {
		for k := i; k < c+i; k++ {
			fmt.Print(matrix[k], " ")
		}
		fmt.Print("\n")
	}
}

func BubbleAsterisks(matrix []string) (res []string) {
	var asterisk_slice []string
	var letter_slice []string

	for i := 0; i < len(matrix); i++ {
		if matrix[i] == "*" {
			asterisk_slice = append(asterisk_slice, matrix[i])
		} else {
			letter_slice = append(letter_slice, matrix[i])
		}
	}

	res = append(asterisk_slice, letter_slice...)

	return
}

func main() {
	var c, r int

	var matrix []string
	first, second := true, false

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		var err error
		if first {
			r, err = strconv.Atoi(scanner.Text())
			first, second = false, true
			if err != nil {
				os.Exit(1)
			}
			_ = r
		} else if second {
			c, err = strconv.Atoi(scanner.Text())
			first, second = false, false
			if err != nil {
				os.Exit(1)
			}
		} else {
			matrix = append(matrix, scanner.Text())
		}
	}

	PrintMatrix(matrix, c)
	matrix = BubbleAsterisks(matrix)
	PrintMatrix(matrix, c)
}

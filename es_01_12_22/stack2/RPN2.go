package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	/*
		Creare una seconda versione RPN2.go in cui l'espressione
		da valutare sia fornita da linea di comando (os.Args)
		invece che da standard input (cioe` tastiera).
		In questo caso se l'espressione contiene moltiplicazioni,
		l'asterisco va messo tra apici ('') o tra virgolette ("").
	*/

	var first *Element

	for i := 1; i < len(os.Args); i++ {
		input := os.Args[i]
		if input == "+" || input == "-" || input == "*" || input == "/" {
			if first == nil || first.next == nil {
				fmt.Println("not enough data")
			} else {
				var v1 float64
				var v2 float64
				first, v1 = popFunction(first)
				first, v2 = popFunction(first)
				switch input {
				case "+":
					first = pushFunction(first, (v1 + v2))
				case "-":
					first = pushFunction(first, (v1 - v2))
				case "*":
					first = pushFunction(first, (v1 * v2))
				case "/":
					first = pushFunction(first, (v1 / v2))
				}
			}
		} else {
			val, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("valore non è operando o cifra")
			} else {
				first = pushFunction(first, val)
			}
		}
		printStack(first)

	}
}

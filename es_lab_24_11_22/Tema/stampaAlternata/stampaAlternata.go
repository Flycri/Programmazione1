package main

import (
	"fmt"
  "bufio"
  "os"
)

func main() {

	/*
		Scrivere un programma stampaAlternata.go che legge da standard input del testo su più righe
    (terminato da EOF) e stampa prima le righe pari e poi le righe dispari (considerate la prima riga del
    testo la riga 1 (e non 0)).
	*/

  var even_s, odd_s string
  var isEven bool

  scanner := bufio.NewScanner(os.Stdin)

  for scanner.Scan() {
    s := scanner.Text()
    if isEven {
      even_s+=s+"\n"
      isEven = false
    } else {
      odd_s+=s+"\n"
      isEven = true
    }

  }
  fmt.Print(even_s)
  fmt.Print(odd_s)

}

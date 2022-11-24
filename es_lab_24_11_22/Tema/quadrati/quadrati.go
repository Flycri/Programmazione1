package main

import (
	"fmt"
  "math"
  "os"
  "strconv"
)

func isSquare(n int) bool {
  return (math.Sqrt(float64(n))-float64(int(math.Sqrt(float64(n))))) == 0
}

func main() {
	/*
    Scrivere una programma quadrati.go che legge da linea di comando una sequenza di numeri interi
    non negativi e stampa solo quelli che sono dei quadrati (1, 4, 9, ecc.). Il programma deve essere dotato
    di una funzione isSquare(n int) bool che restituisce true se n è un quadrato, false altrimenti.
	*/

  /*var input_sequence []int
  var x int
  for i:=1; i<len(os.Args); i++ {
    x,_ = strconv.Atoi(os.Args[i])
    input_sequence = append(input_sequence, x)
  }*/
  for _,v:=range os.Args[1:] {
    x,err := strconv.Atoi(v)
    if err!= nil {
      fmt.Println(err)
      os.Exit(1)
    }
    if isSquare(x) {
    fmt.Println(x)
    }
  }

}

package main

import (
	"fmt"
  "bufio"
  "os"
  "strconv"
  "math"
)

func main() {

	/*
    Scrivere un programma polinomio.go che legge da standard input una riga che contiene dei numeri a, b, ....
    Il programma calcola il valore in x del polinomio
    a + bx + cx^2 + dx^3 ....
    corrispondente alla sequenza dei numeri letti tranne l'ultimo, dove il valore di x è l'ultimo valore della sequenza.
    Ad esempio,
    3 2 0 7 5
    corrisponde al polinomio 3 + 2x + 7x^3 da valutare per x = 5
	*/
  var coeff_slice []float64
  var tot float64
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanWords)
  for scanner.Scan(){
    x,err := strconv.ParseFloat(scanner.Text(),64)
    if err!=nil {
      fmt.Println(err)
      os.Exit(1)
    }
    coeff_slice = append(coeff_slice,x)
  }

  for i,val:=range coeff_slice[:len(coeff_slice)-1] {
    tot += math.Pow(coeff_slice[len(coeff_slice)-1],float64(i)) * val
  }

  fmt.Println(tot)

}

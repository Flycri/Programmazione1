package main

import (
)

func estraiPari(in []int)(out []int) {
  for _,val:=range in {
    if (val%2==0) {
      out = append(out,val)
    }
  }
  return
}

func rimuoviMultipli(m int, in[]int)(out []int) {
  for _,val:=range in {
    if (val%m!=0) {
      out = append(out,val)
    }
  }
  return
}

func main() {

	/*
    estraiPari(in []int) (out []int) che prende una slice di interi e ne restituisce una che contiene solo i nr.
    pari di quella in ingresso
    rimuoviMultipli(m int, in []int) (out []int) che prende un intero e una slice di interi e ne restituisce una che
    contiene solo i multipli del numero passato di quella in ingresso. Es.: rimuoviMultipli(5, in) restituisce una
    slice che contiene solo i multipli di 5 di in.
	*/



}

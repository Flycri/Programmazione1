package main

import "fmt"

func main() {
	/*
		Scrivere un programma cifra_pari.go che, dato un intero da standard input, determina e stampa in
		che posizione (procedendo da destra a sinistra) si trova la prima cifra pari del numero.
		Se il numero non contiene cifre pari, il programma stampa -1.
	*/

	var x,j int
	var containsEven bool

	fmt.Scan(&x)

	for i:=1;x/i>0;i*=10 {
		if (((x/i)%10)%2==0) {
			containsEven=true
			break;
		}
		j++
	}

	if containsEven {
		fmt.Println("il primo numero pari si trova in posizione",j+1,"(da dx)")
	} else {
		fmt.Println("-1")
	}

}
package main

import "fmt"

/*
	Stampare una slice ricorsivamente al contrario
*/

func InversePrint(slice []int) {
	l := len(slice)
	if l == 1 {
		fmt.Print(slice[0])
	} else {
		InversePrint(slice[1:])
		fmt.Print(slice[0])
	}

}

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	InversePrint(slice)

}

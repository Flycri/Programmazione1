package main

import (
	"fmt"
)

func main() {
	var x int
	const div = 1000

	fmt.Scan(&x)
	if (x%div == 0) {
		fmt.Println("il numero termina con 000")
	} else {
		fmt.Println("il numero non termina con 000")
	}
}
package main

import (
	"fmt"
)

func main() {
	var year int
	fmt.Scan(&year)
	if year%100 == 0 {
		if year%400 == 0 {
			fmt.Println("l'anno", year, "è un anno bisestile")
		} else {
			fmt.Println("l'anno", year, "non è un anno bisestile")
		}
	} else if year%4 == 0 {
		fmt.Println("l'anno", year, "è un anno bisestile")
	} else {
		fmt.Println("l'anno", year, "non è un anno bisestile")
	}
}

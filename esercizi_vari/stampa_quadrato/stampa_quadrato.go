package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	l, _ := strconv.Atoi(os.Args[1])

	for i := 0; i < l; i++ {
		if i == 0 || i == (l-1) {
			for j := 0; j < l; j++ {
				fmt.Print("* ")
			}
		} else {
			for j := 0; j < l; j++ {
				if j == 0 || j == (l-1) {
					fmt.Print("* ")
				} else {
					fmt.Print("  ")
				}
			}
		}
		fmt.Print("\n")

	}

}

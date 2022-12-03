package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var temp_tot int
	var tot_array [3]int

	for scanner.Scan() { //fa schifo, lo so, ma sono stanco
		if scanner.Text() == "" {
			if temp_tot > tot_array[2] {
				tot_array[0] = tot_array[1]
				tot_array[1] = tot_array[2]
				tot_array[2] = temp_tot
			} else if temp_tot > tot_array[1] {
				tot_array[0] = tot_array[1]
				tot_array[1] = temp_tot
			} else if temp_tot > tot_array[0] {
				tot_array[0] = temp_tot
			}
			temp_tot = 0
		} else {
			x, _ := strconv.Atoi(scanner.Text())
			temp_tot += x
		}
	}

	fmt.Println(tot_array[0] + tot_array[1] + tot_array[2])

}

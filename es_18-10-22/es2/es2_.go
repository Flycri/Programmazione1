package main

import (
	"fmt"
)

func main() {
	var x, y, z int
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scan(&z)
	if (x>y) {
		x,y = y,x
	}
	if (y>z) {
		y,z = z,y
	}
	if (x>y) {
		x,y = y,x
	}
	fmt.Println(x,y,z)
}
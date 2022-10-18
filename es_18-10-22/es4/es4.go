package main

import (
	"fmt"
)
 func main() {
	var x int
	fmt.Scan(&x)
	if ((x%10)%2!=0)&&((x/10%10)%2!=0)&&((x/100%10)%2!=0) {
		fmt.Println(x,"è composto solo da cifre dispari")
	} else {
		fmt.Println(x,"non è composto solo da cifre dispari")
	}
 }
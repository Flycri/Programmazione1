package main

import (
	"fmt"
)
 func main() {
	var max,min,i,mcd int
	fmt.Scan(&max)
	fmt.Scan(&min)
	if (max<min) {
		max,min = min,max
	}
	i = min
	for mcd==0 {
		if (max%i==0)&&(min%i==0) {
			mcd=i
		} else {
			i--
		}
	}
	fmt.Println("Il massimo comune divisore di",max,"e",min,"è:",mcd)
 }
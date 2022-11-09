package main

import "fmt"

func main() {
	var n, x, tempSum int

	maxSum := 0

	for {
		fmt.Scan(&n)
		if n == 999 {
			break
		}

		tempSum = 0

		for i := 1; n/i != 0; i *= 10 {
			x = n / i % 10
			tempSum += x
		}

		if tempSum > maxSum {
			maxSum = tempSum
		}
	}
	fmt.Println(maxSum)
}

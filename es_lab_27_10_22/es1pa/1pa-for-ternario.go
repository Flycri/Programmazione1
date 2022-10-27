package main
import "fmt"
func main() {
    /*
    Scrivere un programma che dato un numero n stampi tutti i numeri pari da 0 a n compreso
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 0; i <= n; i = i + 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

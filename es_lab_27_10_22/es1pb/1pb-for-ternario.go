package main
import "fmt"
func main() {
    /*
	Scrivere un programma che dato un numero intero in input stampi tutti i numeri da n a 1
	compreso separati da spazi
	*/
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := n; i >= 1; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

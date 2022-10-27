package main
import "fmt"
func main() {
    /*
	 Scrivere un programma che dato un intero n stampi a schermo n asterischi in fila
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

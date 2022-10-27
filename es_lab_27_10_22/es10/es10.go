package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
    /*
    Scrivere un programma es10.go che genera numeri casuali tra 1 e 10,
	stampandoli, fino a quando la loro somma raggiunge un obiettivo fissato (TARGET),
	ad esempio 50. Poi stampa la somma.
    */
    const TARGET = 50
    const MAX = 10
    rand.Seed(time.Now().Unix())
    var n int

    t := 0
    for t < TARGET {
        n = rand.Intn(MAX) + 1
        fmt.Print(n, " ")
        t += n
    }
    fmt.Println(t)
}

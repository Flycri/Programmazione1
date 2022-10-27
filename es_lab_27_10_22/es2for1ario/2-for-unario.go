package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
    /*
    Scrivere un programma che genera un numero random da 1 a 10 e che li stampi a schermo
	finchè la loro somma è maggiore o uguale a 20
    */
    const TARGET = 20
    const MAX = 10
    rand.Seed(time.Now().Unix())
    var n int

    sum := 0
    for sum < TARGET {
        n = rand.Intn(MAX) + 1
        fmt.Print(n, " ")
        sum += n
    }
    fmt.Println(sum)
}

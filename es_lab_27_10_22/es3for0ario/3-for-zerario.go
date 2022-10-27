package main
import (
    "fmt"
    "math/rand"
    "time"
)
func main() {
    /*
	Scrivere un programma che stampi a schermo numeri casuali da 0 a 20 e li stampi separati da spazi finchè
	non viene generato uno 0 ed infine stampi quanti numeri sono stati generati prima dello 0
	*/

    rand.Seed(time.Now().Unix())
    const K = 20
    var n int

    c := 0
    for {
        n = rand.Intn(K+1)
        fmt.Print(n, " ")
        if n == 0 {
            break
        }
        c++
    }
    fmt.Println(c)
}

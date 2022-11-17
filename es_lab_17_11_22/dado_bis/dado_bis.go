package main
 
import (
 	"fmt"
 	"math/rand"
 	"time"
)

  func main() {
 	const nFacce = 6
	var lancio, n int
	var frequenza [nFacce + 1]int

	fmt.Scan(&n)

	rand.Seed(time.Now().Unix())

	for i := 0; i < n; i++ {
		lancio = rand.Intn(nFacce) + 1
		frequenza[lancio]++
	}

	for i:=1;i<len(frequenza);i++ {
		fmt.Printf("%d: %d (%d %% ) ", i, frequenza[i],  frequenza[i]*100/n)
	}
}
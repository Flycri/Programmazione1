/* (Operatori di confronto e logici)
Specifiche: Scrivere un programma Go condizioni.go per testare, una a una, le condizioni nella tabella che segue.
Il programma, per ogni condizione, legge un valore da tastiera (del tipo indicato nella colonna “tipo”) e stampa true o false, a seconda che la condizione sia verificata o no.
Implementare, in un unico programma, una condizione alla volta (preceduta da un commento che indica la condizione da verificare), testarla su almeno due input (uno che la verifica e uno che la falsifica), e solo poi procedere alla successiva.
Scopo dell'esercizio è imparare a usare gli operatori di confronto uguale, diverso, maggiore, minore, maggiore o uguale, minore o uguale (==, !=, >, <, >=, <=) e quelli logici NOT, AND, OR (!, &&, ||).
Nota: ai fini dell'esercizio bastano 2 variabili, una di tipo int e una di tipo float64. */

package main

import (
	"fmt"
	"math"
)

func main() {
	var x int
	var f float64

	fmt.Println("Numero uguale a 10")
	fmt.Scan(&x)
	fmt.Println(x==10,"\n")

	fmt.Println("Numero diverso a 10")
	fmt.Scan(&x)
	fmt.Println(x!=10,"\n")

	fmt.Println("Numero diverso da 10 e da 20")
	fmt.Scan(&x)
	fmt.Println((x!=10)&&(x!=20),"\n")

	fmt.Println("Numero diverso da 10 o da 20")
	fmt.Scan(&x)
	fmt.Println((x!=10)||(x!=20),"\n") //sempre true

	fmt.Println("Numero maggiore o uguale a 10")
	fmt.Scan(&x)
	fmt.Println(x>=10,"\n")

	fmt.Println("Numero compreso fra 10 e 20 compresi")
	fmt.Scan(&x)
	fmt.Println((x>=10)&&(x<=20),"\n")

	fmt.Println("Numero compreso fra 10 e 20 esclusi")
	fmt.Scan(&x)
	fmt.Println((x>10)&&(x<20),"\n")

	fmt.Println("Numero compreso fra 10 e 20, 10 compreso")
	fmt.Scan(&x)
	fmt.Println((x>=10)&&(x<20),"\n")

	fmt.Println("Numero esterno a 10 e 20, compresi")
	fmt.Scan(&x)
	fmt.Println((x<=10)||(x>=20),"\n")

	fmt.Println("Numero compreso fra 10 e 20 compresi, o fra 30 e 40 compresi")
	fmt.Scan(&x)
	fmt.Println(((x>=10)&&(x<=20))||((x>=30)&&(x<=40)),"\n")

	fmt.Println("Numero multiplo di 4 ma non di 100")
	fmt.Scan(&x)
	fmt.Println((x%4==0)&&(x%100!=0),"\n")

	fmt.Println("Numero dispari e compreso fra 0 e 100 compresi")
	fmt.Scan(&x)
	fmt.Println((x%2==1)&&((x>=0)&&(x<=100)),"\n")

	fmt.Println("Numero 'vicino' a 10.0")
	fmt.Scan(&f)
	fmt.Println((f>10-math.Pow(10,-6))&&(f<10+math.Pow(10,-6)),"\n")
}
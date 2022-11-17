package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	/*
		Scrivere un programma usaStandard.go che dichiara le seguenti costanti:

		VOCALI = "aeiouAEIOU"\
		CIFRE = "0123456789"\
		una stringa S1 a piacere lunga 1 solo carattere
		un'altra stringa S2 a piacere, più lunga di 1.

		Il programma deve leggere da standard input una parola e utilizzare le funzioni dei pacchetti 
		strings e strconv per svolgere i seguenti compiti:
		Deve verificare

		se la stringa letta contiene S2 come sottostringa (nel caso stampare un messaggio)
		se la stringa letta contiene almeno una vocale (nel caso stampare un messaggio)

		Deve poi determinare
		3. quante occorrenze di S1 ha la stringa letta (e stampare un messaggio)
		4. in che posizione si trova la prima vocale della stringa letta
		5. in che posizione si trova l'ultima vocale della stringa letta
		e stampare i risultati ottenuti

		6. Deve stampare la stringa ottenuta concatenando S2 3 volte con se stessa.

		Deve infine estrarre le eventuali cifre contenute nella parola letta, concatenarle in una 
		stringa s nello stesso ordine in cui le incontra e
		7. stampare l'int n corrispondente a s usando fmt.Printf("%d\n", n)
		8. stampare il float64 f corrispondente a "0.s" usando fmt.Printf("%f\n", f).
	*/

	const VOCALI, CIFRE, S1, S2 = "aeiouAEIOU", "0123456789", "a", "hihi"

	var input,s string

	fmt.Scan(&input)

	fmt.Printf("La stringa %s contiene la sottostringa %s? %t \n",input,S2,strings.Contains(input,S2))

	fmt.Printf("La stringa %s contiene almeno una vocale? %t \n",input,strings.ContainsAny(input,VOCALI))

	fmt.Printf("La stringa %s contiene %d volte il carattere %s \n",input,strings.Count(input,S1),S1)

	fmt.Printf("La stringa %s contiene la prima vocale nella posizione %d \n",input,strings.IndexAny(input,VOCALI))

	fmt.Printf("La stringa %s contiene l'ultima vocale nella posizione %d \n",input,strings.LastIndexAny(input,VOCALI))

	fmt.Printf("Ripetizione di 3 volte di %s: %s \n",S2,strings.Repeat(S2,3))

	for _,r:=range input {
		println(string(r))
		if strings.ContainsAny(string(r),CIFRE) {
			s+=string(r)
		}
	}

	num1,_:=strconv.Atoi(s)
	fmt.Printf("int corrispondente di %s: %d \n",s,num1)

	num2,_:=strconv.ParseFloat("0."+s,64)
	fmt.Printf("float64 corrispondente di %s: %f \n",s,num2)

}
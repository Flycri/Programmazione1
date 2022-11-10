package main

import (
	"fmt"
	"strconv"
	"strings"
)

func giorniInMese(mese int) int {
	switch mese {
	case 2:
		return 28
	case 4, 6, 9, 11:
		return 30
	default:
		return 31
	}
}

func main() {
	/*
		Scrivi la funzione giorniInMese come da specifiche. Scrivi un main per invocare e testare la 
		funzione. Il programma deve leggere da standard input una stringa nel formato gg-mm-aaaa 
		(vedi funzione Atoi del pacchetto strconv) e stampare "il mese <x> ha <tot> giorni", dove x 
		e tot sono numeri, usando la funzione giorniInMese per determinare tot. Chiama il programma 
		giorni_mese.go e caricalo su upload
	*/

	var date, mmstring string
	var mm int

	fmt.Scan(&date)

	_,date,_ = strings.Cut(date,"-")
	mmstring,_,_ = strings.Cut(date,"-")
	mm,_ = strconv.Atoi(mmstring)

	fmt.Printf("il mese %d ha %d giorni",mm,giorniInMese(mm))

}
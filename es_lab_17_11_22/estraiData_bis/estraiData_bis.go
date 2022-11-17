package main

import (
   "fmt"
   "strings"
   "strconv"
)

func main() {

	/*
		Completare il programma implementando la funzione 
		func stringGMA2GMA(data string) (int, int, int)
		che, data una data nel formato giorno/mese/anno, restituisce i numeri corrispondenti a giorno,
		mese, anno, in quest'ordine.
		Il numero di cifre per giorno, mese, anno non è fissato, sono cioè valide date del 
		tipo 7/3/22, 07/03/2022, 7/03/2022, ...
		Sfruttare le funzioni opportune dei pacchetti strings e strconv.
	*/

   var dataGMA string
   fmt.Print("data nel formato giorno/mese/anno: ")
   fmt.Scan(&dataGMA)
   g, m, a := stringGMA2GMA(dataGMA)
   fmt.Println("giorno", g)
   fmt.Println("mese", num2mese(m))
   fmt.Println("anno", a)

}

func stringGMA2GMA(data string) (int, int, int) {
   var gmaSlice []string
   gmaSlice=strings.Split(data,"/")
   fmt.Println(gmaSlice)
   gmaSlice[0]=strings.TrimPrefix(gmaSlice[0],"0")
   gmaSlice[1]=strings.TrimPrefix(gmaSlice[1],"0")
   if len(gmaSlice[2])==2{
	gmaSlice[2]="20"+gmaSlice[2]
   }

   g,_ := strconv.Atoi(gmaSlice[0])
   m,_ := strconv.Atoi(gmaSlice[1])
   a,_ := strconv.Atoi(gmaSlice[2])

   return g,m,a
}

func num2mese(m int) string {
   var mesi = [13]string{"", "gen", "feb", "mar", "apr", "mag", "giu", "lug", "ago", "set", "ott", "nov", "dic"}
   return mesi[m]
}
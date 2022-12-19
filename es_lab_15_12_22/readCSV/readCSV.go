package main

import (
	"fmt"
	"math"
)

/*
Scrivere un programma `readCSV.go` che legga da stdin una sequenza di dati in formato CSV (Comma Separated Values) così organizzati:

timestamp,temperatura,umidità

Dove:
- `timestamp` è una stringa nel formato YYYY-MM-DD/HH:mm (anno mese giorno ora minuto)
- `temperatura` è un float
- `umidità` è un float
- il separatore è una virgola ","

Il programma deve usare Scanf per leggere ogni riga dell'input.

Il programma deve calcolare il massimo e il minimo dei valori di temperatura e umidità e stamparli con il timestamp
(in formato diverso dall'originale, deve essere nella forma "X(°/%) misura delle ore HH, minuti mm del giorno DD del mese MM
dell'anno YYYY") in cui sono stati misurati.

Ad esempio, il file fornito come input genera questo risultato:

minTemp: 1.0° misura delle ore 4, minuti 22 del giorno 11 del mese 12 dell'anno 2022
maxTemp: 15.0° misura delle ore 9, minuti 22 del giorno 11 del mese 12 dell'anno 2022
minHumid: 49.0% misura delle ore 9, minuti 24 del giorno 11 del mese 12 dell'anno 2022
maxHumid: 91.0% misura delle ore 9, minuti 31 del giorno 6 del mese 12 dell'anno 2022
*/
type Data struct {
	year  int
	month int
	day   int
	hour  int
	min   int
	temp  float64
	hum   float64
}

func (d Data) ToString(temp bool) string {
	if temp {
		return fmt.Sprintf("%.1f° misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d", d.temp, d.hour, d.min, d.day, d.month, d.year)
	} else {
		return fmt.Sprintf("%.1f%% misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d", d.hum, d.hour, d.min, d.day, d.month, d.year)
	}
}

func main() {

	var year, month, day, hour, min int
	var temp, hum float64
	var max_temp_data, max_hum_data, min_temp_data, min_hum_data Data
	fmt.Scanf("%d-%d-%d/%d:%d,%f,%f", &year, &month, &day, &hour, &min, &temp, &hum)
	min_temp_data.temp = math.MaxInt
	min_hum_data.hum = math.MaxInt
	for {
		_, err := fmt.Scanf("%d-%d-%d/%d:%d,%f,%f", &year, &month, &day, &hour, &min, &temp, &hum)
		if err != nil {
			break
		}
		if temp > max_temp_data.temp {
			max_temp_data = Data{year: year, month: month, day: day, hour: hour, min: min, temp: temp, hum: hum}
		}
		if hum > max_hum_data.hum {
			max_hum_data = Data{year: year, month: month, day: day, hour: hour, min: min, temp: temp, hum: hum}
		}

		if temp < min_temp_data.temp {
			min_temp_data = Data{year: year, month: month, day: day, hour: hour, min: min, temp: temp, hum: hum}
		}
		if hum < min_hum_data.hum {
			min_hum_data = Data{year: year, month: month, day: day, hour: hour, min: min, temp: temp, hum: hum}
		}
	}

	fmt.Println("minTemp:", min_temp_data.ToString(true))
	fmt.Println("maxTemp:", max_temp_data.ToString(true))
	fmt.Println("minHumid:", min_hum_data.ToString(false))
	fmt.Println("maxHumid:", max_hum_data.ToString(false))

}

package main

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"net/http"

	"github.com/holizz/terrapin"
)

func pippoFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
	<!doctype html>	
	<title>Titolo</title>
	<h1>TOP TEXT</h1>
	<p>paragrafo</p>
	<img style="height:500px; width:500px" src="https://static.nexilia.it/nextquotidiano/2018/10/melonichan-giorgia-meloni-mari-love-live-lesbica-gender-11.jpg?impolicy=nexilia-4-3">
	<footer><h1>BOTTOM TEXT</h1></footer>
	`))
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
	<!doctype html>	
	<title>Buon Natale</title>
	<h1>Buon Natale</h1>
	<img src="/mainImage.png">
	<footer><h1>oof</h1></footer>
	`))
}

func mainImage(w http.ResponseWriter, r *http.Request) {
	campo := image.NewRGBA(image.Rect(0, 0, 500, 500))
	t := terrapin.NewTerrapin(campo, terrapin.Position{250.0, 450.0})

	kochSnowflake(t, 10, 5)

	png.Encode(w, campo)
}

func kochSnowflake(t *terrapin.Terrapin, l float64, liv int) {
	t.Right(math.Pi - math.Pi/3.0)
	kochCurve(t, l, liv)
	t.Right(math.Pi - math.Pi/3.0)
	kochCurve(t, l, liv)
	t.Right(math.Pi - math.Pi/3.0)
	kochCurve(t, l, liv)
}

func kochCurve(t *terrapin.Terrapin, l float64, liv int) {
	if liv == 0 {
		t.Forward(l)
		return
	}
	kochCurve(t, l, liv-1)
	t.Left(math.Pi / 3.0)
	kochCurve(t, l, liv-1)
	t.Right(math.Pi - math.Pi/3.0)
	kochCurve(t, l, liv-1)
	t.Left(math.Pi / 3.0)
}

func main() {

	/*
		Specifica.
	*/

	fmt.Println("Listening on http://localhost:3000/")
	http.HandleFunc("/pippo", pippoFunc)
	http.HandleFunc("/main", mainPage)
	http.HandleFunc("/mainImage.png", mainImage)
	http.ListenAndServe(":3000", nil)

}

package main

import "fmt"

// Punto è un punto nel piano cartesiano
type Punto struct {
	x, y float64
}

func main() {
	var p1, p2, p3 Punto

	fmt.Print("x e y di P1: ")
	fmt.Scan(&p1.x, &p1.y)

	fmt.Print("x e y di P2: ")
	fmt.Scan(&p2.x, &p2.y)

	fmt.Println(funzUno(p1, p2))

	funzDue(&p1, 2, 3)
	fmt.Println(p1)

	p3 = newPunto(20,30)
	fmt.Println(p3)

	specchiaPunto(&p3)
	fmt.Println(p3)

}

func funzUno(p1, p2 Punto) Punto {
	var p Punto
	p.x = (p1.x + p2.x) / 2
	p.y = (p1.y + p2.y) / 2
	return p
}

func funzDue(p *Punto, horiz, vert float64) {
	(*p).x += horiz
	(*p).y += vert
}

func newPunto(x, y float64) Punto {
	p := Punto{x: x, y: y}
	return p
}

func specchiaPunto(p *Punto) {
	(*p).x  *=-1
}
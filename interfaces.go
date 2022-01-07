package figuras

import "fmt"

type Geometrica interface {
	area() float64
	perimetro() float64
}

func Result(g Geometrica) {
	fmt.Println("Medidas:", g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimetro:", g.perimetro())
}

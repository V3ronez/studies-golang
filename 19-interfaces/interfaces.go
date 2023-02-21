package main

import (
	"fmt"
	"math"
)

type forma interface { // interfaces tem apenas assinaturas de metodos
	area() float64
}
type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func escreverArea(f forma) {
	fmt.Printf("A area do é %0.2f\n", f.area())
}

func (r retangulo) area() float64 {
	return r.altura + r.largura
}
func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{altura: 10, largura: 15}
	escreverArea(r)
	c := circulo{raio: 16}
	escreverArea(c)
}

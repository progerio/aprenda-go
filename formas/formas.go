package formas

import "math"

type Retangulo struct {
	Largura float64
	Altura  float64
}

// metodo do area do retangulo
func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

type Circulo struct {
	Raio float64
}

// metodo area do circulo
func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)
}

type Forma interface {
	Area() float64
}

type Triangulo struct {
	Base   float64
	Altura float64
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) * 0.5
}

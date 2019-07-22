package area

import "math"

// Pi é uma proporção numerica definida pela relacao entre
//o perimetro deum circunferencia e seu diametro
const Pi = 3.1416

// Circ é responsável por cacular a area de circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsável por cacular a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é publicamente visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}

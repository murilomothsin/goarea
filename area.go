package goarea

import "math"

// Pi 'e uma proporcao numerica definida pela relacao entre
// o perimetro de uma circunferencia e seu diametro
const Pi = 3.1416

// Circ calcula a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect calcula a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Nao 'e visivel!
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}

package area

import "math"

// Pi é uma proporção numérica definida pela relação entre o perímetro de uma
// circunferência e o seu diâmetro.
const Pi = 3.1416

// Circle é responsável por calcular a área da circunferência.
func Circle(ray float64) float64 {
	return math.Pow(ray, 2) * Pi
}

// Retangle é responsável por calcular a área de um retângulo.
func Retangle(base, height float64) float64 {
	return base * height
}

// Triangle é responsável por calcular a área de um triângulo.
func Triangle(base, height float64) float64 {
	return triangleEq(base, height)
}

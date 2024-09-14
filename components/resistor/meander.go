package resistor

import "math"

func CountNumberOfLinks(a float64, b float64, l float64) float64 {
	return math.Sqrt((math.Pow(a, 2)/(4*math.Pow((a+b), 2)))+((l)/(a+b))) - (a / (2 * (a + b)))
}

func CountXlengthMeander(a float64, b float64, n float64) float64 {
	return n * (a + b)
}

func CountYlengthMeander(a float64, l float64, n float64) float64 {
	return (l - a*n) / n
}

package resistor

import "math"

func CountXstripes(a float64, b float64, formFactor float64) float64 {
	return b * (1 + math.Sqrt(1+formFactor*(a/b+1)))
}

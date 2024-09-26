package resistor

import "math"

type Meander struct {
	numberOfLinks  float64
	meanderXLength float64
	meanderYLength float64
	meanderArea    float64
}

func (r *Resistor) InitMeander(a float64, b float64) {
	r.meander.numberOfLinks = CountNumberOfLinks(a, b, b*r.GetFromFactor())
	r.meander.meanderXLength = CountXlengthMeander(a, b, r.GetNumberOfLinks())
	r.meander.meanderYLength = CountYlengthMeander(a, b*r.GetFromFactor(), r.GetNumberOfLinks())
	r.meander.meanderArea = CountArea(r.GetXlengthMeander(), r.GetYlengthMeander())
}

func CountNumberOfLinks(a, b, l float64) float64 {
	return math.Ceil(math.Sqrt((math.Pow(a, 2)/(4*math.Pow((a+b), 2)))+((l)/(a+b))) - (a / (2 * (a + b))))
}

func CountXlengthMeander(a, b, n float64) float64 {
	return n * (a + b)
}

func CountYlengthMeander(a, l, n float64) float64 {
	return (l - a*n) / n
}

func CountArea(x, y float64) float64 {
	return x * y
}

func (r *Resistor) GetNumberOfLinks() float64 {
	return r.meander.numberOfLinks
}

func (r *Resistor) GetXlengthMeander() float64 {
	return r.meander.meanderXLength
}

func (r *Resistor) GetYlengthMeander() float64 {
	return r.meander.meanderYLength
}

func (r *Resistor) GetMeanderArea() float64 {
	return r.meander.meanderArea
}

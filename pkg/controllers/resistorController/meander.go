package resistorcontrollerIml

import "math"

// func initMeander(a float64, b float64) {
// 	r.meander.numberOfLinks = CountNumberOfLinks(a, b, b*r.GetFromFactor())
// 	r.meander.meanderXLength = CountXlengthMeander(a, b, r.GetNumberOfLinks())
// 	r.meander.meanderYLength = CountYlengthMeander(a, b*r.GetFromFactor(), r.GetNumberOfLinks())
// 	r.meander.meanderArea = CountArea(r.GetXlengthMeander(), r.GetYlengthMeander())
//}

func countNumberOfLinks(a, b, l float64) float64 {
	return math.Ceil(math.Sqrt((math.Pow(a, 2)/(4*math.Pow((a+b), 2)))+((l)/(a+b))) - (a / (2 * (a + b))))
}

func countXlengthMeander(a, b, n float64) float64 {
	return n * (a + b)
}

func countYlengthMeander(a, l, n float64) float64 {
	return (l - a*n) / n
}

func countArea(x, y float64) float64 {
	return x * y
}

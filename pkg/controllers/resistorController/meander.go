package resistorcontrollerIml

import (
	"math"
	"github.com/shpakunya/pkg/components"
)

func meanderInit(resistor *components.Resistor) {
	resistor.Meander.NumberOfLinks = components.Parameter{Name: "оптимальное число Z-образных звеньев", Symbol: "n",
		Unit: "", Value: countNumberOfLinks(resistor.Rectangle.Width.Value, resistor.Rectangle.Width.Value, resistor.Rectangle.Width.Value * resistor.FormFactor.Value)}
	resistor.Meander.MeanderXLength = components.Parameter{Name: "Длинна контура", Symbol: "X",
		Unit: "мм", Value: countXlengthMeander(resistor.Rectangle.Width.Value, resistor.Rectangle.Width.Value, resistor.Meander.NumberOfLinks.Value)}
	resistor.Meander.MeanderYLength = components.Parameter{Name: "Высота контура", Symbol: "Y",
		Unit: "мм", Value: countYlengthMeander(resistor.Rectangle.Width.Value, resistor.Rectangle.Width.Value * resistor.FormFactor.Value, resistor.Meander.NumberOfLinks.Value)}
	resistor.Meander.MeanderArea = components.Parameter{Name: "Площадь контура", Symbol: "S",
		Unit: "мм²", Value: countArea(resistor.Meander.MeanderXLength.Value, resistor.Meander.MeanderYLength.Value)}
}

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

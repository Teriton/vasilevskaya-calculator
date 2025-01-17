package resistorcontrollerIml

import (
	"math"

	"github.com/shpakunya/pkg/components"
)

type ResistorController struct {
}

func (ResistorController) NewResistor(resistance float64, tolerance float64, power float64,
	material components.ResistorMaterial, env components.Environment) components.Resistor {
	// Init params
	res := components.Resistor{}
	res.Resistance = components.Parameter{Name: "Сопротивление", Symbol: "R", Unit: "Ом", Value: resistance}
	res.Tolerance = components.Parameter{Name: "Допуск", Symbol: "Δ", Unit: "%", Value: tolerance}
	res.Power = components.Parameter{Name: "Рассеиваемая мощность", Symbol: "Pрас", Unit: "мВт", Value: power}
	res.Material = material
	res.Env = env

	autoCalculateInit(&res)
	return res
}

func autoCalculateInit(resistor *components.Resistor) {
	// Main
	resistor.FormFactor = components.Parameter{Name: "Коэфициент формы", Symbol: "k",
		Unit: " ", Value: countFormFactor(resistor.Resistance.Value, resistor.Material.SquareResistance.Value)}
	resistor.GammaRt = components.Parameter{Name: "Температурная погрешность сопротивления резистора", Symbol: "γRt",
		Unit: "%", Value: countGammaRt(resistor.Material.TemperatureCoefficientOfResistance.Value, resistor.Env.Temperature.Value)}
	resistor.GammaRdelta = components.Parameter{Name: "Погрешность воспроизведения геометрических разыеров резис­тора", Symbol: "γRΔ",
		Unit: "%", Value: countGammaRDelta(resistor.Tolerance.Value, resistor.Env.GammaRokv.Value,
			resistor.Env.GammaRcontact.Value, resistor.Material.Senescence.Value, resistor.GammaRt.Value)}
	resistor.FormOfResistor = CountFormOfResistor(resistor.FormFactor.Value)

	// Additional
	initRectange(resistor)
}

func countFormFactor(resistance float64, squareResistance float64) float64 {
	return resistance / squareResistance
}

func countGammaRDelta(tolerance float64, gammaRokv float64, gammaRcontact float64, senescence float64, gammaRt float64) float64 {
	return tolerance - (gammaRokv + gammaRcontact + senescence + gammaRt)
}

func countGammaRt(tkr float64, temperature float64) float64 {
	return math.Abs((tkr * math.Pow(10, -4) * (temperature - 20)) * 100)
}

func CountFormOfResistor(formFactor float64) components.Form {
	switch {
	case formFactor > 0.1 && formFactor < 10:
		return components.RectangleForm
	case formFactor <= 0.1:
		return components.CCPForm
	case formFactor >= 10:
		return components.MeanderForm
	default:
		return "undefined"
	}
}

func tehnRound(num float64, techRound float64) float64 {
	roundness := math.Round(1.0 / techRound)
	return math.Ceil(num*roundness) / roundness
}

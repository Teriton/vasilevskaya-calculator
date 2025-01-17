package resistorcontrollerIml

import (
	"math"

	"github.com/shpakunya/pkg/components"
)

func initRectange(resistor *components.Resistor) {
	resistor.Rectangle.Bp = components.Parameter{Name: "Минимальная ширина резистора, при которой рассеивается заданная мощность", Symbol: "bp",
		Unit: "мм", Value: bpCount(resistor.Material.SquareResistance.Value, resistor.Power.Value,
			resistor.Resistance.Value, resistor.Material.PermissibleSpecificPowerDissipation.Value)}
	resistor.Rectangle.Bdelta = components.Parameter{Name: "Минимальная ширина резистора, при которой обеспечивается заданная точность", Symbol: "bΔ",
		Unit: "мм", Value: bdeltaCount(resistor.FormFactor.Value, resistor.GammaRdelta.Value, resistor.Env.Deltal.Value, resistor.Env.Deltab.Value)}
	resistor.Rectangle.Width = components.Parameter{Name: "Ширина", Symbol: "b",
		Unit: "мм", Value: tehnRound(bCount(resistor.Rectangle.Bp.Value, resistor.Rectangle.Bdelta.Value, resistor.Env.Btehn.Value), resistor.Env.Btehn.Value)}
	resistor.Rectangle.Lp = components.Parameter{Name: "Минимальная длинна резистора, при которой рассеивается заданная мощность", Symbol: "lp",
		Unit: "мм", Value: lpCount(resistor.Material.SquareResistance.Value, resistor.Power.Value,
			resistor.Resistance.Value, resistor.Material.PermissibleSpecificPowerDissipation.Value)}
	resistor.Rectangle.Ldelta = components.Parameter{Name: "Минимальная длинна резистора, при которой обеспечивается заданная точность", Symbol: "lΔ",
		Unit: "мм", Value: ldeltaCount(resistor.FormFactor.Value, resistor.GammaRdelta.Value, resistor.Env.Deltal.Value, resistor.Env.Deltab.Value)}
	resistor.Rectangle.Height = components.Parameter{Name: "Длинна", Symbol: "l",
		Unit: "мм", Value: tehnRound(lCount(resistor.Rectangle.Bp.Value, resistor.Rectangle.Ldelta.Value, resistor.Env.Ltehn.Value), resistor.Env.Ltehn.Value)}
}

func bpCount(squareResistance, power, resistance, permissibleSpecificPowerDissipation float64) float64 {
	return 10 * math.Sqrt((squareResistance*power)/(resistance*permissibleSpecificPowerDissipation))
}

func lpCount(squareResistance, power, resistance, permissibleSpecificPowerDissipation float64) float64 {
	return 10 * math.Sqrt((resistance*power)/(squareResistance*permissibleSpecificPowerDissipation))
}

func bdeltaCount(formFactor, gammaRdelta, deltal, deltab float64) float64 {
	return (deltab + (deltal / formFactor)) / (gammaRdelta / 100.0)
}

func ldeltaCount(formFactor, gammaRdelta, deltal, deltab float64) float64 {
	return (deltab + (deltal * formFactor)) / (gammaRdelta / 100.0)
}

func bCount(bp float64, bdelta float64, btehn float64) float64 {
	return math.Max(bp, math.Max(bdelta, btehn))
}

func lCount(lp, ldelta, ltehn float64) float64 {
	return math.Max(lp, math.Max(ldelta, ltehn))
}

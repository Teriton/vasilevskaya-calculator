package resistorcontrollerIml

import (
	"math"

	"github.com/shpakunya/pkg/components"
)

func ccpInit(resistor *components.Resistor) {
	resistor.Ccp.BpCCP = components.Parameter{Name: "Минимальная ширина резистора, при которой рассеивается заданная мощность", Symbol: "bp",
		Unit: "мм", Value: bpCount(resistor.Material.SquareResistance.Value, resistor.Power.Value,
			resistor.Resistance.Value, resistor.Material.PermissibleSpecificPowerDissipation.Value)}
	resistor.Ccp.BdeltaCCP = components.Parameter{Name: "Минимальная ширина резистора, при которой обеспечивается заданная точность", Symbol: "bΔ",
		Unit: "мм", Value: bdeltaCppCount(resistor.Env.Deltal.Value, resistor.Env.Deltab.Value, 0.04, resistor.Tolerance.Value, resistor.Env.GammaRokv.Value)}
}

func bpCppCount(squareResistance, power, resistance, permissibleSpecificPowerDissipation float64) float64 {
	return bpCount(squareResistance, power, resistance, permissibleSpecificPowerDissipation) / 4.0
}

func bdeltaCppCount(deltaB, deltaL, k, gammaR, gammaRokv float64) float64 {
	return (deltaB + (deltaL / k)) / ((gammaR/100.0 - gammaRokv/100.0) * (1.0/(2.0*k) + 1.0) * math.Log(1+2*k))
}


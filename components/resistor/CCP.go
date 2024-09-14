package resistor

import (
	"example/main/environment"
	"math"
)

func CountBpCCP(resistance float64, P float64, material material) float64 {
	return (1.0 / 4.0 * math.Sqrt((material.squareResistance*P)/(resistance*material.permissibleSpecificPowerDissipation))) * 10
}

func CountBdeltaCCP(deltaB float64, deltaL float64, k float64, gammar float64, env environment.Environment) float64 {
	return (deltaB + (deltaL / k)) / ((gammar/100.0 - env.GetGammaRokv()/100.0) * (1.0/(2.0*k) + 1.0) * math.Log(1-2*k))
}

package resistor

import (
	"math"
)

type CCP struct {
	bpCCP     float64
	bdeltaCCP float64
}

func (r *Resistor) initCCP() {
	r.ccp.bpCCP = CountBpCCP(r.resistance, r.power, r.material)
	r.ccp.bdeltaCCP = CountBdeltaCCP(r.environment.GetDeltab(), r.environment.GetDeltal(), 1, r.tolerance, r.environment.GetGammaRokv())
}

func CountBpCCP(resistance, P float64, material material) float64 {
	return (1.0 / 4.0 * math.Sqrt((material.squareResistance*P)/(resistance*material.permissibleSpecificPowerDissipation))) * 10
}

func CountBdeltaCCP(deltaB, deltaL, k, gammar, gammaRokv float64) float64 {
	return (deltaB + (deltaL / k)) / ((gammar/100.0 - gammaRokv/100.0) * (1.0/(2.0*k) + 1.0) * math.Log(1+2*k))
}

func (r *Resistor) GetBpCCP() float64 {
	return r.ccp.bpCCP
}

func (r *Resistor) GetBdeltaCCP() float64 {
	return r.ccp.bdeltaCCP
}

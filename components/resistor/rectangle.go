package resistor

import "math"

func (r *Resistor) bpCount() float64 {
	return 10 * math.Sqrt((r.material.squareResistance*r.power)/(r.resistance*r.material.permissibleSpecificPowerDissipation))
}

func (r *Resistor) lpCount() float64 {
	return 10 * math.Sqrt((r.resistance*r.power)/(r.material.squareResistance*r.material.permissibleSpecificPowerDissipation))
}

func (r *Resistor) bdeltaCount() float64 {
	var deltal, deltab float64 = 0.01, 0.01
	return (deltab + (deltal / r.GetFromFactor())) / (r.GetGammaRdelta() / 100.0)
}

func (r *Resistor) ldeltaCount() float64 {
	var deltal, deltab float64 = 0.01, 0.01
	return (deltab + (deltal * r.GetFromFactor())) / (r.GetGammaRdelta() / 100.0)
}

func (r *Resistor) bCount() float64 {
	switch {
	case r.bp > r.bdelta:
		return r.bp
	case r.bdelta > r.btehn:
		return r.bdelta
	default:
		return r.btehn
	}
}

func (r *Resistor) lCount() float64 {
	switch {
	case r.lp > r.ldelta:
		return r.lp
	case r.ldelta > r.ltehn:
		return r.ldelta
	default:
		return r.ltehn
	}
}

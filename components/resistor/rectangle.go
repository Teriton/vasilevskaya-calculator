package resistor

import "math"

type Rectangle struct {
	bp     float64
	bdelta float64
	width  float64
	lp     float64
	ldelta float64
	height float64
}

func (r *Resistor) initRectange() {
	r.rectangle.bp = BpCount(r.material.squareResistance, r.power, r.resistance, r.material.permissibleSpecificPowerDissipation)
	r.rectangle.bdelta = BdeltaCount(r.formFactor, r.gammaRdelta, r.environment.GetDeltal(), r.environment.GetDeltab())
	r.rectangle.width = TehnRound(BCount(r.GetBp(), r.GetBdelta(), r.btehn))
	r.rectangle.lp = LpCount(r.material.squareResistance, r.power, r.resistance, r.material.permissibleSpecificPowerDissipation)
	r.rectangle.ldelta = LdeltaCount(r.formFactor, r.gammaRdelta, r.environment.GetDeltal(), r.environment.GetDeltab())
	r.rectangle.height = TehnRound(r.lCount(r.GetLp(), r.GetLdelta(), r.ltehn))
}

func BpCount(squareResistance, power, resistance, permissibleSpecificPowerDissipation float64) float64 {
	return 10 * math.Sqrt((squareResistance*power)/(resistance*permissibleSpecificPowerDissipation))
}

func LpCount(squareResistance, power, resistance, permissibleSpecificPowerDissipation float64) float64 {
	return 10 * math.Sqrt((resistance*power)/(squareResistance*permissibleSpecificPowerDissipation))
}

func BdeltaCount(formFactor, gammaRdelta, deltal, deltab float64) float64 {
	return (deltab + (deltal / formFactor)) / (gammaRdelta / 100.0)
}

func LdeltaCount(formFactor, gammaRdelta, deltal, deltab float64) float64 {
	return (deltab + (deltal * formFactor)) / (gammaRdelta / 100.0)
}

func BCount(bp float64, bdelta float64, btehn float64) float64 {
	return math.Max(bp, math.Max(bdelta, btehn))
}

func (r *Resistor) lCount(lp, ldelta, ltehn float64) float64 {
	return math.Max(lp, math.Max(ldelta, ltehn))
}

// Getters
func (r *Resistor) GetBp() float64 {
	return r.rectangle.bp
}

func (r *Resistor) GetBdelta() float64 {
	return r.rectangle.bdelta
}

func (r *Resistor) GetWidth() float64 {
	return r.rectangle.width
}

func (r *Resistor) GetLp() float64 {
	return r.rectangle.lp
}

func (r *Resistor) GetLdelta() float64 {
	return r.rectangle.ldelta
}

func (r *Resistor) GetHeight() float64 {
	return r.rectangle.height
}

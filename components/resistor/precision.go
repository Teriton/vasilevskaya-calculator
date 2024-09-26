package resistor

type Precision struct {
	bmin, bmax       float64
	lmin, lmax       float64
	rokvmin, rokvmax float64
	rmin, rmax       float64
}

func (r *Resistor) precisionInit() {
	r.precision.bmax = r.GetWidth() + r.environment.GetDeltab()
	r.precision.bmin = r.GetWidth() - r.environment.GetDeltab()
	r.precision.lmin = r.GetHeight() - r.environment.GetDeltal()
	r.precision.lmax = r.GetHeight() + r.environment.GetDeltal()
	r.precision.rokvmax = r.environment.GetGammaRokv()*0.01*r.material.squareResistance + r.material.squareResistance
	r.precision.rokvmin = r.material.squareResistance - r.environment.GetGammaRokv()*0.01*r.material.squareResistance
	r.precision.rmax = r.resistance * (1 + (r.tolerance-r.gammaRt-r.material.senescence)/100)
	r.precision.rmin = r.resistance * (1 - (r.tolerance-r.gammaRt-r.material.senescence)/100)
}

// Getters

func (r *Resistor) Getbmax() float64 {
	return r.precision.bmax
}

func (r *Resistor) Getbmin() float64 {
	return r.precision.bmin
}

func (r *Resistor) Getlmin() float64 {
	return r.precision.lmin
}

func (r *Resistor) Getlmax() float64 {
	return r.precision.lmax
}

func (r *Resistor) Getrokvmax() float64 {
	return r.precision.rokvmax
}

func (r *Resistor) Getrokvmin() float64 {
	return r.precision.rokvmin
}

func (r *Resistor) Getrmin() float64 {
	return r.precision.rmin
}

func (r *Resistor) Getrmax() float64 {
	return r.precision.rmax
}

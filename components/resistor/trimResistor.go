package resistor

import (
	"example/main/environment"
	"fmt"
	"math"
)

func (r *Resistor) InitTrim(bdelta float64) {
	r.gammaRdeltaTrim = CalculateGammaRdeltaTrim(r.formFactor, r.width, bdelta)
	r.gammaR = CalculateGammaR(r.material, r.gammaRdeltaTrim, *r.environment, r.gammaRt)
	r.mOfTrim = CalculateMofTrim(r.gammaR, r.gammaRdeltaTrim)
	bmax := r.width + r.environment.GetDeltab()
	bmin := r.width - r.environment.GetDeltab()
	lmin := r.height - r.environment.GetDeltal()
	rokvmax := r.environment.GetGammaRokv()*0.01*r.material.squareResistance + r.material.squareResistance
	rokvmin := r.material.squareResistance - r.environment.GetGammaRokv()*0.01*r.material.squareResistance
	// rmax := r.resistance + r.resistance*r.gammaR*0.01
	// rmin := r.resistance - r.resistance*r.gammaR*0.01
	rmax := r.resistance * (1 + (r.tolerance-r.gammaRt-r.material.senescence)/100)
	rmin := r.resistance * (1 - (r.tolerance-r.gammaRt-r.material.senescence)/100)
	r.lnTrim = CalculateLn(rmax, bmin, rokvmax)
	r.rdashminTrim = CalculateRdashmin(rokvmin, r.lnTrim, bmax)
	r.deltarTrim = CalculateDeltarTrim(rmin, r.rdashminTrim)
	// BackThen
	r.loTrim = CalculateLo(rmin, bmax, rokvmin)
	r.ltune = CalculateLtune(r.loTrim, r.lnTrim)
	r.deltaLrTrim = CalculateDeltaLr(r.ltune, r.numberOfLinks)
	r.deltaLdashTrim = CalculateDeltaLdash(rokvmin, r.deltaLrTrim, bmax)
	r.deltaRTrim = CalculateDeltaR(r.deltarTrim, r.mOfTrim)
	r.lpodg = CalculateLpodg(r.ltune, r.mOfTrim)
	r.lsum = CalculateLsum(r.lnTrim, r.lpodg, r.mOfTrim)

	// NewWay
	gammakf := CalculateGammaKf(r.height, r.width, r.environment.GetDeltal(), r.environment.GetDeltab())
	n := CalculateNTrim(gammakf, r.environment.GetGammaRokv(), r.tolerance, r.material.senescence)
	if r.resistance == 130000.0 {
		fmt.Println("bmax: ", bmax)
		fmt.Println("bmin: ", bmin)
		fmt.Println("lmin: ", lmin)
		fmt.Println("rokvmax: ", rokvmax)
		fmt.Println("rokvmin: ", rokvmin)
		fmt.Println("rmax: ", rmax)
		fmt.Println("rmin: ", rmin)
		fmt.Println(gammakf, n)
	}
}

func CalculateGammaR(mat material, gammaRdeltaTrim float64, env environment.Environment, gammaRt float64) float64 {
	return env.GetGammaRokv() + gammaRdeltaTrim + gammaRt + mat.GetSenescence() + env.GetGammaRcontact()
}

func CalculateGammaRdeltaTrim(formFactor float64, b float64, deltab float64) float64 {
	return 2 * ((1 + (1 / formFactor)) / ((b / deltab) - (deltab / b))) * 100
}

func CalculateMofTrim(gammaR float64, gammaRdeltaTrim float64) float64 {
	return math.Ceil(gammaR / gammaRdeltaTrim)
}

func CalculateLn(rmax float64, bmin float64, rokvmax float64) float64 {
	return ((rmax * bmin) / rokvmax) - 0.01
}

func CalculateRdashmin(rokvmin float64, lmin float64, bmax float64) float64 {
	return (rokvmin * lmin) / bmax
}

func CalculateLo(rmin float64, bmax float64, rokvmin float64) float64 {
	return (rmin * bmax) / rokvmin
}

func CalculateLtune(lo float64, ln float64) float64 {
	return lo - ln
}

func CalculateDeltarTrim(rmin float64, rdashmin float64) float64 {
	return rmin - rdashmin
}

func CalculateDeltaLr(ltune float64, m float64) float64 {
	return ltune / m
}

func CalculateDeltaLdash(rokvmin float64, deltaLr float64, bmax float64) float64 {
	return (rokvmin * deltaLr) / bmax
}

func CalculateDeltaR(deltaR float64, m float64) float64 {
	return deltaR / m
}

func CalculateLpodg(ltune float64, m float64) float64 {
	return ltune / m
}

func CalculateLsum(ln float64, lpodg float64, m float64) float64 {
	return ln + lpodg*m
}

func CalculateGammaKf(l float64, b float64, deltal float64, deltab float64) float64 {
	return (deltal/l + deltab/b) * 100
}

func CalculateNTrim(gammakf float64, gammarokv float64, gammaR float64, gammaSt float64) float64 {
	return math.Round((gammakf + gammarokv) / (gammaR - gammaSt))
}

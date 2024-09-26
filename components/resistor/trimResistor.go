package resistor

import (
	"example/main/environment"
	"math"
)

type TrimLength struct {
	gammaRdeltaTrim float64
	gammaR          float64
	mOfTrim         float64
	lnTrim          float64
	loTrim          float64
	rdashminTrim    float64
	ltune           float64
	deltarTrim      float64
	deltaLrTrim     float64
	deltaLdashTrim  float64
	deltaRTrim      float64
	lpodg           float64
	lsum            float64
	gammakf         float64
}

func (r *Resistor) initTrimLength() {
	r.trimLength.gammaRdeltaTrim = CalculateGammaRdeltaTrim(r.formFactor, r.GetWidth(), r.environment.GetDeltab())
	r.trimLength.gammaR = CalculateGammaR(r.material, r.GetGammaRdeltaTrimL(), *r.environment, r.gammaRt)
	r.trimLength.mOfTrim = CalculateMofTrim(r.GetGammaRTrimL(), r.GetGammaRdeltaTrimL())
	r.trimLength.lnTrim = CalculateLn(r.Getrmax(), r.Getbmin(), r.Getrokvmax())
	r.trimLength.rdashminTrim = CalculateRdashmin(r.Getrokvmin(), r.GetlnTrimL(), r.Getbmax())
	r.trimLength.deltarTrim = CalculateDeltarTrim(r.Getrmin(), r.GetRdashminTrimL())
	// BackThen
	r.trimLength.loTrim = CalculateLo(r.Getrmin(), r.Getbmax(), r.Getrokvmin())
	r.trimLength.ltune = CalculateLtune(r.GetloTrimL(), r.GetlnTrimL())
	r.trimLength.deltaLrTrim = CalculateDeltaLr(r.GetLtuneTrimL(), r.GetNumberOfLinks())
	r.trimLength.deltaLdashTrim = CalculateDeltaLdash(r.Getrokvmin(), r.GetDeltaLrTrimL(), r.Getbmax())
	r.trimLength.deltaRTrim = CalculateDeltaR(r.GetDeltaRTrimL(), r.GetMOfTrimL())
	r.trimLength.lpodg = CalculateLpodg(r.GetLtuneTrimL(), r.GetMOfTrimL())
	r.trimLength.lsum = CalculateLsum(r.GetlnTrimL(), r.GetLpodgTrimL(), r.GetMOfTrimL())

	// NewWay
	r.trimLength.gammakf = CalculateGammaKf(r.GetHeight(), r.GetWidth(), r.environment.GetDeltal(), r.environment.GetDeltab())
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

// Getters

func (r *Resistor) GetGammaRdeltaTrimL() float64 {
	return r.trimLength.gammaRdeltaTrim
}

func (r *Resistor) GetGammaRTrimL() float64 {
	return r.trimLength.gammaR
}

func (r *Resistor) GetMOfTrimL() float64 {
	return r.trimLength.mOfTrim
}

func (r *Resistor) GetlnTrimL() float64 {
	return r.trimLength.lnTrim
}

func (r *Resistor) GetloTrimL() float64 {
	return r.trimLength.loTrim
}

func (r *Resistor) GetRdashminTrimL() float64 {
	return r.trimLength.rdashminTrim
}

func (r *Resistor) GetLtuneTrimL() float64 {
	return r.trimLength.ltune
}
func (r *Resistor) GetDeltarTrimL() float64 {
	return r.trimLength.deltarTrim
}
func (r *Resistor) GetDeltaLrTrimL() float64 {
	return r.trimLength.deltaLrTrim
}

func (r *Resistor) GetDeltaLdashTrimL() float64 {
	return r.trimLength.deltaLdashTrim
}
func (r *Resistor) GetDeltaRTrimL() float64 {
	return r.trimLength.deltaRTrim
}

func (r *Resistor) GetLpodgTrimL() float64 {
	return r.trimLength.lpodg
}
func (r *Resistor) GetLsumTrimL() float64 {
	return r.trimLength.lsum
}
func (r *Resistor) GetGammakfTrimL() float64 {
	return r.trimLength.gammakf
}

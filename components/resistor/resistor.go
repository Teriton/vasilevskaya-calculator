package resistor

import (
	"example/main/environment"
	"math"
)

type Resistor struct {
	resistance  float64
	tolerance   float64
	power       float64
	material    material
	environment *environment.Environment
	btehn       float64
	ltehn       float64

	// Auto Calculated
	formFactor     float64
	gammaRt        float64
	gammaRdelta    float64
	formOfResistor Form

	// Rectangle
	bp     float64
	bdelta float64
	width  float64
	lp     float64
	ldelta float64
	height float64

	// Stripes

	stripesXlength float64

	// Meander
	numberOfLinks  float64
	meanderXLength float64
	meanderYLength float64

	// CCP
	bpCCP     float64
	bdeltaCCP float64

	// Trim
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

func NewResistor(resistance float64, tolerance float64, power float64, material material, enviroment *environment.Environment) *Resistor {
	resistor := new(Resistor)
	resistor.resistance = resistance
	resistor.tolerance = tolerance
	resistor.power = power
	resistor.material = material
	resistor.environment = enviroment

	resistor.btehn = 0.1
	resistor.ltehn = 0.3

	// Auto Calculated
	resistor.autoCalculateInit()
	return resistor
}

// Getters

func (r *Resistor) GetFromFactor() float64 {
	return r.formFactor
}

func (r *Resistor) GetResistance() float64 {
	return r.resistance
}

func (r *Resistor) GetGammaRt() float64 {
	return r.gammaRt
}

func (r *Resistor) GetTolerance() float64 {
	return r.tolerance
}

func (r *Resistor) GetPower() float64 {
	return r.power
}

func (r *Resistor) GetFormOfResistor() Form {
	return r.formOfResistor
}

func (r *Resistor) GetGammaRdelta() float64 {
	return r.gammaRdelta
}

func (r *Resistor) GetMaterial() *material {
	return &r.material
}

func (r *Resistor) GetBp() float64 {
	return r.bp
}

func (r *Resistor) GetBdelta() float64 {
	return r.bdelta
}

func (r *Resistor) GetWidth() float64 {
	return r.width
}

func (r *Resistor) GetLp() float64 {
	return r.lp
}

func (r *Resistor) GetLdelta() float64 {
	return r.ldelta
}

func (r *Resistor) GetHeight() float64 {
	return r.height
}

func (r *Resistor) GetNumberOfLinks() float64 {
	return r.numberOfLinks
}

func (r *Resistor) GetXlengthMeander() float64 {
	return r.meanderXLength
}

func (r *Resistor) GetYlengthMeander() float64 {
	return r.meanderYLength
}

func (r *Resistor) GetXlengthStripes() float64 {
	return r.stripesXlength
}

func (r *Resistor) GetBpCCP() float64 {
	return r.bpCCP
}

func (r *Resistor) GetBdeltaCCP() float64 {
	return r.bdeltaCCP
}

func (r *Resistor) GetGammaRdeltaTrim() float64 {
	return r.gammaRdeltaTrim
}

func (r *Resistor) GetGammaR() float64 {
	return r.gammaR
}

func (r *Resistor) GetMOfTrim() float64 {
	return r.mOfTrim
}

func (r *Resistor) GetlnTrim() float64 {
	return r.lnTrim
}

func (r *Resistor) GetloTrim() float64 {
	return r.loTrim
}

func (r *Resistor) GetRdashminTrim() float64 {
	return r.rdashminTrim
}

func (r *Resistor) GetLtune() float64 {
	return r.ltune
}
func (r *Resistor) GetDeltarTrim() float64 {
	return r.deltarTrim
}
func (r *Resistor) GetDeltaLrTrim() float64 {
	return r.deltaLrTrim
}

func (r *Resistor) GetDeltaLdashTrim() float64 {
	return r.deltaLdashTrim
}
func (r *Resistor) GetDeltaRTrim() float64 {
	return r.deltaRTrim
}

func (r *Resistor) GetLpodg() float64 {
	return r.lpodg
}
func (r *Resistor) GetLsum() float64 {
	return r.lsum
}
func (r *Resistor) GetGammakf() float64 {
	return r.gammakf
}

// Setters

func (r *Resistor) SetMaterial(material material) {
	r.material = material
	r.autoCalculateInit()
}

// Calculations

func (r *Resistor) autoCalculateInit() {
	r.formFactor = CountFormFactor(r.resistance, r.material.squareResistance)
	r.gammaRt = CountGammaRt(r.material.temperatureCoefficientOfResistance, r.environment.GetTemperature())
	r.gammaRdelta = CountGammaRDelta(r.tolerance, r.material.squareResistance, r.environment.GetGammaRcontact(), r.material.senescence, r.gammaRt)
	r.formOfResistor = CountFormOfResistor(r.formFactor)
	// Rectangle
	r.bp = r.bpCount()
	r.bdelta = r.bdeltaCount()
	r.width = r.bCount()
	r.lp = r.lpCount()
	r.ldelta = r.ldeltaCount()
	r.height = r.lCount()
	// Stripes
	a, b := r.GetWidth(), r.GetWidth()
	r.stripesXlength = CountXstripes(a, b, r.GetFromFactor())
	// Meander
	r.InitMeander(a, b)
	// CCP
	r.bpCCP = CountBpCCP(r.resistance, r.power, r.material)
	r.bdeltaCCP = CountBdeltaCCP(0.01, 0.01, 1, r.tolerance, *r.environment)
	// Trim
	r.InitTrim(0.01)
}

func CountFormFactor(resistance float64, squareResistance float64) float64 {
	return resistance / squareResistance
}

func CountGammaRDelta(tolerance float64, gammaRokv float64, gammaRcontact float64, senescence float64, gammaRt float64) float64 {
	return tolerance - (gammaRokv + gammaRcontact + senescence + gammaRt)
}

func CountGammaRt(tkr float64, temperature float64) float64 {
	return math.Abs((tkr * math.Pow(10, -4) * (temperature - 20)) * 100)
}

func CountFormOfResistor(formFactor float64) Form {
	switch {
	case formFactor > 0.1 && formFactor < 10:
		return Rectangle
	case formFactor <= 0.1:
		return CCP
	case formFactor >= 10:
		return Meander
	default:
		return "undefined"
	}
}

func SetMaterialsForResistors(arrOfRes []Resistor, mat material) {
	for i := range arrOfRes {
		(arrOfRes)[i].SetMaterial(mat)
	}
}

func CalculateAndSetMaterial(arrOfRes []Resistor) {
	RoOpt := CalculateRoOpt(arrOfRes)
	mat := CalculateMaterial(RoOpt)
	SetMaterialsForResistors(arrOfRes, mat)

}

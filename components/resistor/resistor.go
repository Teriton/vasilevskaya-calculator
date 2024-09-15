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

func (r *Resistor) GetFormOfResistor() Form {
	return r.formOfResistor
}

func (r *Resistor) GetGammaRdelta() float64 {
	return r.gammaRdelta
}

func (r *Resistor) GetMaterial() string {
	return r.material.name
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

// Setters

func (r *Resistor) SetMaterial(material material) {
	r.material = material
	r.autoCalculateInit()
}

// Calculations

func (r *Resistor) autoCalculateInit() {
	r.formFactor = r.countFormFactor()
	r.gammaRt = r.countGammaRt()
	r.gammaRdelta = r.countGammaRDelta()
	r.formOfResistor = r.countFormOfResistor()
	// Rectangle
	r.bp = r.bpCount()
	r.bdelta = r.bdeltaCount()
	r.width = r.bCount()
	r.lp = r.lpCount()
	r.ldelta = r.ldeltaCount()
	r.height = r.lCount()
	// Stripes
	a, b := r.GetWidth()*2, r.GetWidth()
	r.stripesXlength = CountXstripes(a, b, r.GetFromFactor())
	// Meander
	r.numberOfLinks = CountNumberOfLinks(a, b, r.GetHeight())
	r.meanderXLength = CountXlengthMeander(a, b, r.numberOfLinks)
	r.meanderYLength = CountYlengthMeander(a, r.GetHeight(), r.GetNumberOfLinks())
	// CCP
	r.bpCCP = CountBpCCP(r.resistance, r.power, r.material)
	r.bdeltaCCP = CountBdeltaCCP(0.01, 0.01, 1, r.tolerance, *r.environment)
}

func (r *Resistor) countFormFactor() float64 {
	return r.resistance / r.material.squareResistance
}

func (r *Resistor) countGammaRDelta() float64 {
	return r.tolerance - (r.environment.GetGammaRokv() + r.environment.GetGammaRcontact() + r.material.senescence + r.gammaRt)
}

func (r *Resistor) countGammaRt() float64 {
	return math.Abs((r.material.temperatureCoefficientOfResistance * math.Pow(10, -4) * (r.environment.GetTemperature() - 20)) * 100)
}

func (r *Resistor) countFormOfResistor() Form {
	switch {
	case r.formFactor > 0.1 && r.formFactor < 10:
		return Rectangle
	case r.formFactor <= 0.1:
		return CCP
	case r.formFactor >= 10:
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

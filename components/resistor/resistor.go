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

	//Tolerance
	precision Precision
	// Rectangle
	rectangle Rectangle

	// Stripes

	stripesXlength float64

	// Meander
	meander Meander

	// CCP
	ccp CCP

	// Trim
	trimLength TrimLength
	trimWide   TrimWide
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

func (r *Resistor) GetEnvironment() *environment.Environment {
	return r.environment
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
func (r *Resistor) GetXlengthStripes() float64 {
	return r.stripesXlength
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
	r.gammaRdelta = CountGammaRDelta(r.tolerance, r.environment.GetGammaRokv(), r.environment.GetGammaRcontact(), r.material.senescence, r.gammaRt)
	r.formOfResistor = CountFormOfResistor(r.formFactor)
	// Rectangle
	r.initRectange()
	// Precision
	r.precisionInit()
	// Stripes
	a, b := r.GetWidth(), r.GetWidth()
	r.stripesXlength = CountXstripes(a, b, r.GetFromFactor())
	// Meander
	r.InitMeander(a, b)
	// CCP
	r.initCCP()
	// Trim
	r.initTrimLength()
	r.InitTrimWide()
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
		return RectangleForm
	case formFactor <= 0.1:
		return CCPForm
	case formFactor >= 10:
		return MeanderForm
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

func TehnRound(num float64) float64 {
	return math.Ceil(num*10) / 10
}

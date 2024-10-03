package environment

type Environment struct {
	temperature   float64
	gammaRokv     float64
	gammaRcontact float64
	deltab        float64
	deltal        float64

	kz       float64
	gammaC0  float64
	gammaCst float64
	delataA  float64
}

func InitEnvironment(temperature float64) *Environment {
	environment := new(Environment)
	environment.temperature = temperature
	environment.gammaRokv = 4.0
	environment.gammaRcontact = 2.0
	environment.deltab = 0.01
	environment.deltal = 0.01

	environment.kz = 3
	environment.gammaC0 = 10
	environment.gammaCst = 3
	environment.delataA = 0.001
	return environment
}

func (e *Environment) GetTemperature() float64 {
	return e.temperature
}

func (e *Environment) GetGammaRokv() float64 {
	return e.gammaRokv
}

func (e *Environment) GetGammaRcontact() float64 {
	return e.gammaRcontact
}

func (e *Environment) GetDeltab() float64 {
	return e.deltab
}

func (e *Environment) GetDeltal() float64 {
	return e.deltal
}

func (e *Environment) GetKz() float64 {
	return e.kz
}

func (e *Environment) GetGammaC0() float64 {
	return e.gammaC0
}

func (e *Environment) GetGammaCst() float64 {
	return e.gammaCst
}

func (e *Environment) GetDaltaA() float64 {
	return e.delataA
}

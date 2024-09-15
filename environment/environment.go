package environment

type Environment struct {
	temperature   float64
	gammaRokv     float64
	gammaRcontact float64
	deltab        float64
	deltal        float64
}

func InitEnvironment(temperature float64) *Environment {
	environment := new(Environment)
	environment.temperature = temperature
	environment.gammaRokv = 4.0
	environment.gammaRcontact = 2.0
	environment.deltab = 0.01
	environment.deltal = 0.01
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

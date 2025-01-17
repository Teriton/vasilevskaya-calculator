package components

type Environment struct {
	//Global
	Temperature Parameter

	//Resistor related
	GammaRokv     Parameter
	GammaRcontact Parameter
	Deltab        Parameter
	Deltal        Parameter
	Btehn         Parameter
	Ltehn         Parameter
}

package calculator

type Calculator struct {
	ResistorController
	EnvironmentController
}

func InitCalculator(resistorController ResistorController, environmentController EnvironmentController) Calculator {
	calc := Calculator{}
	calc.ResistorController = resistorController
	calc.EnvironmentController = environmentController
	return calc
}

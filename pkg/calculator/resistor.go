package calculator

import "github.com/shpakunya/pkg/components"

type ResistorController interface {
	NewResistor(resistance float64, tolerance float64, power float64, material components.ResistorMaterial, env components.Environment) components.Resistor
	GetMaterials() []components.ResistorMaterial
	GenerateMaterial(SquareResistance, PermissibleSpecificPowerDissipation, TemperatureCoefficientOfResistance, Senescence float64) components.ResistorMaterial
}

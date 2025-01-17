package resistorcontrollerIml

import "github.com/shpakunya/pkg/components"

// TODO Узнать названия и вписать
func precisionInit(resistor *components.Resistor) {
	resistor.Precision.Bmax = components.Parameter{Name: "TODO", Symbol: "bmax",
		Unit: "мм", Value: resistor.Rectangle.Width.Value + resistor.Env.Deltab.Value}
	resistor.Precision.Bmin = components.Parameter{Name: "TODO", Symbol: "bmin",
		Unit: "мм", Value: resistor.Rectangle.Width.Value - resistor.Env.Deltab.Value}
	resistor.Precision.Lmax = components.Parameter{Name: "TODO", Symbol: "lmax",
		Unit: "мм", Value: resistor.Rectangle.Height.Value + resistor.Env.Deltal.Value}
	resistor.Precision.Lmax = components.Parameter{Name: "TODO", Symbol: "lmin",
		Unit: "мм", Value: resistor.Rectangle.Height.Value - resistor.Env.Deltal.Value}
	resistor.Precision.Rokvmax = components.Parameter{Name: "TODO", Symbol: "ρквmax",
		Unit: "Ом/кв", Value: resistor.Env.GammaRokv.Value*0.01*resistor.Material.SquareResistance.Value + resistor.Material.SquareResistance.Value}
	resistor.Precision.Rokvmin = components.Parameter{Name: "TODO", Symbol: "ρквmin",
		Unit: "Ом/кв", Value: resistor.Material.SquareResistance.Value - resistor.Env.GammaRokv.Value*0.01*resistor.Material.SquareResistance.Value}
	resistor.Precision.Rmax = components.Parameter{Name: "TODO", Symbol: "Rmax",
		Unit: "Ом", Value: resistor.Resistance.Value * (1 + (resistor.Tolerance.Value-resistor.GammaRdelta.Value-resistor.Material.Senescence.Value)/100)}
	resistor.Precision.Rmax = components.Parameter{Name: "TODO", Symbol: "Rmin",
		Unit: "Ом", Value: resistor.Resistance.Value * (1 - (resistor.Tolerance.Value-resistor.GammaRdelta.Value-resistor.Material.Senescence.Value)/100)}
}

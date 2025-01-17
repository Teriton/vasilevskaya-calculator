package resistorcontrollerIml

import "github.com/shpakunya/pkg/components"

func precisionInit(resistor *components.Resistor) {
	resistor.Precision.Bmax = components.Parameter{Name: "Максимальное значение ширины", Symbol: "bmax",
		Unit: "мм", Value: resistor.Rectangle.Width.Value + resistor.Env.Deltab.Value}
	resistor.Precision.Bmin = components.Parameter{Name: "Минимальное значение ширины", Symbol: "bmin",
		Unit: "мм", Value: resistor.Rectangle.Width.Value - resistor.Env.Deltab.Value}
	resistor.Precision.Lmax = components.Parameter{Name: "Максимальное значение ширины", Symbol: "lmax",
		Unit: "мм", Value: resistor.Rectangle.Height.Value + resistor.Env.Deltal.Value}
	resistor.Precision.Lmin = components.Parameter{Name: "Минимальное значение ширины", Symbol: "lmin",
		Unit: "мм", Value: resistor.Rectangle.Height.Value - resistor.Env.Deltal.Value}
	resistor.Precision.Rokvmax = components.Parameter{Name: "Максимальное значение сопротивления квадрата резистивной пленки", Symbol: "ρквmax",
		Unit: "Ом/кв", Value: resistor.Env.GammaRokv.Value*0.01*resistor.Material.SquareResistance.Value + resistor.Material.SquareResistance.Value}
	resistor.Precision.Rokvmin = components.Parameter{Name: "Минимальное значение сопротивления квадрата резистивной пленки", Symbol: "ρквmin",
		Unit: "Ом/кв", Value: resistor.Material.SquareResistance.Value - resistor.Env.GammaRokv.Value*0.01*resistor.Material.SquareResistance.Value}
	resistor.Precision.Rmax = components.Parameter{Name: "Максимальное значение сопротивления", Symbol: "Rmax",
		Unit: "Ом", Value: resistor.Resistance.Value * (1 + (resistor.Tolerance.Value-resistor.GammaRdelta.Value-resistor.Material.Senescence.Value)/100)}
	resistor.Precision.Rmin = components.Parameter{Name: "Минимальное значение сопротивления", Symbol: "Rmin",
		Unit: "Ом", Value: resistor.Resistance.Value * (1 - (resistor.Tolerance.Value-resistor.GammaRdelta.Value-resistor.Material.Senescence.Value)/100)}
}

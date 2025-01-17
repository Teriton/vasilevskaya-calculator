package environmentControllerIml

import "github.com/shpakunya/pkg/components"

type EnvironmentController struct {
}

func (EnvironmentController) NewEnvironment(temperature float64) components.Environment {
	environment := components.Environment{}
	environment.Temperature = components.Parameter{Name: "Температура", Symbol: "T",
		Unit: "°C", Value: temperature}
	ressistorEnvInit(&environment)
	return environment
}

func ressistorEnvInit(environment *components.Environment) {
	environment.GammaRokv = components.Parameter{Name: "Погрешность воспроизведения сопротивления квадрата резис­ тивной пленки", Symbol: "γρкв",
		Unit: "%", Value: 4.0}
	environment.GammaRcontact = components.Parameter{Name: "Погрешность сопротивления, вносимая контактами", Symbol: "γRк",
		Unit: "%", Value: 2.0}
	environment.Deltab = components.Parameter{Name: "Точность воспроизведения геометрнн резисторов, ширины", Symbol: "Δb",
		Unit: "мм", Value: 0.01}
	environment.Deltal = components.Parameter{Name: "Точность воспроизведения геометрнн резисторов, длинны", Symbol: "Δl",
		Unit: "мм", Value: 0.01}
	environment.Btehn = components.Parameter{Name: "Минимальная ширина резистора", Symbol: "bтехн",
		Unit: "мм", Value: 0.1}
	environment.Ltehn = components.Parameter{Name: "Минимальная длинна резистора", Symbol: "lтехн",
		Unit: "мм", Value: 0.3}
}

package components

type ParameterMain[T any] struct {
	Name   string
	Symbol string
	Unit   string
	Value  T
}

type Parameter ParameterMain[float64]

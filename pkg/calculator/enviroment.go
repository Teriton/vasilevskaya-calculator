package calculator

import "github.com/shpakunya/pkg/components"

type EnvironmentController interface {
	NewEnvironment(temperature float64) components.Environment
}

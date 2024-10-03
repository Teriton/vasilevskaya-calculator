package capacitor

import "example/main/environment"

type Capacitor struct {
	capacity  float64
	urab      float64
	tolerance float64

	env      *environment.Environment
	material Material

	areaMoreThan5 AreaMoreThan5
}

func NewCapacitor(capacity float64, urab float64, tolerance float64, material Material, env *environment.Environment) *Capacitor {
	cap := new(Capacitor)

	cap.capacity = capacity
	cap.urab = urab
	cap.tolerance = tolerance

	cap.env = env
	cap.material = material

	cap.autoCalculateInit()

	return cap
}

func (c *Capacitor) autoCalculateInit() {
	c.initAreaMoreThan5()
}

// Setters

func (c *Capacitor) SetMaterial(material Material) {
	c.material = material
	c.autoCalculateInit()
}

// Getters

func (c *Capacitor) GetCapacity() float64 {
	return c.capacity
}

func (c *Capacitor) GetUrab() float64 {
	return c.urab
}

func (c *Capacitor) GetTolerance() float64 {
	return c.tolerance
}

func (c *Capacitor) GetMaterial() *Material {
	return &c.material
}
func (c *Capacitor) GetAreaMoreThan5() *AreaMoreThan5 {
	return &c.areaMoreThan5
}

func (c *Capacitor) GetEnv() *environment.Environment {
	return c.env
}

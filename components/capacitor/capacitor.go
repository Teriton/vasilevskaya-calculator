package capacitor

import (
	"example/main/environment"
	"math"
)

type Capacitor struct {
	capacity  float64
	urab      float64
	tolerance float64

	env      *environment.Environment
	material Material

	areaMoreThan5 AreaMoreThan5
	area15        Area15
	ctripledash0  float64
}

func NewCapacitor(capacity float64, urab float64, tolerance float64, material Material, env *environment.Environment) *Capacitor {
	cap := new(Capacitor)

	cap.capacity = capacity
	cap.urab = urab
	cap.tolerance = tolerance

	cap.env = env
	cap.material = material
	cap.ctripledash0 = 1000000000000000000

	cap.autoCalculateInit()
	return cap
}

func (c *Capacitor) autoCalculateInit() {
	c.initAreaMoreThan5()
	c.area15Init()
}

func TehnRound(num float64) float64 {
	return math.Ceil(num*10) / 10
}

func CalculateCtripledash0(arrOfCaps []Capacitor) float64 {
	var smallestCapacitor Capacitor = arrOfCaps[0]
	for i, j := range arrOfCaps {
		if j.GetCapacity() < smallestCapacitor.GetCapacity() {
			smallestCapacitor = arrOfCaps[i]
		}
	}
	return smallestCapacitor.GetCapacity() / 0.01
}

func SetCtripledash0ForCapacitors(arrOfCaps []Capacitor, num float64) {
	for i := range arrOfCaps {
		(arrOfCaps)[i].SetCtripledash0(num)
	}
}

// Setters

func (c *Capacitor) SetMaterial(material Material) {
	c.material = material
	c.autoCalculateInit()
}

func (c *Capacitor) SetCtripledash0(num float64) {
	c.ctripledash0 = num
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
func (c *Capacitor) GetCtripledash0() float64 {
	return c.ctripledash0
}

func (c *Capacitor) GetArea15() *Area15 {
	return &c.area15
}

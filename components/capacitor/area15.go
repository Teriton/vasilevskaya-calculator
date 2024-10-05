package capacitor

// НЕ РАБОЧАИЙ КОД

type Area15 struct {
	area float64
	k    float64
}

func (c *Capacitor) area15Init() {
	c.area15.k = CalculateK(c.GetAreaMoreThan5().GetArea())
	c.area15.area = CalculateArea15(c.GetCapacity(), c.GetArea15().GetK(), c.GetAreaMoreThan5().GetRealD())
}

func CalculateK(s float64) float64 {
	return 1.3785 - s*(0.175/2)
}

func CalculateArea15(c, k, d float64) float64 {
	return c / (0.0885 * k * (d / 1000))
}

// Getters

func (c *Area15) GetArea() float64 {
	return c.area
}

func (c *Area15) GetK() float64 {
	return c.k
}

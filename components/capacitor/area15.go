package capacitor

// НЕ РАБОЧАИЙ КОД

type Area15 struct {
	area float64
	k    float64
}

func (c *Capacitor) area15Init() {
	c.area15.k = CalculateK(c.GetAreaMoreThan5().GetArea())
	c.area15.area = CalculateArea15(c.GetCapacity(), c.GetArea15().GetK(), c.GetAreaMoreThan5().GetRealD(), c.GetMaterial().Gete())
}

func CalculateK(s float64) float64 {
	return (199.0 / 150.0) - s*(1.0/15.0)
}

func CalculateArea15(c, k, d, e float64) float64 {
	return (c * k * (d / 1000)) / (0.008854 * e)
}

// Getters

func (c *Area15) GetArea() float64 {
	return c.area
}

func (c *Area15) GetK() float64 {
	return c.k
}

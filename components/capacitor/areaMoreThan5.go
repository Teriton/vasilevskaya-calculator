package capacitor

import "math"

type AreaMoreThan5 struct {
	d            float64
	gammaCt      float64
	gammaSdop    float64
	cdash0       float64
	cdoubledash0 float64
	c0           float64
	area         float64
}

func (c *Capacitor) initAreaMoreThan5() {
	c.areaMoreThan5.d = CalculateD(c.urab, c.env.GetKz(), c.GetMaterial().GetElStrength())
	c.areaMoreThan5.gammaCt = CalculateGammaCt(c.GetEnv().GetTemperature(), c.GetMaterial().GetTKE())
	c.areaMoreThan5.gammaSdop = CalculateGammaSdop(
		c.GetTolerance(),
		c.GetEnv().GetGammaC0(),
		c.GetAreaMoreThan5().GetgammaCt(),
		c.GetEnv().GetGammaCst(),
	)
	c.areaMoreThan5.cdash0 = CalculateCdash0(c.GetMaterial().Gete(), c.GetAreaMoreThan5().Getd())
	c.areaMoreThan5.cdoubledash0 = CalculateCdoubledash0(
		c.GetCapacity(),
		c.GetAreaMoreThan5().GetGammaSdop(),
		c.GetEnv().GetDaltaA(),
	)
	c.areaMoreThan5.c0 = CalculateC0(c.GetAreaMoreThan5().GetCdash0(), c.GetAreaMoreThan5().GetDdoubledash0())
	c.areaMoreThan5.area = CalculateArea(c.GetAreaMoreThan5().GetC0(), c.GetCapacity())
}

func CalculateD(urab, kz, Epr float64) float64 {
	return 10 * ((urab * kz) / (Epr * 100))
}

func CalculateGammaSdop(gammaC, gammaC0, gammaCt, gammaCst float64) float64 {
	return gammaC - gammaC0 - gammaCt - gammaCst
}

func CalculateGammaCt(t, TKE float64) float64 {
	return (TKE * math.Pow(10, -4) * (t - 20)) * 100
}

func CalculateCdash0(e, d float64) float64 {
	return 0.0885 * (e / (d * math.Pow(10, -2)))
}

func CalculateCdoubledash0(c, gammaSdop, deltaA float64) float64 { // пф/мм^2
	return c * math.Pow((gammaSdop/100)/(2*deltaA*10), 2)
}

func CalculateC0(cdash0, cdoubledash0 float64) float64 { // пф/мм^2
	return math.Min(cdash0, cdoubledash0)
}

func CalculateArea(c0, capacity float64) float64 {
	return capacity / c0
}

// Getters

func (c *AreaMoreThan5) Getd() float64 {
	return c.d
}

func (c *AreaMoreThan5) GetgammaCt() float64 {
	return c.gammaCt
}

func (c *AreaMoreThan5) GetGammaSdop() float64 {
	return c.gammaSdop
}

func (c *AreaMoreThan5) GetCdash0() float64 {
	return c.cdash0
}

func (c *AreaMoreThan5) GetDdoubledash0() float64 {
	return c.cdoubledash0
}

func (c *AreaMoreThan5) GetC0() float64 {
	return c.c0
}

func (c *AreaMoreThan5) GetArea() float64 {
	return c.area
}

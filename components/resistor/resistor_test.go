package resistor

import (
	"example/main/environment"
	"testing"
)

type formFactorTest struct {
	res, reskv, ff float64
}

var ffTests = []formFactorTest{
	{1, 1, 1},
	{1000, 500, 2},
	{5000, 10000, 0.5},
}

func TestCountFormFactor(t *testing.T) {
	for _, test := range ffTests {
		if output := CountFormFactor(test.res, test.reskv); output != test.ff {
			t.Errorf("Form factor  %f, but expected  %f", output, test.ff)
		}
	}
}

type gammaRtTest struct {
	tkr, temperature, gammaRt float64
}

var gammaRtTests = []gammaRtTest{
	{15, 70, 7.5},
}

func TestCountGammaRt(t *testing.T) {
	for _, test := range gammaRtTests {
		if output := CountGammaRt(test.tkr, test.temperature); output != test.gammaRt {
			t.Errorf("GammaRt  %f, but expected  %f", output, test.gammaRt)
		}
	}
}

type gammaRdeltaTest struct {
	tolerance, gammaRokv, gammaRcontact, senescence, gammaRt, output float64
}

var gammaRdeltaTests = []gammaRdeltaTest{
	{20, 4, 2, 0, 0.5, 13.5},
}

func TestCountGammaRDelta(t *testing.T) {
	for _, test := range gammaRdeltaTests {
		if output := CountGammaRDelta(test.tolerance, test.gammaRokv, test.gammaRcontact, test.senescence, test.gammaRt); output != test.output {
			t.Errorf("GammaRdelta  %f, but expected  %f", output, test.output)
		}
	}
}

type formTest struct {
	formFactor float64
	output     Form
}

var formTests = []formTest{
	{100, MeanderForm},
	{1, RectangleForm},
	{0.1, CCPForm},
}

func TestCountFormOfResistor(t *testing.T) {
	for _, test := range formTests {
		if output := CountFormOfResistor(test.formFactor); output != test.output {
			t.Errorf("Form is  %s, but expected  %s", output, test.output)
		}
	}
}

func initResistors() []Resistor {
	env := environment.InitEnvironment(70)
	arrRes := []Resistor{
		*NewResistor(20000.0, 10.0, 0.005, GetHollowMaterial(), env),
		*NewResistor(1000000.0, 20.0, 0.120, GetHollowMaterial(), env),
		*NewResistor(100000.0, 10.0, 0.0015, GetHollowMaterial(), env),
		*NewResistor(2700.0, 20.0, 0.060, GetHollowMaterial(), env),
		*NewResistor(100000.0, 20.0, 0.080, GetHollowMaterial(), env),
		*NewResistor(1000.0, 10.0, 0.150, GetHollowMaterial(), env),
		*NewResistor(16000.0, 10.0, 0.050, GetHollowMaterial(), env),
		*NewResistor(470000.0, 20.0, 0.100, GetHollowMaterial(), env),
		*NewResistor(4700.0, 20.0, 0.055, GetHollowMaterial(), env),
		*NewResistor(470.0, 20.0, 0.010, GetHollowMaterial(), env),
	}
	return arrRes
}

func TestSetMaterialsForResistors(t *testing.T) {
	arrOfRes := initResistors()
	SetMaterialsForResistors(arrOfRes, GetMaterials()[0])
	for _, j := range arrOfRes {
		if *j.GetMaterial() != GetMaterials()[0] {
			t.Errorf("Material is %s, but expected %s", j.material.name, GetMaterials()[0].name)
		}
	}
}

func TestCalculateAndSetMaterial(t *testing.T) {
	arrOfRes := initResistors()
	CalculateAndSetMaterial(arrOfRes)
	for _, j := range arrOfRes {
		if *j.GetMaterial() != GetMaterials()[0] {
			t.Errorf("Material is %s, but expected %s", j.material.name, GetMaterials()[0].name)
		}
	}
}

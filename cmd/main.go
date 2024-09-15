package main

import (
	"example/main/components/resistor"
	consoleoutput "example/main/consoleOutput"
	"example/main/environment"
	"fmt"
)

func main() {
	env := environment.InitEnvironment(70.0)

	// arrRes := []resistor.Resistor{
	// 	// *resistor.NewResistor(20000.0, 10.0, 0.005, resistor.GetHollowMaterial(), env),
	// 	// *resistor.NewResistor(1000000.0, 20.0, 0.120, resistor.GetHollowMaterial(), env),
	// 	// *resistor.NewResistor(100000.0, 10.0, 0.0015, resistor.GetHollowMaterial(), env),
	// 	// *resistor.NewResistor(2700.0, 20.0, 0.060, resistor.GetHollowMaterial(), env),
	// 	// *resistor.NewResistor(100000.0, 20.0, 0.080, resistor.GetHollowMaterial(), env),
	// 	// *resistor.NewResistor(1000.0, 10.0, 0.150, resistor.GetHollowMaterial(), env),
	// 	// *resistor.NewResistor(16000.0, 10.0, 0.050, resistor.GetHollowMaterial(), env),
	// 	// *resistor.NewResistor(470000.0, 20.0, 0.100, resistor.GetHollowMaterial(), env),
	// 	// *resistor.NewResistor(4700.0, 20.0, 0.055, resistor.GetHollowMaterial(), env),
	// 	// *resistor.NewResistor(470.0, 20.0, 0.010, resistor.GetHollowMaterial(), env),
	// }

	arrRes := []resistor.Resistor{
		*resistor.NewResistor(1200.0, 10.0, 0.150, resistor.GetHollowMaterial(), env),
		*resistor.NewResistor(2200.0, 20.0, 0.050, resistor.GetHollowMaterial(), env),
		*resistor.NewResistor(6000.0, 10.0, 0.040, resistor.GetHollowMaterial(), env),
		*resistor.NewResistor(1000.0, 20.0, 0.010, resistor.GetHollowMaterial(), env),
		*resistor.NewResistor(120000.0, 20.0, 0.051, resistor.GetHollowMaterial(), env),
		*resistor.NewResistor(130000.0, 2.0, 0.090, resistor.GetHollowMaterial(), env),
		*resistor.NewResistor(1000000.0, 10.0, 0.012, resistor.GetHollowMaterial(), env),
		*resistor.NewResistor(500000.0, 20.0, 0.032, resistor.GetHollowMaterial(), env),
		*resistor.NewResistor(12000.0, 20.0, 0.045, resistor.GetHollowMaterial(), env),
		*resistor.NewResistor(666000.0, 20.0, 0.123, resistor.GetHollowMaterial(), env),
	}

	resistor.CalculateAndSetMaterial(arrRes)

	fmt.Println(resistor.CalculateRoOpt(arrRes))
	fmt.Println(arrRes[0].GetMaterial())

	// res := *resistor.NewResistor(110000, 5, 100, resistor.GetHollowMaterial(), env)

	// for _, j := range resistor.GetMaterials() {
	// 	res.SetMaterial(j)
	// 	if res.GetGammaRdelta() > 0 {
	// 		fmt.Println("===================================")
	// 		fmt.Println(res.GetMaterial())
	// 		fmt.Println(res.GetGammaRdelta())
	// 		fmt.Println("===================================")
	// 	}

	// }

	for i, j := range arrRes {
		fmt.Println("======================")
		fmt.Println("Номер:", i+1)
		fmt.Println("Сопротивление: ", j.GetResistance())
		fmt.Println("Коэффициент формы: ", j.GetFromFactor())
		fmt.Println("gammaRt: ", j.GetGammaRt())
		fmt.Println("gammaRdelta: ", j.GetGammaRdelta())
		fmt.Println("Форма: ", j.GetFormOfResistor())
		fmt.Println("bp: ", j.GetBp())
		fmt.Println("bdelta: ", j.GetBdelta())
		fmt.Println("Ширина: ", j.GetWidth())
		fmt.Println("lp: ", j.GetLp())
		fmt.Println("ldelta: ", j.GetLdelta())
		fmt.Println("Длинна: ", j.GetHeight())
		fmt.Println("Колличество звеньев мендра: ", j.GetNumberOfLinks())
		fmt.Println("X меандр: ", j.GetXlengthMeander())
		fmt.Println("Y меандр: ", j.GetYlengthMeander())
		fmt.Println("X полоски: ", j.GetXlengthStripes())
		fmt.Println("ЦКП bp: ", j.GetBpCCP())
		fmt.Println("ЦКП bdelta: ", j.GetBdeltaCCP())
		fmt.Println("======================")
	}

	consoleoutput.RenderTableOfResistors(arrRes)

	consoleoutput.RenderMaterialsCheck(resistor.CheckAllMaterials(arrRes))
	for _, mat := range resistor.GetMaterials() {
		resistor.SetMaterialsForResistors(arrRes, mat)
		consoleoutput.RenderTableOfResistors(arrRes)
	}

	//server.RunServer()

}

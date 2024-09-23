package main

import (
	"example/main/components/resistor"
	consoleoutput "example/main/consoleOutput"
	"example/main/environment"
	"example/main/server"
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
		*resistor.NewResistor(1200.0, 20.0, 0.040, resistor.GetHollowMaterial(), env),
		*resistor.NewResistor(2200.0, 20.0, 0.050, resistor.GetHollowMaterial(), env),
		*resistor.NewResistor(6000.0, 10.0, 0.040, resistor.GetHollowMaterial(), env),
		*resistor.NewResistor(1100.0, 20.0, 0.010, resistor.GetHollowMaterial(), env),
		*resistor.NewResistor(120000.0, 20.0, 0.051, resistor.GetHollowMaterial(), env),
		*resistor.NewResistor(130000.0, 20.0, 0.090, resistor.GetHollowMaterial(), env),
		*resistor.NewResistor(100000.0, 10.0, 0.012, resistor.GetHollowMaterial(), env),
		*resistor.NewResistor(500000.0, 20.0, 0.032, resistor.GetHollowMaterial(), env),
		*resistor.NewResistor(12000.0, 2.0, 0.045, resistor.GetHollowMaterial(), env),
		*resistor.NewResistor(600000.0, 20.0, 0.123, resistor.GetHollowMaterial(), env),
	}
	// Main
	// arrRes := []resistor.Resistor{
	// 	*resistor.NewResistor(30000.0, 10.0, 0.015, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(2000.0, 2.0, 0.020, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(10000.0, 20.0, 0.045, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(100000.0, 20.0, 0.040, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(50000.0, 10.0, 0.038, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(80000.0, 20.0, 0.026, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(55000.0, 10.0, 0.018, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(20000.0, 10.0, 0.030, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(35000.0, 20.0, 0.022, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(90000.0, 20.0, 0.024, resistor.GetHollowMaterial(), env),
	// }
	// Полина
	// arrRes := []resistor.Resistor{
	// 	*resistor.NewResistor(20000.0, 10.0, 0.020, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(1100.0, 20.0, 0.015, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(4000.0, 10.0, 0.036, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(110000.0, 20.0, 0.0185, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(1700.0, 20.0, 0.025, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(16000.0, 2.0, 0.020, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(40000.0, 10.0, 0.020, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(8000.0, 20.0, 0.038, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(2000.0, 10.0, 0.040, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(50000.0, 20.0, 0.016, resistor.GetHollowMaterial(), env),
	// }
	// Геля
	// arrRes := []resistor.Resistor{
	// 	*resistor.NewResistor(4700.0, 10.0, 0.010, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(100000.0, 20.0, 0.020, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(20000.0, 10.0, 0.040, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(18000.0, 2.0, 0.030, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(76000.0, 20.0, 0.065, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(30000.0, 20.0, 0.020, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(55000.0, 10.0, 0.070, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(31000.0, 20.0, 0.050, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(65000.0, 10.0, 0.055, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(47000.0, 10.0, 0.060, resistor.GetHollowMaterial(), env),
	// }
	// Лена
	// arrRes := []resistor.Resistor{
	// 	*resistor.NewResistor(20000.0, 10.0, 0.040, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(1000.0, 10.0, 0.015, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(3000.0, 20.0, 0.020, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(100000.0, 10.0, 0.080, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(2000.0, 10.0, 0.050, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(15000.0, 20.0, 0.100, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(30000.0, 2.0, 0.010, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(55000.0, 10.0, 0.090, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(150000.0, 10.0, 0.030, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(40000.0, 20.0, 0.010, resistor.GetHollowMaterial(), env),
	// }
	// Француз
	// arrRes := []resistor.Resistor{
	// 	*resistor.NewResistor(8000.0, 10.0, 0.010, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(3600.0, 2.0, 0.015, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(8100.0, 20.0, 0.075, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(10000.0, 10.0, 0.005, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(10000.0, 10.0, 0.050, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(600000.0, 10.0, 0.120, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(22000.0, 20.0, 0.020, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(14000.0, 20.0, 0.015, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(15000.0, 10.0, 0.006, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(33000.0, 20.0, 0.060, resistor.GetHollowMaterial(), env),
	// }
	// Роман
	// arrRes := []resistor.Resistor{
	// 	*resistor.NewResistor(10000.0, 10.0, 0.005, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(1000000.0, 10.0, 0.120, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(12000.0, 20.0, 0.025, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(36000.0, 10.0, 0.020, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(150000.0, 2.0, 0.050, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(25000.0, 20.0, 0.015, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(730000.0, 10.0, 0.150, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(10000.0, 10.0, 0.030, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(330000.0, 10.0, 0.075, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(14000.0, 20.0, 0.020, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(10000.0, 10.0, 0.015, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(225000.0, 10.0, 0.075, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(15000.0, 10.0, 0.050, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(88000.0, 20.0, 0.055, resistor.GetHollowMaterial(), env),
	// }
	// Дима
	// arrRes := []resistor.Resistor{
	// 	*resistor.NewResistor(1000.0, 10.0, 0.005, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(5000.0, 20.0, 0.010, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(12000.0, 10.0, 0.012, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(20000.0, 20.0, 0.040, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(500.0, 2.0, 0.400, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(1000000.0, 10.0, 0.010, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(40000.0, 20.0, 0.020, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(62000.0, 10.0, 0.010, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(1600000.0, 20.0, 0.150, resistor.GetHollowMaterial(), env),
	// 	*resistor.NewResistor(4000.0, 10.0, 0.009, resistor.GetHollowMaterial(), env),
	// }
	// arrRes := []resistor.Resistor{
	// 	*resistor.NewResistor(130000.0, 2.0, 0.090, resistor.GetHollowMaterial(), env),
	// }

	resistor.CalculateAndSetMaterial(arrRes)

	fmt.Println(resistor.CalculateRoOpt(arrRes))
	// resistor.SetMaterialsForResistors(arrRes, resistor.GetMaterials()[0])
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

	for _, j := range arrRes {
		// fmt.Println("======================")
		// fmt.Println("Номер:", i+1)
		// fmt.Println("Сопротивление: ", j.GetResistance())
		// fmt.Println("Коэффициент формы: ", j.GetFromFactor())
		// fmt.Println("gammaRt: ", j.GetGammaRt())
		// fmt.Println("gammaRdelta: ", j.GetGammaRdelta())
		// fmt.Println("Форма: ", j.GetFormOfResistor())
		// fmt.Println("bp: ", j.GetBp())
		// fmt.Println("bdelta: ", j.GetBdelta())
		// fmt.Println("Ширина: ", j.GetWidth())
		// fmt.Println("lp: ", j.GetLp())
		// fmt.Println("ldelta: ", j.GetLdelta())
		// fmt.Println("Длинна: ", j.GetHeight())
		// fmt.Println("Колличество звеньев мендра: ", j.GetNumberOfLinks())
		// fmt.Println("X меандр: ", j.GetXlengthMeander())
		// fmt.Println("Y меандр: ", j.GetYlengthMeander())
		// fmt.Println("X полоски: ", j.GetXlengthStripes())
		// fmt.Println("ЦКП bp: ", j.GetBpCCP())
		// fmt.Println("ЦКП bdelta: ", j.GetBdeltaCCP())
		// fmt.Println("gammaRdelta подгоночного: ", j.GetGammaRdeltaTrim())
		// fmt.Println("gammaR подгоночного:", j.GetGammaR())
		// fmt.Println("Количество секций подгонки", j.GetMOfTrim())
		// fmt.Println("l нерегулироемое: ", j.GetlnTrim())
		// fmt.Println("l регулируемое: ", j.GetLtune())
		// fmt.Println("l общее: ", j.GetloTrim())
		// fmt.Println("R'min: ", j.GetRdashminTrim())
		// fmt.Println("deltaR: ", j.GetDeltaRTrim())
		// fmt.Println("Deltaldash: ", j.GetDeltaLdashTrim())
		// fmt.Println("DeltaR: ", j.GetDeltaRTrim())
		// fmt.Println("SumL: ", j.GetLsum())
		// fmt.Println("deltaR: ", j.GetDeltaRTrim())
		// fmt.Println("lpodg: ", j.GetLpodg())
		// fmt.Println("deltaLr:", j.GetDeltaLrTrim())
		// fmt.Println("======================")

		if j.GetGammaRdelta() < 0 {
			consoleoutput.PodgResistorPrint(&j)
		}
	}

	// resistor.SetMaterialsForResistors(arrRes, resistor.GetMaterials()[0])
	consoleoutput.RenderEnviroment(arrRes)
	consoleoutput.RenderTableOfResistors(arrRes)
	consoleoutput.RenderTableOfResistorsPractos1(arrRes)

	consoleoutput.RenderMaterialsCheck(resistor.CheckAllMaterials(arrRes))
	for _, mat := range resistor.GetMaterials() {
		resistor.SetMaterialsForResistors(arrRes, mat)
		consoleoutput.RenderTableOfResistors(arrRes)
	}

	server.RunServer()

}

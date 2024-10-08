package main

import (
	"example/main/components/capacitor"
	"example/main/environment"
	"example/main/server"
	"fmt"
)

// func main() {
// 	env := environment.InitEnvironment(70.0)

// 	// arrRes := []resistor.Resistor{
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
// 	//

// 	arrRes := []resistor.Resistor{
// 		*resistor.NewResistor(100.0, 10.0, 0.05, resistor.GetHollowMaterial(), env),
// 		// *resistor.NewResistor(1200.0, 20.0, 0.040, resistor.GetHollowMaterial(), env),
// 		// *resistor.NewResistor(2200.0, 10.0, 0.050, resistor.GetHollowMaterial(), env),
// 		// *resistor.NewResistor(6000.0, 10.0, 0.040, resistor.GetHollowMaterial(), env),
// 		// *resistor.NewResistor(1100.0, 10.0, 0.010, resistor.GetHollowMaterial(), env),
// 		// *resistor.NewResistor(120000.0, 20.0, 0.051, resistor.GetHollowMaterial(), env),
// 		// *resistor.NewResistor(130000.0, 20.0, 0.090, resistor.GetHollowMaterial(), env),
// 		// *resistor.NewResistor(100000.0, 10.0, 0.012, resistor.GetHollowMaterial(), env),
// 		// *resistor.NewResistor(500000.0, 20.0, 0.032, resistor.GetHollowMaterial(), env),
// 		// *resistor.NewResistor(12000.0, 2.0, 0.045, resistor.GetHollowMaterial(), env),
// 		// *resistor.NewResistor(600000.0, 20.0, 0.123, resistor.GetHollowMaterial(), env),
// 	}

// 	// arrRes2 := []resistor.Resistor{
// 	// 	*resistor.NewResistor(1200.0, 20.0, 0.040, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(2200.0, 10.0, 0.050, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(6000.0, 10.0, 0.040, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(1100.0, 10.0, 0.010, resistor.GetHollowMaterial(), env),
// 	// }
// 	// Main
// 	// arrRes := []resistor.Resistor{
// 	// 	*resistor.NewResistor(30000.0, 10.0, 0.015, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(2000.0, 2.0, 0.020, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(10000.0, 20.0, 0.045, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(100000.0, 20.0, 0.040, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(50000.0, 10.0, 0.038, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(80000.0, 20.0, 0.026, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(55000.0, 10.0, 0.018, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(20000.0, 10.0, 0.030, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(35000.0, 20.0, 0.022, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(90000.0, 20.0, 0.024, resistor.GetHollowMaterial(), env),
// 	// }
// 	// Полина
// 	// arrRes := []resistor.Resistor{
// 	// 	*resistor.NewResistor(20000.0, 10.0, 0.020, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(1100.0, 20.0, 0.015, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(4000.0, 10.0, 0.036, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(110000.0, 20.0, 0.0185, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(1700.0, 20.0, 0.025, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(16000.0, 2.0, 0.020, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(40000.0, 10.0, 0.020, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(8000.0, 20.0, 0.038, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(2000.0, 10.0, 0.040, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(50000.0, 20.0, 0.016, resistor.GetHollowMaterial(), env),
// 	// }
// 	// Геля
// 	// arrRes := []resistor.Resistor{
// 	// 	*resistor.NewResistor(4700.0, 10.0, 0.010, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(100000.0, 20.0, 0.020, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(20000.0, 10.0, 0.040, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(18000.0, 2.0, 0.030, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(76000.0, 20.0, 0.065, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(30000.0, 20.0, 0.020, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(55000.0, 10.0, 0.070, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(31000.0, 20.0, 0.050, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(65000.0, 10.0, 0.055, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(47000.0, 10.0, 0.060, resistor.GetHollowMaterial(), env),
// 	// }
// 	// Лена
// 	// arrRes := []resistor.Resistor{
// 	// 	*resistor.NewResistor(20000.0, 10.0, 0.040, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(1000.0, 10.0, 0.015, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(3000.0, 20.0, 0.020, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(100000.0, 10.0, 0.080, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(2000.0, 10.0, 0.050, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(15000.0, 20.0, 0.100, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(30000.0, 2.0, 0.010, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(55000.0, 10.0, 0.090, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(150000.0, 10.0, 0.030, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(40000.0, 20.0, 0.010, resistor.GetHollowMaterial(), env),
// 	// }
// 	// Француз
// 	// arrRes := []resistor.Resistor{
// 	// 	*resistor.NewResistor(8000.0, 10.0, 0.010, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(3600.0, 2.0, 0.015, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(8100.0, 20.0, 0.075, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(10000.0, 10.0, 0.005, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(10000.0, 10.0, 0.050, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(600000.0, 10.0, 0.120, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(22000.0, 20.0, 0.020, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(14000.0, 20.0, 0.015, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(15000.0, 10.0, 0.006, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(33000.0, 20.0, 0.060, resistor.GetHollowMaterial(), env),
// 	// }
// 	// Роман
// 	// arrRes := []resistor.Resistor{
// 	// 	*resistor.NewResistor(10000.0, 10.0, 0.005, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(1000000.0, 10.0, 0.120, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(12000.0, 20.0, 0.025, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(36000.0, 10.0, 0.020, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(150000.0, 2.0, 0.050, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(25000.0, 20.0, 0.015, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(730000.0, 10.0, 0.150, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(10000.0, 10.0, 0.030, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(330000.0, 10.0, 0.075, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(14000.0, 20.0, 0.020, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(10000.0, 10.0, 0.015, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(225000.0, 10.0, 0.075, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(15000.0, 10.0, 0.050, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(88000.0, 20.0, 0.055, resistor.GetHollowMaterial(), env),
// 	// }
// 	// Дима
// 	// arrRes := []resistor.Resistor{
// 	// 	*resistor.NewResistor(1000.0, 10.0, 0.005, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(5000.0, 20.0, 0.010, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(12000.0, 10.0, 0.012, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(20000.0, 20.0, 0.040, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(500.0, 2.0, 0.400, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(1000000.0, 10.0, 0.010, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(40000.0, 20.0, 0.020, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(62000.0, 10.0, 0.010, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(1600000.0, 20.0, 0.150, resistor.GetHollowMaterial(), env),
// 	// 	*resistor.NewResistor(4000.0, 10.0, 0.009, resistor.GetHollowMaterial(), env),
// 	// }
// 	// arrRes := []resistor.Resistor{
// 	// 	*resistor.NewResistor(130000.0, 2.0, 0.090, resistor.GetHollowMaterial(), env),
// 	// }

// 	// resistor.CalculateAndSetMaterial(arrRes)

// 	fmt.Println(resistor.CalculateRoOpt(arrRes))
// 	resistor.SetMaterialsForResistors(arrRes, resistor.GetMaterials()[0])
// 	fmt.Println(arrRes[0].GetMaterial())

// 	// res := *resistor.NewResistor(110000, 5, 100, resistor.GetHollowMaterial(), env)

// 	// for _, j := range resistor.GetMaterials() {
// 	// 	res.SetMaterial(j)
// 	// 	if res.GetGammaRdelta() > 0 {
// 	// 		fmt.Println("===================================")
// 	// 		fmt.Println(res.GetMaterial())
// 	// 		fmt.Println(res.GetGammaRdelta())
// 	// 		fmt.Println("===================================")
// 	// 	}

// 	// }

// 	for i, j := range arrRes {
// 		if i == 0 {
// 			fmt.Println("======================")
// 			fmt.Println("Номер:", i+1)
// 			fmt.Println("Сопротивление: ", j.GetResistance())
// 			fmt.Println("Коэффициент формы: ", j.GetFromFactor())
// 			fmt.Println("gammaRt: ", j.GetGammaRt())
// 			fmt.Println("gammaRdelta: ", j.GetGammaRdelta())
// 			fmt.Println("Форма: ", j.GetFormOfResistor())
// 			fmt.Println("bp: ", j.GetBp())
// 			fmt.Println("bdelta: ", j.GetBdelta())
// 			fmt.Println("Ширина: ", j.GetWidth())
// 			fmt.Println("lp: ", j.GetLp())
// 			fmt.Println("ldelta: ", j.GetLdelta())
// 			fmt.Println("Длинна: ", j.GetHeight())
// 			fmt.Println("Колличество звеньев мендра: ", j.GetNumberOfLinks())
// 			fmt.Println("X меандр: ", j.GetXlengthMeander())
// 			fmt.Println("Y меандр: ", j.GetYlengthMeander())
// 			fmt.Println("X полоски: ", j.GetXlengthStripes())
// 			fmt.Println("ЦКП bp: ", j.GetBpCCP())
// 			fmt.Println("ЦКП bdelta: ", j.GetBdeltaCCP())
// 			fmt.Println("gammaRdelta подгоночного: ", j.GetGammaRdeltaTrimL())
// 			fmt.Println("gammaR подгоночного:", j.GetGammaRTrimL())
// 			fmt.Println("Количество секций подгонки", j.GetMOfTrimL())
// 			fmt.Println("l нерегулироемое: ", j.GetlnTrimL())
// 			fmt.Println("l регулируемое: ", j.GetLtuneTrimL())
// 			fmt.Println("l общее: ", j.GetloTrimL())
// 			fmt.Println("l прям общее: ", j.GetlooTrimL())
// 			fmt.Println("R'min: ", j.GetRdashminTrimL())
// 			fmt.Println("deltaR: ", j.GetDeltaRTrimL())
// 			fmt.Println("Deltaldash: ", j.GetDeltaLdashTrimL())
// 			fmt.Println("DeltaR: ", j.GetDeltaRTrimL())
// 			fmt.Println("SumL: ", j.GetLsumTrimL())
// 			fmt.Println("deltaR: ", j.GetDeltaRTrimL())
// 			fmt.Println("lpodg: ", j.GetLpodgTrimL())
// 			fmt.Println("deltaLr:", j.GetDeltaLrTrimL())
// 			fmt.Println("rmin:", j.Getrmin())
// 			fmt.Println("rmax:", j.Getrmax())
// 			fmt.Println("lmin:", j.Getlmin())
// 			fmt.Println("lmax:", j.Getlmax())
// 			fmt.Println("R'min: ", j.GetrdashminTrimW())
// 			fmt.Println("deltaR: ", j.GetdeltaRTrimW())
// 			fmt.Println("deltar: ", j.GetdeltarTrimW())
// 			fmt.Println("bdeltaWide: ", j.GetdeltabiTrimW())
// 			fmt.Println("iR: ", j.GetiRTrimW())
// 			fmt.Println("Nwide: ", int(j.GetnTrimW()))
// 			// fmt.Println("ir: ", j.GetirTrimW())
// 			fmt.Println("bo - bn: ", j.GetboTrimW()-j.GetbnTrimW())
// 			var sum float64
// 			for _, j := range *j.GetdeltabiTrimW() {
// 				sum += j
// 			}
// 			fmt.Println("sum: ", sum)
// 			fmt.Println("======================")
// 		}
// 		if j.GetGammaRdelta() < 0 {
// 			consoleoutput.PodgWideResistorPrint(&j)
// 		}
// 	}

// 	// resistor.SetMaterialsForResistors(arrRes, resistor.GetMaterials()[0])
// 	consoleoutput.RenderEnviroment(arrRes)
// 	consoleoutput.RenderTableOfResistors(arrRes)
// 	consoleoutput.RenderTableOfResistorsPractos1(arrRes)

// 	consoleoutput.RenderMaterialsCheck(resistor.CheckAllMaterials(arrRes))
// 	for _, mat := range resistor.GetMaterials() {
// 		resistor.SetMaterialsForResistors(arrRes, mat)
// 		consoleoutput.RenderTableOfResistors(arrRes)
// 	}

// 	server.RunServer()

// }

func main() {

	env := environment.InitEnvironment(70.0)

	arrOfCups := []capacitor.Capacitor{
		*capacitor.NewCapacitor(4500, 6, 20, capacitor.GetHollowMaterial(), env),
		*capacitor.NewCapacitor(400, 6, 20, capacitor.GetHollowMaterial(), env),
		*capacitor.NewCapacitor(1700, 6, 20, capacitor.GetHollowMaterial(), env),
	}

	capacitor.SetMaterialsForCapacitors(arrOfCups, capacitor.Materials[0])
	capacitor.SetCtripledash0ForCapacitors(arrOfCups, arrOfCups[0].GetMaterial().GetCud())

	for i, j := range arrOfCups {
		fmt.Println("==============")
		fmt.Println("N:", i)
		fmt.Println("C: ", j.GetCapacity())
		fmt.Println("Uраб: ", j.GetUrab())
		fmt.Println("Точность: ", j.GetTolerance())
		fmt.Println("e: ", j.GetMaterial().Gete())
		fmt.Println("Площадь: ", j.GetAreaMoreThan5().GetArea())
		fmt.Println("d: ", j.GetAreaMoreThan5().Getd())
		fmt.Println("gammaCt: ", j.GetAreaMoreThan5().GetgammaCt())
		fmt.Println("gammaSdop: ", j.GetAreaMoreThan5().GetGammaSdop())
		fmt.Println("C'0: ", j.GetAreaMoreThan5().GetCdash0())
		fmt.Println("C''0: ", j.GetAreaMoreThan5().GetDdoubledash0())
		fmt.Println("C'''0: ", j.GetCtripledash0())
		fmt.Println("C0: ", j.GetAreaMoreThan5().GetC0())
		fmt.Println("A1: ", j.GetAreaMoreThan5().GetA1())
		fmt.Println("B1: ", j.GetAreaMoreThan5().GetB1())
		fmt.Println("A2: ", j.GetAreaMoreThan5().GetA2())
		fmt.Println("B2: ", j.GetAreaMoreThan5().GetB2())
		fmt.Println("A3: ", j.GetAreaMoreThan5().GetA3())
		fmt.Println("B3: ", j.GetAreaMoreThan5().GetB3())
		fmt.Println("Итоговая площадь: ", j.GetAreaMoreThan5().GetRealArea(), " мм^2")
		fmt.Println("Итоговая толщина: ", j.GetAreaMoreThan5().GetRealD(), " мкм")
		fmt.Println("Площадь  1-5: ", j.GetArea15().GetArea(), "мм^2")
		fmt.Println("Коэфицциент для 1-5: ", j.GetArea15().GetK())
	}
	server.RunServer()
}

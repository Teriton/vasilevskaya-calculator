package consoleoutput

import (
	"example/main/components/resistor"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func RenderTableOfResistors(arrOfRes []resistor.Resistor) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	if len(arrOfRes) > 0 {
		t.AppendHeader(table.Row{"",
			"Материал:" + arrOfRes[0].GetMaterial().GetName(),
			"Roкв: " + strconv.FormatFloat(arrOfRes[0].GetMaterial().GetSquareResistance(), 'g', 10, 64),
			"Roopt: " + strconv.FormatFloat(resistor.CalculateRoOpt(arrOfRes), 'g', 10, 64),
		})
	}
	t.AppendHeader(table.Row{"#", "Сопротивление", "Коэффициент формы", "gammaRdelta", "Форма", "Ширина", "Длинна"})
	for i, j := range arrOfRes {
		t.AppendRow(table.Row{
			i + 1,
			j.GetResistance(),
			j.GetFromFactor(),
			j.GetGammaRdelta(),
			j.GetFormOfResistor(),
			j.GetWidth(),
			j.GetHeight(),
		})
	}
	t.SetStyle(table.StyleColoredBlackOnYellowWhite)
	t.AppendSeparator()
	t.Render()
}

func RenderEnviroment(arrOfRes []resistor.Resistor) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Temperature", "Rokv", "Rocontact", "GammaRt"})
	t.AppendRow(
		table.Row{
			arrOfRes[0].GetEnvironment().GetTemperature(),
			arrOfRes[0].GetEnvironment().GetGammaRokv(),
			arrOfRes[0].GetEnvironment().GetGammaRcontact(),
			arrOfRes[0].GetGammaRt(),
		},
	)
	t.SetStyle(table.StyleColoredBlackOnYellowWhite)
	t.AppendSeparator()
	t.Render()
}

func RenderTableOfResistorsPractos1(arrOfRes []resistor.Resistor) {
	t := table.NewWriter()
	t1 := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t1.SetOutputMirror(os.Stdout)
	// if len(arrOfRes) > 0 {
	// 	t.AppendHeader(table.Row{"",
	// 		"Материал:" + arrOfRes[0].GetMaterial().GetName(),
	// 		"Roкв: " + strconv.FormatFloat(arrOfRes[0].GetMaterial().GetSquareResistance(), 'g', 10, 64),
	// 		"Roopt: " + strconv.FormatFloat(resistor.CalculateRoOpt(arrOfRes), 'g', 10, 64),
	// 	})
	// }
	t.AppendHeader(table.Row{"#", "R,Ом", "P,Вт", "gammaR,%", "кф", "bp,мм", "bdelta,мм", "bmax,мм", "b,мм", "l,мм", "L,мм", "B,мм", "n", "S, мм"})
	t1.AppendHeader(table.Row{"#", "R,Ом", "P,Вт", "gammaR,%", "кф", "lp,мм", "ldelta,мм", "lmax,мм", "b,мм", "l,мм"})
	for i, j := range arrOfRes {
		if j.GetFromFactor() > 1 {
			t.AppendRow(table.Row{
				i + 1,
				j.GetResistance(),
				j.GetPower(),
				j.GetTolerance(),
				j.GetFromFactor(),
				j.GetBp(),
				j.GetBdelta(),
				j.GetWidth(),
				j.GetWidth(),
				j.GetHeight(),
				j.GetXlengthMeander(),
				j.GetYlengthMeander(),
				j.GetNumberOfLinks(),
				j.GetMeanderArea(),
			})
		} else {
			t1.AppendRow(table.Row{
				i + 1,
				j.GetResistance(),
				j.GetPower(),
				j.GetTolerance(),
				j.GetFromFactor(),
				j.GetLp(),
				j.GetLdelta(),
				j.GetHeight(),
				j.GetWidth(),
				j.GetHeight(),
			})
		}

	}
	t.SetStyle(table.StyleColoredBlackOnYellowWhite)
	t.AppendSeparator()
	t1.SetStyle(table.StyleColoredBlackOnYellowWhite)
	t1.AppendSeparator()
	t.Render()
	t1.Render()
}

func RenderMaterialsCheck(arrOfCheckedMaterials []resistor.CheckMaterial) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	currentRow := table.Row{"#", "Название", "Сопротивление, Ом/кв", "Мощность, Вт/см^2", "ТКС*10^-4,1/град", "Старение,%"}
	for i := range arrOfCheckedMaterials[0].GammaRdeltas {
		currentRow = append(currentRow, "R"+strconv.Itoa(i+1))
	}
	t.AppendHeader(currentRow)
	for i, mat := range arrOfCheckedMaterials {
		currentRow = table.Row{
			i + 1,
			mat.Material.GetName(),
			mat.Material.GetSquareResistance(),
			mat.Material.GetPermissibleSpecificPowerDissipation(),
			mat.Material.GetTemperatureCoefficientOfResistance(),
			mat.Material.GetSenescence(),
		}
		for _, j := range mat.GammaRdeltas {
			currentRow = append(currentRow, strconv.FormatFloat(j, 'g', 10, 64))
		}
		t.AppendRow(currentRow)
	}
	t.SetRowPainter(table.RowPainter(func(row table.Row) text.Colors {
		var count int
		for i, j := range row {
			// fmt.Printf("%d %s\n", i, j)
			if i > 5 {
				num, err := strconv.ParseFloat(j.(string), 32)
				if err != nil {
					break
				}
				if num < 0 {
					count++
				}
			}
		}
		switch count {
		case 0:
			return text.Colors{text.BgGreen, text.FgBlack}
		case 1:
			return text.Colors{text.BgHiYellow, text.FgBlack}
		default:
			return text.Colors{text.BgRed, text.FgBlack}
		}
	}))
	t.SetStyle(table.StyleColoredBright)
	t.AppendSeparator()
	t.Render()
}

func PodgWideResistorPrint(res *resistor.Resistor) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Сопротивление", "Точность", "Мощность", "Общая ширина", "breg", "bn", "m", "deltab", "Rmin"})
	t.AppendRow(
		table.Row{
			res.GetResistance(),
			res.GetTolerance(),
			res.GetPower(),
			res.GetboTrimW(),
			res.GetbregTrimW(),
			res.GetbnTrimW(),
			res.GetMOfTrimL(),
			res.GetdeltabiTrimW(),
			res.GetrdashminTrimW(),
		},
	)
	t.SetStyle(table.StyleColoredBright)
	t.AppendSeparator()
	t.Render()
}

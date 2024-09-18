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

func RenderTableOfResistorsPractos1(arrOfRes []resistor.Resistor) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	// if len(arrOfRes) > 0 {
	// 	t.AppendHeader(table.Row{"",
	// 		"Материал:" + arrOfRes[0].GetMaterial().GetName(),
	// 		"Roкв: " + strconv.FormatFloat(arrOfRes[0].GetMaterial().GetSquareResistance(), 'g', 10, 64),
	// 		"Roopt: " + strconv.FormatFloat(resistor.CalculateRoOpt(arrOfRes), 'g', 10, 64),
	// 	})
	// }
	t.AppendHeader(table.Row{"#", "R,Ом", "P,Вт", "gammaR,%", "кф", "bp,мм", "bdelta,мм", "bmax,мм", "b,мм", "l,мм", "L,мм", "B,мм", "n"})
	for i, j := range arrOfRes {
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
		})
	}
	t.SetStyle(table.StyleColoredBlackOnYellowWhite)
	t.AppendSeparator()
	t.Render()
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

package datacontroller

import (
	"github.com/shpakunya/pkg/components"
)

type DataController struct {
	ReadMaterialsRes
}

type ReadMaterialsRes interface {
	GetMatreadMaterialsRes() (length int, names []string, sq []float64, ppd []float64, tcr []float64, sen []float64)
}

func convertToMaterial(name string, sr float64, ppd float64, tcr float64, s float64) components.ResistorMaterial {
	return components.ResistorMaterial{
		Name: components.ParameterMain[string]{Name: "Название материала", Symbol: "", Unit: "", Value: name},
		SquareResistance: components.Parameter{Name: "Квадратное сопротивление", Symbol: "ρкв",
			Unit: "Ом/кв", Value: sr},
		PermissibleSpecificPowerDissipation: components.Parameter{Name: "Допустимая удельная мощность рассеивания",
			Unit: "Вт/см^2", Value: ppd},
		TemperatureCoefficientOfResistance: components.Parameter{Name: "ТКС", Symbol: "aR",
			Unit: "1/град * 10^-4", Value: tcr},
		Senescence: components.Parameter{Name: "Необратимое относительное изменение сопротивления после 1000 ч работы под нагрузкой",
			Unit: "(1 Вт/см^2 при 85 °C), %", Value: s},
	}
}

func (dc DataController) GetMaterials() []components.ResistorMaterial {
	num, names, sqs, ppds, tcrs, sens := dc.GetMatreadMaterialsRes()
	resistorMaterials := make([]components.ResistorMaterial, 0)
	for i := range num {
		resistorMaterials = append(resistorMaterials, convertToMaterial(names[i], sqs[i], ppds[i], tcrs[i], sens[i]))
	}
	return resistorMaterials
}

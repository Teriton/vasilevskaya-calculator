package resistor

import "math"

type material struct {
	name                                string
	squareResistance                    float64
	permissibleSpecificPowerDissipation float64
	temperatureCoefficientOfResistance  float64
	senescence                          float64
}

func GetHollowMaterial() material {
	return material{
		name:                                "Hollow",
		squareResistance:                    0,
		permissibleSpecificPowerDissipation: 0,
		temperatureCoefficientOfResistance:  0,
		senescence:                          0,
	}
}

func GetMaterials() []material {

	return []material{
		{
			name:                                "Кермет К50-С", // 0
			squareResistance:                    10000.0,
			permissibleSpecificPowerDissipation: 2.0,
			temperatureCoefficientOfResistance:  1.0,
			senescence:                          0,
		},
		{
			name:                                "Хром", // 1
			squareResistance:                    500.0,
			permissibleSpecificPowerDissipation: 1.0,
			temperatureCoefficientOfResistance:  2.0,
			senescence:                          0,
		},
		{
			name:                                "Тантал", // 2
			squareResistance:                    100.0,
			permissibleSpecificPowerDissipation: 3.0,
			temperatureCoefficientOfResistance:  -2.0,
			senescence:                          0,
		},
		{
			name:                                "Нихром", // 3
			squareResistance:                    300.0,
			permissibleSpecificPowerDissipation: 2.0,
			temperatureCoefficientOfResistance:  2.5,
			senescence:                          0,
		},
		{
			name:                                "Сплав МЛТ3-М", //4
			squareResistance:                    500.0,
			permissibleSpecificPowerDissipation: 2.0,
			temperatureCoefficientOfResistance:  2.4,
			senescence:                          0.5,
		},
		{
			name:                                "Кермет К-20С", //5
			squareResistance:                    3000.0,
			permissibleSpecificPowerDissipation: 2.0,
			temperatureCoefficientOfResistance:  0.5,
			senescence:                          0,
		},
		{
			name:                                "Сплав 1004", // 6
			squareResistance:                    50000.0,
			permissibleSpecificPowerDissipation: 5.0,
			temperatureCoefficientOfResistance:  15.0,
			senescence:                          2,
		},
		{
			name:                                "Сплав 1714", // 7
			squareResistance:                    500.0,
			permissibleSpecificPowerDissipation: 5.0,
			temperatureCoefficientOfResistance:  2.0,
			senescence:                          1,
		},
		{
			name:                                "Сплав 2005", // 8
			squareResistance:                    500000.0,
			permissibleSpecificPowerDissipation: 5.0,
			temperatureCoefficientOfResistance:  12.0,
			senescence:                          2,
		},
		{
			name:                                "Сплав 2310", // 9
			squareResistance:                    50000.0,
			permissibleSpecificPowerDissipation: 5.0,
			temperatureCoefficientOfResistance:  12.0,
			senescence:                          2,
		},
		{
			name:                                "Сплав 3001", // 10
			squareResistance:                    3000.0,
			permissibleSpecificPowerDissipation: 5.0,
			temperatureCoefficientOfResistance:  1.0,
			senescence:                          0.5,
		},
		{
			name:                                "Сплав 3710", // 11
			squareResistance:                    3000.0,
			permissibleSpecificPowerDissipation: 5.0,
			temperatureCoefficientOfResistance:  1.0,
			senescence:                          0.5,
		},
		{
			name:                                "Сплав 4206", // 12
			squareResistance:                    1000.0,
			permissibleSpecificPowerDissipation: 2.0,
			temperatureCoefficientOfResistance:  0.5,
			senescence:                          0.5,
		},
		{
			name:                                "Сплав 4400", // 13
			squareResistance:                    5000.0,
			permissibleSpecificPowerDissipation: 10.0,
			temperatureCoefficientOfResistance:  3.0,
			senescence:                          0,
		},
		{
			name:                                "Сплав 4800", // 14
			squareResistance:                    1000.0,
			permissibleSpecificPowerDissipation: 5.0,
			temperatureCoefficientOfResistance:  2,
			senescence:                          1,
		},
		{
			name:                                "Сплав 5006", // 15
			squareResistance:                    20.0,
			permissibleSpecificPowerDissipation: 5.0,
			temperatureCoefficientOfResistance:  0.5,
			senescence:                          2,
		},
		{
			name:                                "Сплав 5402", // 16
			squareResistance:                    100.0,
			permissibleSpecificPowerDissipation: 2.0,
			temperatureCoefficientOfResistance:  0.5,
			senescence:                          1,
		},
		{
			name:                                "Сплав 5406 К", // 17
			squareResistance:                    500.0,
			permissibleSpecificPowerDissipation: 2.0,
			temperatureCoefficientOfResistance:  0.5,
			senescence:                          1,
		},
		{
			name:                                "Сплав 5406 Н", // 18
			squareResistance:                    500.0,
			permissibleSpecificPowerDissipation: 2.0,
			temperatureCoefficientOfResistance:  0.3,
			senescence:                          1,
		},
	}
}

func (m *material) GetName() string {
	return m.name
}

func (m *material) GetSquareResistance() float64 {
	return m.squareResistance
}

func (m *material) GetPermissibleSpecificPowerDissipation() float64 {
	return m.permissibleSpecificPowerDissipation
}

func (m *material) GetTemperatureCoefficientOfResistance() float64 {
	return m.temperatureCoefficientOfResistance
}

func (m *material) GetSenescence() float64 {
	return m.senescence
}

type CheckMaterial struct {
	Material     material
	GammaRdeltas []float64
}

func CheckAllMaterials(arrOfRes []Resistor) []CheckMaterial {
	var arrOfCheckedMaterials []CheckMaterial
	materials := GetMaterials()
	for _, mat := range materials {
		var currentrlyCheckingMaterial CheckMaterial
		currentrlyCheckingMaterial.Material = mat
		for _, res := range arrOfRes {
			res.SetMaterial(mat)
			currentrlyCheckingMaterial.GammaRdeltas = append(currentrlyCheckingMaterial.GammaRdeltas, res.gammaRdelta)
		}
		arrOfCheckedMaterials = append(arrOfCheckedMaterials, currentrlyCheckingMaterial)
	}
	return arrOfCheckedMaterials
}

func CalculateRoOpt(arrOfMat []Resistor) float64 {
	var sum, antisum float64

	for _, j := range arrOfMat {
		sum += j.GetResistance()
		antisum += (1 / j.GetResistance())
	}

	return math.Sqrt(sum / antisum)
}

func CalculateMaterial(roOpt float64) material {
	arrMat := GetMaterials()
	var mat material = arrMat[0]

	for _, j := range arrMat {
		if math.Abs(j.squareResistance-roOpt) < math.Abs(mat.squareResistance-roOpt) {
			mat = j
		}
	}
	return mat
}

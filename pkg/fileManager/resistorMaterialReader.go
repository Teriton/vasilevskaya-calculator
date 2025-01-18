package filemanager

import (
	"log"
	"os"
	"strconv"
)

type materialResistor struct {
	Name                                string
	SquareResistance                    float64
	PermissibleSpecificPowerDissipation float64
	TemperatureCoefficientOfResistance  float64
	Senescence                          float64
}

func (FileManager) GetMatreadMaterialsRes() (length int, names []string, sq []float64, ppd []float64, tcr []float64, sen []float64) {
	mats := convertStrsToMaterials(readFileContent("./data/resistorMaterials.csv"))
	length = len(mats)
	for _, mat := range mats {
		names = append(names, mat.Name)
		sq = append(sq, mat.SquareResistance)
		ppd = append(ppd, mat.PermissibleSpecificPowerDissipation)
		tcr = append(tcr, mat.TemperatureCoefficientOfResistance)
		sen = append(sen, mat.Senescence)
	}
	return
}

func convertStrsToMaterials(matStrs [][]string) []materialResistor {
	mats := make([]materialResistor, 0)
	for num, mat := range matStrs[1:] {
		if len(mat) < 5 {
			log.Fatalf("FileManager: not enough parameters of resistor material, check material number %d", num)
			os.Exit(1)
		}
		mats = append(mats, materialResistor{
			Name:                                mat[0],
			SquareResistance:                    checkParam(strconv.ParseFloat(mat[1], 64)),
			PermissibleSpecificPowerDissipation: checkParam(strconv.ParseFloat(mat[2], 64)),
			TemperatureCoefficientOfResistance:  checkParam(strconv.ParseFloat(mat[3], 64)),
			Senescence:                          checkParam(strconv.ParseFloat(mat[4], 64)),
		})
	}
	return mats
}

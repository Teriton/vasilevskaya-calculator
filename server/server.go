package server

import (
	"encoding/json"
	"example/main/components/resistor"
	"example/main/environment"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type OutputData struct {
	Resistance float64
	Tolerance  float64
	Power      float64

	NameOfMaterial                      string
	SquareResistance                    float64
	PermissibleSpecificPowerDissipation float64
	TemperatureCoefficientOfResistance  float64
	Senescence                          float64

	Temperature   float64
	GammaRokv     float64
	GammaRcontact float64
	Deltab        float64
	Deltal        float64
	Btehn         float64
	Ltehn         float64

	GammaRt        float64
	FormOfResistor string

	FormFactor  float64
	GammaRdelta float64
	RoOpt       float64

	Bp     float64
	Bdelta float64
	Width  float64
	Lp     float64
	Ldelta float64
	Length float64

	NOfMeander  float64
	Xmeander    float64
	Ymeander    float64
	MeanderArea float64

	LSumTrim float64
	DeltaLr  float64
	RminTrim float64

	Bmin, Bmax       float64
	Lmin, Lmax       float64
	Rokvmin, Rokvmax float64
	Rmin, Rmax       float64

	N        float64
	Bo       float64
	Bn       float64
	Rdashmin float64
	Deltar   float64
	DeltaR   float64
	Ir       []float64
	IR       []float64
	Deltabi  []float64
	Breg     float64

	GammaRdeltaTrim float64
	GammaR          float64
	MOfTrim         float64
	LnTrim          float64
	LoTrim          float64
	RdashminTrim    float64
	Ltune           float64
	DeltaLrTrim     float64
	DeltaLdashTrim  float64
	Lpodg           float64
	Lsum            float64
	Gammakf         float64
}

type ResistorJSON struct {
	Resistance string `json:"resistance"`
	Power      string `json:"power"`
	Tolerance  string `json:"tolerance"`
}
type InputJSON struct {
	Temperature float64        `json:"temperature"`
	Material    int            `json:"material"`
	Res         []ResistorJSON `json:"res"`
}

func countSingleResistor(c *gin.Context) {
	// var res resistor.Resistor
	resistance, _ := strconv.ParseFloat(c.Query("resistance"), 32) // shortcut for c.Request.URL.Query().Get("lastname")
	power, _ := strconv.ParseFloat(c.Query("power"), 32)
	tolerance, _ := strconv.ParseFloat(c.Query("tolerance"), 32)
	env := environment.InitEnvironment(70)

	arrRes := []resistor.Resistor{
		*resistor.NewResistor(resistance, tolerance, power, resistor.GetHollowMaterial(), env),
	}

	resistor.CalculateAndSetMaterial(arrRes)

	c.JSON(200, gin.H{
		"formFactor": arrRes[0].GetFromFactor(),
		"material":   arrRes[0].GetMaterial().GetName(),
	})
}

func countArrOfRes(c *gin.Context) {
	var arrResJSON InputJSON

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &arrResJSON)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	var arrRes []resistor.Resistor
	for _, j := range arrResJSON.Res {
		resistance, _ := strconv.ParseFloat(j.Resistance, 32)
		power, _ := strconv.ParseFloat(j.Power, 32)
		tolerance, _ := strconv.ParseFloat(j.Tolerance, 32)
		env := environment.InitEnvironment(arrResJSON.Temperature)
		arrRes = append(arrRes, *resistor.NewResistor(resistance, tolerance, power, resistor.GetHollowMaterial(), env))
	}

	if arrResJSON.Material < 0 || arrResJSON.Material > len(resistor.GetMaterials()) {
		resistor.CalculateAndSetMaterial(arrRes)
	} else {
		resistor.SetMaterialsForResistors(arrRes, resistor.GetMaterials()[arrResJSON.Material])
	}

	var outputData []OutputData

	for _, j := range arrRes {
		outputData = append(outputData, OutputData{
			Resistance: j.GetResistance(),
			Tolerance:  j.GetTolerance(),
			Power:      j.GetPower(),

			NameOfMaterial:                      j.GetMaterial().GetName(),
			SquareResistance:                    j.GetMaterial().GetSquareResistance(),
			PermissibleSpecificPowerDissipation: j.GetMaterial().GetPermissibleSpecificPowerDissipation(),
			TemperatureCoefficientOfResistance:  j.GetMaterial().GetTemperatureCoefficientOfResistance(),
			Senescence:                          j.GetMaterial().GetSenescence(),

			Temperature:   j.GetEnvironment().GetTemperature(),
			GammaRokv:     j.GetEnvironment().GetGammaRokv(),
			GammaRcontact: j.GetEnvironment().GetGammaRcontact(),
			Deltab:        j.GetEnvironment().GetDeltab(),
			Deltal:        j.GetEnvironment().GetDeltal(),

			GammaRt:        j.GetGammaRt(),
			FormOfResistor: string(j.GetFormOfResistor()),

			FormFactor:  j.GetFromFactor(),
			GammaRdelta: j.GetGammaRdelta(),
			RoOpt:       resistor.CalculateRoOpt(arrRes),

			Bp:     j.GetBp(),
			Bdelta: j.GetBdelta(),
			Width:  j.GetWidth(),
			Lp:     j.GetLp(),
			Ldelta: j.GetLdelta(),
			Length: j.GetHeight(),

			NOfMeander:  j.GetNumberOfLinks(),
			Xmeander:    j.GetXlengthMeander(),
			Ymeander:    j.GetYlengthMeander(),
			MeanderArea: j.GetMeanderArea(),

			Bmin:    j.Getbmin(),
			Bmax:    j.Getbmax(),
			Lmin:    j.Getlmin(),
			Lmax:    j.Getlmax(),
			Rokvmin: j.Getrokvmin(),
			Rokvmax: j.Getrokvmax(),
			Rmin:    j.Getrmin(),
			Rmax:    j.Getrmax(),

			N:        j.GetnTrimW(),
			Bo:       j.GetboTrimW(),
			Bn:       j.GetbnTrimW(),
			Rdashmin: j.GetrdashminTrimW(),
			Deltar:   j.GetdeltarTrimW(),
			DeltaR:   j.GetdeltaRTrimW(),
			Ir:       *j.GetirTrimW(),
			IR:       *j.GetiRTrimW(),
			Deltabi:  *j.GetdeltabiTrimW(),
			Breg:     j.GetbregTrimW(),

			GammaRdeltaTrim: j.GetGammaRdeltaTrimL(),
			GammaR:          j.GetGammaRTrimL(),
			MOfTrim:         j.GetMOfTrimL(),
			LnTrim:          j.GetlnTrimL(),
			LoTrim:          j.GetloTrimL(),
			RdashminTrim:    j.GetRdashminTrimL(),
			Ltune:           j.GetLtuneTrimL(),
			DeltaLrTrim:     j.GetDeltaLrTrimL(),
			DeltaLdashTrim:  j.GetDeltaLdashTrimL(),
			Lpodg:           j.GetLpodgTrimL(),
			Lsum:            j.GetLsumTrimL(),
			Gammakf:         j.GetGammakfTrimL(),
		})
	}

	c.IndentedJSON(200, outputData)
}

type ResitsorNums struct {
	Nums []string `form:"res" collection_format:"csv"`
}

func resistorsQuery(c *gin.Context) {
	var numsString ResitsorNums
	var nums []float64
	if c.ShouldBind(&numsString) == nil {
		fmt.Println("log")
		fmt.Println(numsString.Nums)
	}
	for _, j := range numsString.Nums {
		num, err := strconv.ParseFloat(j, 64)
		if err == nil {
			nums = append(nums, num)
		}
		fmt.Println(j)
	}
	fmt.Println(nums)
	c.String(200, "Success")
}

func mainFraim(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main frame",
	})
}

type MaterialJSON struct {
	Number                              int
	Name                                string
	SquareResistance                    float64
	PermissibleSpecificPowerDissipation float64
	TemperatureCoefficientOfResistance  float64
	Senescence                          float64
}

func resistorMaterials(c *gin.Context) {
	var arrOfMaterials []MaterialJSON
	for i, j := range resistor.GetMaterials() {
		arrOfMaterials = append(arrOfMaterials, MaterialJSON{
			Number:                              i,
			Name:                                j.GetName(),
			SquareResistance:                    j.GetSquareResistance(),
			PermissibleSpecificPowerDissipation: j.GetPermissibleSpecificPowerDissipation(),
			TemperatureCoefficientOfResistance:  j.GetTemperatureCoefficientOfResistance(),
			Senescence:                          j.GetSenescence(),
		})
	}
	c.IndentedJSON(200, arrOfMaterials)
}

func RunServer() {
	router := gin.Default()
	router.Use(cors.Default())

	router.LoadHTMLGlob("./index/*")
	router.POST("/countResistor", countSingleResistor)
	router.GET("/resistorMaterials", resistorMaterials)
	router.POST("/arrOfRes", countArrOfRes)
	router.GET("/", mainFraim)
	router.GET("/resistors", resistorsQuery)
	router.Run(":8080")
}

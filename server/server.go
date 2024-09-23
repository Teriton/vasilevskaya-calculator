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
	FormFactor  float64
	Temperature float64
	Form        string
	Material    string
	GammaRdelta float64
	RoOpt       float64
	Bp          float64
	Bdelta      float64
	Width       float64
	Lp          float64
	Ldelta      float64
	Length      float64
	NOfMeander  float64
	Xmeander    float64
	Ymeander    float64
	LSumTrim    float64
	MOfTrim     float64
	Lpodg       float64
	DeltaR      float64
	DeltaLr     float64
	RminTrim    float64
}

type ResistorJSON struct {
	Resistance string `json:"resistance"`
	Power      string `json:"power"`
	Tolerance  string `json:"tolerance"`
}
type InputJSON struct {
	Temperature string         `json:"temperature"`
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
	temperature, err := strconv.ParseFloat(arrResJSON.Temperature, 64)
	if err != nil {
		return
	}
	var arrRes []resistor.Resistor
	for _, j := range arrResJSON.Res {
		resistance, _ := strconv.ParseFloat(j.Resistance, 32)
		power, _ := strconv.ParseFloat(j.Power, 32)
		tolerance, _ := strconv.ParseFloat(j.Tolerance, 32)
		env := environment.InitEnvironment(temperature)
		arrRes = append(arrRes, *resistor.NewResistor(resistance, tolerance, power, resistor.GetHollowMaterial(), env))
	}

	// resistor.CalculateAndSetMaterial(arrRes)
	resistor.SetMaterialsForResistors(arrRes, resistor.GetMaterials()[0])

	var outputData []OutputData

	for _, j := range arrRes {
		outputData = append(outputData, OutputData{
			FormFactor:  j.GetFromFactor(),
			Form:        string(j.GetFormOfResistor()),
			Material:    j.GetMaterial().GetName(),
			GammaRdelta: j.GetGammaRdelta(),
			RoOpt:       resistor.CalculateRoOpt(arrRes),
			Bp:          j.GetBp(),
			Bdelta:      j.GetBdelta(),
			Width:       j.GetWidth(),
			Lp:          j.GetLp(),
			Ldelta:      j.GetLdelta(),
			Length:      j.GetHeight(),
			NOfMeander:  j.GetNumberOfLinks(),
			Xmeander:    j.GetXlengthMeander(),
			Ymeander:    j.GetYlengthMeander(),
			LSumTrim:    j.GetLsum(),
			MOfTrim:     j.GetMOfTrim(),
			Lpodg:       j.GetLpodg(),
			DeltaR:      j.GetDeltaLrTrim(),
			DeltaLr:     j.GetDeltaLrTrim(),
			RminTrim:    j.GetRdashminTrim(),
			Temperature: j.GetEnvironment().GetTemperature(),
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

func RunServer() {
	router := gin.Default()
	router.Use(cors.Default())

	router.LoadHTMLGlob("./index/*")
	router.POST("/countResistor", countSingleResistor)
	router.POST("/arrOfRes", countArrOfRes)
	router.GET("/", mainFraim)
	router.GET("/resistors", resistorsQuery)
	router.Run(":8080")
}

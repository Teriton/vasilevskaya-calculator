package server

import (
	"encoding/json"
	"example/main/components/resistor"
	"example/main/environment"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OutputData struct {
	FormFactor  float64
	Form        string
	Material    string
	GammaRdelta float64
	RoOpt       float64
}

type ResistorJSON struct {
	Resistance string `json:"resistance"`
	Power      string `json:"power"`
	Tolerance  string `json:"tolerance"`
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
		"material":   arrRes[0].GetMaterial(),
	})
}

func countArrOfRes(c *gin.Context) {
	var arrResJSON []ResistorJSON

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
	for _, j := range arrResJSON {
		resistance, _ := strconv.ParseFloat(j.Resistance, 32)
		power, _ := strconv.ParseFloat(j.Power, 32)
		tolerance, _ := strconv.ParseFloat(j.Tolerance, 32)
		env := environment.InitEnvironment(70)
		arrRes = append(arrRes, *resistor.NewResistor(resistance, tolerance, power, resistor.GetHollowMaterial(), env))
	}

	resistor.CalculateAndSetMaterial(arrRes)

	var outputData []OutputData

	for _, j := range arrRes {
		outputData = append(outputData, OutputData{
			FormFactor:  j.GetFromFactor(),
			Form:        string(j.GetFormOfResistor()),
			Material:    j.GetMaterial().GetName(),
			GammaRdelta: j.GetGammaRdelta(),
			RoOpt:       resistor.CalculateRoOpt(arrRes),
		})
	}

	c.IndentedJSON(200, outputData)

}

func RunServer() {
	router := gin.Default()
	router.GET("/countResistor", countSingleResistor)
	router.POST("/arrOfRes", countArrOfRes)
	router.Run(":8080")
}

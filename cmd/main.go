package main

import (
	"fmt"
	"reflect"

	"github.com/shpakunya/pkg/calculator"
	datacontroller "github.com/shpakunya/pkg/controllers/dataController"
	environmentControllerIml "github.com/shpakunya/pkg/controllers/environmentController"
	resistorcontrollerIml "github.com/shpakunya/pkg/controllers/resistorController"
	filemanager "github.com/shpakunya/pkg/fileManager"
)

func traverseStruct(v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Struct {
			fmt.Println()
			traverseStruct(field)
		} else {
			fmt.Print(field, " ")
		}
	}
}

func main() {
	calc := calculator.InitCalculator(resistorcontrollerIml.ResistorController{
		DataController: datacontroller.DataController{
			ReadMaterialsRes: filemanager.FileManager{},
		},
	}, environmentControllerIml.EnvironmentController{})
	env := calc.NewEnvironment(100)
	res := calc.NewResistor(10000, 20, 35, calc.GetMaterials()[0], env)
	v := reflect.ValueOf(res)
	traverseStruct(v)
}

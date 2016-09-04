package main

import (
	"cms/src/controllers"
	"fmt"

	"reflect"
)

func main1() {
	baseController := new(controllers.BaseController)
	typ := reflect.TypeOf(baseController)
	baseMethod := make(map[string]bool, typ.NumMethod())
	for i := 0; i < typ.NumMethod(); i++ {
		baseMethod[typ.Method(i).Name] = true
	}

	controller := new(controllers.MainController)
	controllerType := reflect.TypeOf(controller)

	for i := 0; i < controllerType.NumMethod(); i++ {
		methodName := controllerType.Method(i).Name
		if _, ok := baseMethod[methodName]; !ok {
			fmt.Println(methodName)
		}
	}

}

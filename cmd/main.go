package main

import (
	"fmt"

	"github.com/putcn/ncml/ops"
	"github.com/putcn/ncml/vars"
)

func main() {
	opRepo := ops.GetOpRepo()
	opFlow := []ops.OpDesc{
		ops.OpDesc{
			Name: "example",
			Vars: []string{"var1", "var2"},
		},
		ops.OpDesc{
			Name: "example1",
			Vars: []string{"var3", "var4"},
		},
	}

	scope := vars.NewScope(nil)

	fmt.Println("main: start to create op instances")

	for _, desc := range opFlow {
		//create op instances
		op := opRepo.CreateByName(desc.Name)
		for _, varname := range desc.Vars {
			scope.CreateVar(varname, "float32")
		}
		op.Forward(scope) //testing only

		//create vars
	}

	//create backward pass by dependencies

	//run
}

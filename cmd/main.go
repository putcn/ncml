package main

import (
	"context"
	"fmt"

	"github.com/putcn/ncml/exec"
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
	scope1 := vars.NewScope(scope) //testing with nested scopes
	execq := exec.NewExecq()

	fmt.Println("main: start to create op instances")

	for _, desc := range opFlow {
		//create op instances
		op := opRepo.CreateOpByName(desc.Name)
		for _, varname := range desc.Vars {
			scope.CreateVar(varname, "float32")
		}
		execq.Add(op.Forward, scope1)
	}

	fmt.Println("main: testing passid", scope1.FetchVar("var1"))

	//create backward pass by dependencies

	//run
	ctx := context.Background()
	execq.Run(ctx)
	fmt.Println("main: testing passid", scope1.FetchVar("var1"))
}

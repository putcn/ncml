package main

import (
	"github.com/putcn/ncml/ops"
)

func main() {
	opRepo := ops.GetOpRepo()
	opFlow := []ops.OpDesc{ops.OpDesc{Name: "example"}, ops.OpDesc{Name: "example1"}}

	for _, desc := range opFlow {
		//create op instances
		op := opRepo.GetByName(desc.Name)
		op.Forward() //testing only

		//create vars
	}

	//create backward pass by dependencies

	//run
}

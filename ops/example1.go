package ops

import (
	"fmt"

	"github.com/putcn/ncml/vars"
)

type Example1_op struct {
	testfield string
}

func (e *Example1_op) Forward(scope *vars.Scope) {
	fmt.Printf("example_op1: %s\n", e.testfield)
}

func (e *Example1_op) Backward(scope *vars.Scope) {

}

func init() {
	fmt.Println("example_op1: registering op example 1")
	GetOpRepo().Register("example1", func() Operator {
		return &Example1_op{
			testfield: "example 1 forward pass",
		}
	})
}

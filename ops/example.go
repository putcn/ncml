package ops

import (
	"fmt"

	"github.com/putcn/ncml/vars"
)

type Example_op struct {
	Operator
	testfield string
}

func (e *Example_op) Forward(scope *vars.Scope) {
	fmt.Printf("example_op: %s\n", e.testfield)
	tempVar := scope.FetchVar("var1")
	fmt.Printf("example_op: %s my value\n", tempVar)
}

func (e *Example_op) Backward(scope *vars.Scope) {

}

func init() {
	fmt.Println("example_op: registering op example")
	GetOpRepo().Register("example", func() Operator {
		return &Example_op{
			testfield: "example forward pass",
		}
	})
}

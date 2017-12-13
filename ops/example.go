package ops

import (
	"fmt"
)

type Example_op struct {
	testfield string
}

func (e *Example_op) Forward() {
	fmt.Printf("%s\n", e.testfield)
}

func (e *Example_op) Backward() {

}

func init() {
	fmt.Println("registering op example")
	GetOpRepo().Register("example", func() Operator {
		return &Example_op{
			testfield: "example forward pass",
		}
	})
}

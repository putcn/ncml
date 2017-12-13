package ops

import "fmt"

type Example1_op struct {
	testfield string
}

func (e *Example1_op) Forward() {
	fmt.Printf("%s\n", e.testfield)
}

func (e *Example1_op) Backward() {

}

func init() {
	fmt.Println("registering op example 1")
	GetOpRepo().Register("example1", func() Operator {
		return &Example1_op{
			testfield: "example 1 forward pass",
		}
	})
}

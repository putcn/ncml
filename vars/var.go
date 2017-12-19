package vars

import (
	"fmt"
)

type Var struct {
	typeName string
	value    interface{}
	PassId   int
}

func NewVar(name string, typeName string) *Var {
	fmt.Println("var: creating var instance", name)
	return &Var{
		typeName: typeName,
		value:    name, //for testing only
		PassId:   0,
	}
}

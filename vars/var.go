package vars

import (
	"fmt"
)

type Var struct {
	typeName string
	value    interface{}
}

func NewVar(name string, typeName string) *Var {
	fmt.Println("var: creating var instance")
	return &Var{
		typeName: typeName,
		value:    name, //for testing only
	}
}

package vars

import (
	"fmt"
)

type Var struct {
	IVar
	TypeName string
	Value    interface{}
	PassID   int
}

type IVar interface {
	setPassID(int)
	getPassID() int
}

func NewVar(name string, typeName string) *Var {
	fmt.Println("var: creating var instance", name)
	return &Var{
		TypeName: typeName,
		Value:    name, //for testing only
		PassID:   0,
	}
}

func (v *Var) setPassID(passID int) {
	mu.Lock()
	defer mu.Unlock()

	v.PassID = passID
}

func (v *Var) getPassID() int {
	return v.PassID
}

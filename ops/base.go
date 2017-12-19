package ops

import (
	"github.com/putcn/ncml/vars"
)

type Operator interface {
	Forward(*vars.Scope)
	Backward(*vars.Scope)
}

type OpCreator func() Operator

type OpDesc struct {
	Name string
	Vars []string
}

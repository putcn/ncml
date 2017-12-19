package exec

import (
	"context"
	"fmt"

	"github.com/putcn/ncml/vars"
)

type Execq struct {
	fstore []func(*vars.Scope)
	sstore []*vars.Scope
}

func NewExecq() *Execq {
	return &Execq{
		fstore: []func(*vars.Scope){},
		sstore: []*vars.Scope{},
	}
}

func (e *Execq) Add(f func(*vars.Scope), scope *vars.Scope) {
	fmt.Println("new func added")
	e.fstore = append(e.fstore, f)
	e.sstore = append(e.sstore, scope)
}

func (e *Execq) Run(ctx context.Context) {
	fmt.Println("execting whole queue")
	for index, f := range e.fstore {
		f(e.sstore[index])
	}
}

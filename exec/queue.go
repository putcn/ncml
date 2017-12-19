package exec

import (
	"context"
	"fmt"

	"github.com/putcn/ncml/vars"
)

var passId int

type Execq struct {
	fstore []func(*vars.Scope)
	sstore []*vars.Scope
	PassId int
}

func generateId() int {
	passId += 1
	return passId
}

func NewExecq() *Execq {
	e := Execq{
		fstore: []func(*vars.Scope){},
		sstore: []*vars.Scope{},
		PassId: 0,
	}
	e.PreparePass()
	return &e
}

func (e *Execq) Add(f func(*vars.Scope), scope *vars.Scope) {
	fmt.Println("exec: new func added")
	e.fstore = append(e.fstore, f)
	e.sstore = append(e.sstore, scope)
}

func (e *Execq) Run(ctx context.Context) {
	fmt.Println("exec: execting whole queue")
	for index, f := range e.fstore {
		s := e.sstore[index]
		s.PassId = e.PassId
		f(s)
	}
}

func (e *Execq) PreparePass() {
	e.PassId = generateId()
}

func init() {
	passId = 0
}

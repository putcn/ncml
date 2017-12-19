package exec

import (
	"context"
	"fmt"

	"github.com/putcn/ncml/vars"
)

var passID int

type Execq struct {
	fstore []func(*vars.Scope)
	sstore []*vars.Scope
	PassID int
}

func generateID() int {
	passID++
	return passID
}

func NewExecq() *Execq {
	e := Execq{
		fstore: []func(*vars.Scope){},
		sstore: []*vars.Scope{},
		PassID: 0,
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
		s.PassID = e.PassID
		f(s)
	}
}

func (e *Execq) PreparePass() {
	e.PassID = generateID()
}

func init() {
	passID = 0
}

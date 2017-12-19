package vars

import (
	"fmt"
)

type Scope struct {
	parent *Scope
	vars   map[string]*Var
}

func NewScope(parent *Scope) *Scope {
	fmt.Println("scope: creating new scope")
	return &Scope{
		vars:   map[string]*Var{},
		parent: parent,
	}
}

func (s *Scope) CreateVar(name string, typeName string) *Var {
	fmt.Println("scope: creating var and store it")
	if typeName == "" {
		typeName = "float32"
	}
	newVar := NewVar(name, typeName)
	s.StoreVar(name, newVar)
	return newVar
}

func (s *Scope) StoreVar(name string, val *Var) *Var {
	if val != nil {
		s.vars[name] = val
	}
	return s.vars[name]
}

func (s *Scope) FetchVar(name string) *Var {
	fmt.Println("scope: fetching var for ", name)
	return s.vars[name]
}

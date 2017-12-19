package vars

import (
	"fmt"
	"sync"
)

type Scope struct {
	id     int
	parent *Scope
	vars   map[string]*Var
	PassId int
}

var currentId int
var mu sync.Mutex

func generateId() int {
	mu.Lock()
	defer mu.Unlock()
	currentId = currentId + 1
	fmt.Println("creating new scope id", currentId)
	return currentId
}

func NewScope(parent *Scope) *Scope {
	fmt.Println("scope: creating new scope")
	return &Scope{
		id:     generateId(),
		vars:   map[string]*Var{},
		parent: parent,
		PassId: 0,
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
		val.PassId = s.PassId
		s.vars[name] = val
	}
	return s.vars[name]
}

func (s *Scope) FetchVar(name string) *Var {
	fmt.Println("scope: fetching var for ", name)
	varInstance, ok := s.vars[name]
	if !ok {
		if s.parent != nil {
			fmt.Println("scope: var not found, trying with parent")
			varInstance = s.parent.FetchVar(name)
		} else {
			fmt.Println("scope: no parent, return nil")
			varInstance = nil
		}
	} else {
		fmt.Println("scope: var fount, return the var")
	}
	return varInstance
}

func init() {
	currentId = 0
}

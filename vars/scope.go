package vars

import (
	"fmt"
	"sync"
)

type Scope struct {
	id     int
	parent *Scope
	vars   map[string]*Var
	PassID int
}

var currentID int
var mu sync.Mutex

func generateID() int {
	mu.Lock()
	defer mu.Unlock()
	currentID++
	fmt.Println("creating new scope id", currentID)
	return currentID
}

func NewScope(parent *Scope) *Scope {
	fmt.Println("scope: creating new scope")
	return &Scope{
		id:     generateID(),
		vars:   map[string]*Var{},
		parent: parent,
		PassID: 0,
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
		val.PassID = s.PassID
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
	currentID = 0
}

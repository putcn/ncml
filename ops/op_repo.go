package ops

import (
	"fmt"
	"sync"
)

var _opRepo *OpRepo
var mu sync.Mutex

type OpRepo struct {
	opMap map[string]OpCreator
}

func GetOpRepo() *OpRepo {
	if _opRepo == nil {
		fmt.Println("oprepo: no repo instance, creating one")
		mu.Lock()
		defer mu.Unlock()
		_opRepo = &OpRepo{
			opMap: make(map[string]OpCreator),
		}
	}
	return _opRepo
}

func (o *OpRepo) CreateOpByName(name string) Operator {
	fmt.Println("oprepo: creating op instance for ", name)
	return o.opMap[name]()
}

func (o *OpRepo) Register(name string, creator OpCreator) {
	mu.Lock()
	defer mu.Unlock()
	o.opMap[name] = creator
}

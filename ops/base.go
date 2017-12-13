package ops

type Operator interface {
	Forward()
	Backward()
}

type OpCreator func() Operator

type OpDesc struct {
	Name string
}

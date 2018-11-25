package main

// Doc: https://design-patterns.readthedocs.io/zh_CN/latest/structural_patterns/facade.html

type OperationA interface {
	doA()
}
type OperationB interface {
	doB()
}
type OperationC interface {
	doC()
}

type Facade struct {
	a OperationA
	b OperationB
	c OperationC
}

func (f *Facade) operation() {
	f.a.doA()
	f.b.doB()
	f.c.doC()
}

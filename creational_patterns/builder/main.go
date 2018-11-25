package main

// Actor:
// 1. Builder
// 2. ConcreteBuilder
// 3. Director
// 4. Product

type Product interface{}

type Builder interface {
	buildPartA()
	buildPartB()
	buildPartC()
	getResult() Product
}

type BuilderA struct{}

func (b *BuilderA) buildPartA() {}
func (b *BuilderA) buildPartB() {}
func (b *BuilderA) buildPartC() {}
func (b *BuilderA) getResult() Product {}

type Director struct {
	builder Builder
}

func (d *Director) construct() Product {
	d.builder.buildPartA()
	d.builder.buildPartB()
	d.builder.buildPartC()
	return d.builder.getResult()
}

func (d *Director) setBuilder(builder Builder) {
	d.builder = builder
}

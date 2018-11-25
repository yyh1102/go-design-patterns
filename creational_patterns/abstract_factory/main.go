package main

// Actor:
// 1. AbstractFactory
// 2. ConcreteFactory
// 3. AbstractFactory
// 4. Product

type ProductA interface {
	use()
}

type ProductB interface {
	eat()
}

type Factory interface {
	createProductA() ProductA
	createProductB() ProductB
}

type FactoryOne struct{}

func (f *FactoryOne) createProductA() ProductA {}
func (f *FactoryOne) createProductB() ProductB {}

type FactoryTwo struct{}

func (f *FactoryTwo) createProductA() ProductA {}
func (f *FactoryTwo) createProductB() ProductB {}

package main

import "fmt"

// Actor:
// 1. Product
// 2. Factory
// 3. ConcreteProduct
// 4. ConcreteFactory

type Product interface {
	work()
}

type ProductA struct{}

func (p *ProductA) work() {
	fmt.Println("product A")
}

type ProductB struct{}

func (p *ProductB) work() {
	fmt.Println("product B")
}

type Factory interface {
	create() Product
}

type FactoryA struct{}

func (f *FactoryA) create() Product {
	return &ProductA{}
}

type FactoryB struct{}

func (f *FactoryB) create() Product {
	return &ProductB{}
}

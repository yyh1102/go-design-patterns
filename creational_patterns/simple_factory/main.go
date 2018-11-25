package main

import "fmt"

// Actor:
// 1. Factory
// 2. Product
// 3. ConcreteProduct

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

type Factory struct{}

func (f *Factory) createProduct(pname string) Product {
	if pname == "A" {
		return &ProductA{}
	} else if pname == "B" {
		return &ProductB{}
	}
	return nil
}

func main() {
	factory := &Factory{}
	pa := factory.createProduct("A")
	pb := factory.createProduct("B")
	pa.work()
	pb.work()
}

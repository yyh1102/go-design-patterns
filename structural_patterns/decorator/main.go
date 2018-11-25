package main

import "fmt"

// Doc: https://design-patterns.readthedocs.io/zh_CN/latest/structural_patterns/decorator.html
// Actor:
// 1. Component
// 2. ConcreteComponent
// 3. Decorator
// 4. ConcreteDecorator

type Transform interface {
	move()
}

type Car struct{}

func (c *Car) move() {
	fmt.Println("a car move")
}

type Robot struct {
	transform Transform
}

func (r *Robot) move() {
	r.transform.move()
	// Add other behaviours
}

type Airplane struct {
	transform Transform
}

func (a *Airplane) move() {
	a.transform.move()
	// Add other behaviours
}

package main

import "fmt"

// Doc: https://design-patterns.readthedocs.io/zh_CN/latest/structural_patterns/bridge.html
// Actor:
// 1. Abstraction
// 2. RefinedAbstraction
// 3. Implementor
// 4. ConcreteImplementor

type Shape interface {
	typ() string
}

type Color interface {
	rgb() string
}

type Panel interface {
	draw()
}

type PanelImpl struct {
	shape Shape
	color Color
}

func NewPanel(shape Shape, color Color) Panel {
	return &PanelImpl{
		shape: shape,
		color: color,
	}
}

func (p *PanelImpl) draw() {
	fmt.Printf("draw a %v with rgb %v", p.shape.typ(), p.color.rgb())
}

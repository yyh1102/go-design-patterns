package main

import "fmt"

// Docs: https://design-patterns.readthedocs.io/zh_CN/latest/behavioral_patterns/strategy.html
// Actor:
// 1. Context
// 2. Strategy
// 3. ConcreteStrategy

type Strategy interface {
	Algorithm()
}

type Context interface {
	Algorithm()
	SetStrategy(Strategy)
}

type StrategyA struct{}

func (s *StrategyA) Algorithm() {
	fmt.Println("strategy A")
}

type StrategyB struct{}

func (s *StrategyB) Algorithm() {
	fmt.Println("strategy B")
}

type ContextImpl struct {
	strategy Strategy
}

func (c *ContextImpl) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *ContextImpl) Algorithm() {
	c.strategy.Algorithm()
}

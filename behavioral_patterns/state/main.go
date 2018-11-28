package main

// Docs: https://design-patterns.readthedocs.io/zh_CN/latest/behavioral_patterns/state.html
// Actor:
// 1. Context
// 2. State
// 3. ConcreteState

type State interface {
	Handle(Context)
}

type StateA struct{}

func (s *StateA) Handle(ctx Context) {
	ctx.ChangeState(&StateB{})
}

type StateB struct{}

func (s *StateB) Handle(ctx Context) {
	ctx.ChangeState(&StateA{})
}

type Context interface {
	ChangeState(state State)
	Request()
}

type ContextImpl struct {
	state State
}

func (c *ContextImpl) ChangeState(state State) {
	c.state = state
}

func (c *ContextImpl) Request() {
	c.state.Handle(c)
}

func main() {

}

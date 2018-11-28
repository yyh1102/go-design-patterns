package mediator

import "fmt"

// Docs: https://design-patterns.readthedocs.io/zh_CN/latest/behavioral_patterns/mediator.html#
// Actor:
// 1. Mediator
// 2. ConcreteMediator
// 3. Colleague
// 4. ConcreteColleague

type Mediator interface {
	Register(who int, colleague Colleague)
	Operation(who int, msg string)
}

type Colleague interface {
	Send(who int, msg string)
	Receive(msg string)
}

type MediatorImpl struct {
	set map[int]Colleague
}

func NewMediator() *MediatorImpl {
	return &MediatorImpl{
		set: make(map[int]Colleague),
	}
}

func (m *MediatorImpl) Register(who int, colleague Colleague) {
	m.set[who] = colleague
}

func (m *MediatorImpl) Operation(who int, msg string) {
	w, ok := m.set[who]
	if !ok {
		return
	}
	w.Receive(msg)
}

type ColleagueImpl struct {
	mediator Mediator
}

func (c *ColleagueImpl) Send(who int, msg string) {
	fmt.Printf("Send message %s to %d\n", msg, who)
	c.mediator.Operation(who, msg)
}

func (c *ColleagueImpl) Receive(msg string) {
	fmt.Printf("Receive message %s\n", msg)
}

package main

import "fmt"

// Docs: https://design-patterns.readthedocs.io/zh_CN/latest/behavioral_patterns/command.html
// Actor:
// 1. Command
// 2. ConcreteCommand
// 3. Invoker
// 4. Receiver
// 5. Client

type Command interface {
	Execute()
}

type Receiver interface {
	Action()
}

type Invoker interface {
	Call()
}

type EatCommand struct {
	receiver Receiver
}

func (c *EatCommand) Execute() {
	fmt.Println("execute command eat")
}

type Xiaoming struct{}

func (x *Xiaoming) Action() {
	fmt.Println("receiver action")
}

type Xiaohong struct {
	command Command
}

func (x *Xiaohong) setCommand(command Command) {
	x.command = command
}

func (x *Xiaohong) Call() {
	fmt.Println("invoker call")
	x.command.Execute()
}

func main() {
	receiver := &Xiaoming{}
	command := &EatCommand{receiver}
	invoker := &Xiaohong{command}
	invoker.Call()
}

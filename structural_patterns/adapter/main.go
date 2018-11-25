package main

// Doc: https://design-patterns.readthedocs.io/zh_CN/latest/structural_patterns/adapter.html
// Actor:
// 1. Target
// 2. Adapter
// 3. Adaptee

type Target interface {
	request()
}

type Adaptee struct{}

func (a *Adaptee) specificRequest() {}

type Adapter struct {
	adaptee *Adaptee
}

func NewAdapter(adaptee *Adaptee) *Adapter {
	return &Adapter{adaptee}
}

func (a *Adapter) request() {
	a.adaptee.specificRequest()
}

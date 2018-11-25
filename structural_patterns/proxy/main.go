package main

// Doc: https://design-patterns.readthedocs.io/zh_CN/latest/structural_patterns/proxy.html
// Actor:
// 1. Subject
// 2. Proxy
// 3. RealSubject

type Subject interface {
	request()
}

type RealSubject struct{}

func (r *RealSubject) request() {}

type Proxy struct {
	subject Subject
}

func (p *Proxy) request() {
	p.beforeRequest()
	p.subject.request()
	p.afterRequest()
}

func (p *Proxy) beforeRequest() {}
func (p *Proxy) afterRequest()  {}

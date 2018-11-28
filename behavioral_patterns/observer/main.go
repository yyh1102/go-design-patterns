package observer

// Docs: https://design-patterns.readthedocs.io/zh_CN/latest/behavioral_patterns/observer.html
// Actor:
// 1. Subject
// 2. ConcreteSubject
// 3. Observer
// 4. ConcreteObserver

type Observer interface {
	Subscribe(...interface{}) Subscription
	Post(interface{})
}

type Subscription interface {
	Chan() chan interface{}
	Unsubscribe()
}

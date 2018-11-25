package main

// Docs: https://design-patterns.readthedocs.io/zh_CN/latest/structural_patterns/flyweight.html
// Actor:
// 1. Flyweight
// 2. ConcreteFlyweight
// 3. UnsharedConcreteFlyweight
// 4. FlyweightFactory

type Flyweight interface {
	operation()
}

type ConcreteFlyweight struct{}

func (c *ConcreteFlyweight) operation() {}

type Factory interface {
	getObject(string) Flyweight
}

type ObjectsPool struct {
	objects map[string]Flyweight
}

func NewObjectsPool() Factory {
	return &ObjectsPool{
		objects: make(map[string]Flyweight),
	}
}

func (pool *ObjectsPool) getObject(key string) Flyweight {
	if obj, ok := pool.objects[key]; ok {
		return obj
	}

	pool.objects[key] = &ConcreteFlyweight{}
	return pool.objects[key]
}

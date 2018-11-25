package main

var instance *Singleton

type Singleton struct{}

func NewSingleton() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}

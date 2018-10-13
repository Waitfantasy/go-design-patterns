package main

type IteratorInterface interface {
	HasNext() bool
	Next() interface{}
}

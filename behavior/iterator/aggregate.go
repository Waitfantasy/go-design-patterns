package main

type AggregateInterface interface {
	Iterator() IteratorInterface
}

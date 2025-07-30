package main

type Queue struct {
	Size uint
}

type Queuer interface {
	Queue(item any) bool
	Unqueue() any
}

package main

import "fmt"

var ProtocolCollections = map[string]func(node *Node) Protocoler{
	"osc":       NewOSCProtocol,
	"http":      NewHttpProtocol,
	"websocket": NewWebsocketProtocol,
}

type Message struct {
	ID string
}

type Protocol struct {
	Name     string
	Node     *Node
	IsInput  bool
	IsOutput bool
	Callback func(...any)
}

func (proto *Protocol) Send(args any) {
	fmt.Println("base struct protocol does not implement Send method")
}

func (proto *Protocol) CanInput() bool {
	return proto.IsInput
}

func (proto *Protocol) CanOutput() bool {
	return proto.IsOutput
}

type Protocoler interface {
	Send(any)
	CanInput() bool
	CanOutput() bool
	GetProtocol() *Protocol
}

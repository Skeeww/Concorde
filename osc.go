package main

import (
	"fmt"
)

type OSCProtocol struct {
	*Protocol
}

func (proto *OSCProtocol) Send(args ...any) {
	fmt.Println("send osc", args)
}

func NewOSCProtocol(node *Node) Protocoler {
	return &OSCProtocol{
		Protocol: &Protocol{
			Name:     "osc",
			Node:     node,
			IsInput:  false,
			IsOutput: true,
			Callback: func(a ...any) {},
		},
	}
}

func (proto *OSCProtocol) GetProtocol() *Protocol {
	return proto.Protocol
}

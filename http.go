package main

import (
	"fmt"
	"net/http"
)

type HttpProtocol struct {
	*Protocol
	Client *http.Client
}

func NewHttpProtocol(node *Node) Protocoler {
	return &HttpProtocol{
		Protocol: &Protocol{
			Name:     "http",
			Node:     node,
			IsInput:  false,
			IsOutput: true,
			Callback: func(a ...any) {},
		},
		Client: http.DefaultClient,
	}
}

func (proto *HttpProtocol) Send(args any) {
	httpMethod, ok := args.(string)
	if !ok {
		fmt.Println("missing HTTP method, can't send the message")
		return
	}

	switch httpMethod {
	case "GET":
		fmt.Println("Send HTTP GET")
	case "POST":
		fmt.Println("Send HTTP POST")
	default:
		fmt.Println("Unsupported method", httpMethod)
	}
}

func (proto *HttpProtocol) GetProtocol() *Protocol {
	return proto.Protocol
}

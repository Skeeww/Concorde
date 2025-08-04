package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebsocketProtocol struct {
	*Protocol
	Client *websocket.Conn
}

func NewWebsocketProtocol(node *Node) Protocoler {
	socket, _, err := websocket.DefaultDialer.Dial(node.Address, http.Header{})
	if err != nil {
		panic(err)
	}

	protocol := &WebsocketProtocol{
		Protocol: &Protocol{
			Name:     "websocket",
			Node:     node,
			IsInput:  true,
			IsOutput: true,
			Callback: func(a ...any) {},
		},
		Client: socket,
	}

	go func() {
		defer func() {
			if err := socket.Close(); err != nil {
				fmt.Println("socket close errored", err)
			}
		}()

		for {
			select {
			case <-protocol.Node.Context.Done():
				return
			default:
				msgType, raw, err := socket.ReadMessage()
				if err != nil {
					fmt.Println("socket error on read json", err)
					continue
				}
				if msgType != 1 {
					continue
				}
				protocol.Callback(string(raw))
			}
		}
	}()

	return protocol
}

func (proto *WebsocketProtocol) Send(args any) {
	if err := proto.Client.WriteJSON(args); err != nil {
		fmt.Println("error on writeJSON", err)
		return
	}
}

func (proto *WebsocketProtocol) GetProtocol() *Protocol {
	return proto.Protocol
}

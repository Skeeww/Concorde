package protocols

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebsocketProtocol struct {
	Context context.Context
	Address string
	Client  *websocket.Conn
}

func NewWebsocketProtocol(params ProtocolerParameters) Protocoler {
	socket, _, err := websocket.DefaultDialer.Dial(params.Address, http.Header{})
	if err != nil {
		panic(err)
	}

	protocol := WebsocketProtocol{
		Context: params.Context,
		Address: params.Address,
		Client:  socket,
	}

	go func() {
		defer func() {
			if err := socket.Close(); err != nil {
				fmt.Println("socket close errored", err)
			}
		}()

		for {
			select {
			case <-protocol.Context.Done():
				return
			default:
				msg := new(map[string]any)
				if err := socket.ReadJSON(msg); err != nil {
					fmt.Println("socket error on read json", err)
					continue
				}
				protocol.OnMessage(msg)
			}
		}
	}()

	return protocol
}

func (proto WebsocketProtocol) CanInput() bool {
	return true
}

func (proto WebsocketProtocol) CanOutput() bool {
	return true
}

func (proto WebsocketProtocol) Send(args ...any) {
	if len(args) == 0 {
		return
	}

	if err := proto.Client.WriteJSON(args); err != nil {
		fmt.Println("error on writeJSON", err)
		return
	}
}

func (proto WebsocketProtocol) OnMessage(args ...any) {
	fmt.Println("message received", args)
}

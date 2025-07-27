package protocols

import "context"

var ProtocolCollections = map[string]func(params ProtocolerParameters) Protocoler{
	"osc":       NewOSCProtocol,
	"http":      NewHttpProtocol,
	"websocket": NewWebsocketProtocol,
}

type Message struct {
	ID string
}

type Protocoler interface {
	Send(...any)
	OnMessage(...any)
	CanInput() bool
	CanOutput() bool
}

type ProtocolerParameters struct {
	Context context.Context
	Address string
}

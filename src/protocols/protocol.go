package protocols

type Message struct {
	ID string
}

type Protocoler interface {
	Send()
	OnMessage(func())
}

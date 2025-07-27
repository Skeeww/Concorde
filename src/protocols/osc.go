package protocols

import "fmt"

type OSCProtocol struct {
	Address string
}

func (proto OSCProtocol) CanInput() bool {
	return false
}

func (proto OSCProtocol) CanOutput() bool {
	return true
}

func (proto OSCProtocol) Send(args ...any) {
	fmt.Println("send", args)
}

func (proto OSCProtocol) OnMessage(args ...any) {

}

func NewOSCProtocol(params ProtocolerParameters) Protocoler {
	return OSCProtocol{
		Address: params.Address,
	}
}

package protocols

import (
	"fmt"
	"net/http"
)

type HttpProtocol struct {
	Address string
	Client  *http.Client
}

func NewHttpProtocol(params ProtocolerParameters) Protocoler {
	return HttpProtocol{
		Address: params.Address,
		Client:  http.DefaultClient,
	}
}

func (proto HttpProtocol) CanInput() bool {
	return false
}

func (proto HttpProtocol) CanOutput() bool {
	return true
}

func (proto HttpProtocol) Send(args ...any) {
	if len(args) == 0 {
		return
	}

	httpMethod, ok := args[0].(string)
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

func (proto HttpProtocol) OnMessage(args ...any) {
	fmt.Println("http protocol does not support OnMessage callback")
}

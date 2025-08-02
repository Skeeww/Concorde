package main

import (
	"net"

	"github.com/go-viper/mapstructure/v2"
)

type OSCProtocol struct {
	*Protocol
	Client net.Conn
}

type OSCAction struct {
	Address string         `mapstructure:"address"`
	Args    map[string]any `mapstructure:",remain"`
}

func (proto *OSCProtocol) Send(args any) {
	action := new(OSCAction)
	if err := mapstructure.Decode(args, action); err != nil {
		panic(err)
	}

	// message := osc.NewOSCMessage(action.Address)
	// for k, v := range action.Args {
	// 	switch k {
	// 	case "int32":
	// 		message.WithInt32(v.(int32))
	// 	}
	// }
	// payload, err := message.MarshalBinary()
	// if err != nil {
	// 	panic(err)
	// }
	// proto.Client.Write(payload)
}

func NewOSCProtocol(node *Node) Protocoler {
	udpSocket, err := net.Dial("udp", node.Address)
	if err != nil {
		panic(err)
	}

	return &OSCProtocol{
		Protocol: &Protocol{
			Name:     "osc",
			Node:     node,
			IsInput:  false,
			IsOutput: true,
			Callback: func(a ...any) {},
		},
		Client: udpSocket,
	}
}

func (proto *OSCProtocol) GetProtocol() *Protocol {
	return proto.Protocol
}

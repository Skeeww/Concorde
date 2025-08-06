package main

import (
	"errors"
	"fmt"
	"net"

	"github.com/Skeeww/Concorde/internal/osc"
)

type OSCProtocol struct {
	*Protocol
	Client net.Conn
}

func CreateOSCMessageFromAction(args any) (*osc.OSCMessage, error) {
	action, ok := args.(map[string]any)
	if !ok {
		return nil, errors.New("cannot cast OSC Action args to map")
	}

	address, ok := action["address"].(string)
	if !ok {
		return nil, errors.New("cannot find address of type string in OSC Action args")
	}
	message := osc.NewOSCMessage(address)

	arguments, ok := action["args"].([]any)
	if ok {
		for _, argument := range arguments {
			argumentMap := argument.(map[string]any)
			switch argumentMap["type"].(string) {
			case "int32":
				message.WithInt32(int32(argumentMap["value"].(int)))
			case "float32":
				message.WithFloat32(float32(argumentMap["value"].(float64)))
			case "string":
				message.WithString(argumentMap["value"].(string))
			default:
				logger.Println("Unknown OSC argument of type", argumentMap["type"])
			}
		}
	}

	return message, nil
}

func (proto *OSCProtocol) Send(args any) {
	message, err := CreateOSCMessageFromAction(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer, err := message.MarshalBinary()
	if err != nil {
		fmt.Println(err)
		return
	}

	proto.Client.Write(buffer)
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

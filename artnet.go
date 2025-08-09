package main

import (
	"context"
	"net"

	"github.com/jsimonetti/go-artnet"
)

type ArtNetProtocol struct {
	*Protocol
}

func (proto *ArtNetProtocol) Send(args any) {
}

func NewArtNetProtocol(ctx context.Context, node *Node) Protocoler {
	artsubnet := "2.0.0.0/8"
	_, cidrnet, _ := net.ParseCIDR(artsubnet)
	unicastAddrs, _ := net.InterfaceAddrs()
	var ip net.IP
	for _, addr := range unicastAddrs {
		logger.Println("Unicast address found:", addr.String())
		ip = addr.(*net.IPNet).IP
		if cidrnet.Contains(ip) {
			logger.Println("Will be using the following unicast addr:", ip.String())
			break
		}
	}

	log := artnet.NewDefaultLogger()
	c := artnet.NewController(node.Name, ip, log)
	if err := c.Start(); err != nil {
		panic(err)
	}

	return &ArtNetProtocol{
		Protocol: &Protocol{
			Name:     "artnet",
			Node:     node,
			IsInput:  false,
			IsOutput: true,
			Callback: func(a ...any) {},
		},
	}
}

func (proto *ArtNetProtocol) GetProtocol() *Protocol {
	return proto.Protocol
}

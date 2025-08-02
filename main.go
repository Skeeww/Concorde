package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config, err := LoadConfigurationFromYAMLFile("/Users/noan/Desktop/Dev/Concorde/config.yaml")
	if err != nil {
		panic(err)
	}

	nodes, err := ParseNodesFromYAML(ctx, config)
	if err != nil {
		panic(err)
	}

	links, err := ParseLinksFromYAML(config, nodes)
	if err != nil {
		panic(err)
	}

	for _, node := range nodes {
		fmt.Println("Find node", node.Name, "address:", node.Address)
	}

	for _, link := range links {
		fmt.Println("link", link.Input.Name, "on", link.Conditions)
		for _, action := range link.Actions {
			fmt.Printf("[%s]-->[%s] send %s", link.Input.Name, action.Output.Name, action.Action)
		}

		link.Input.Protocol.GetProtocol().Callback = func(a ...any) {
			for _, cond := range link.Conditions {
				if !cond.Eval(a[0]) {
					return
				}
			}
			for _, action := range link.Actions {
				action.Output.Protocol.Send(action.Action)
			}
		}
	}

	select {}
}

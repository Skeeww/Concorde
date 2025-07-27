package main

import "fmt"

func main() {
	config, err := LoadConfigurationFromYAMLFile("/Users/noan/Desktop/Dev/Concorde/config.yaml")
	if err != nil {
		panic(err)
	}

	nodes, err := ParseNodesFromYAML(config)
	if err != nil {
		panic(err)
	}

	links, err := ParseLinksFromYAML(config, nodes)
	if err != nil {
		panic(err)
	}

	for _, node := range nodes {
		fmt.Println("Find node", node.Name, "address:", node.Address)
		node.Protocol.OnMessage(func() {
			fmt.Println("Hello!")
		})
	}

	for _, link := range links {
		fmt.Println("link", link.Input.Name, "on", link.Conditions)
		for _, actions := range link.Actions {
			fmt.Printf("[%s]-->[%s] send %s", link.Input.Name, actions.Output.Name, actions.Action)
		}
	}

	select {}
}

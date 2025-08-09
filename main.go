package main

import (
	"context"
	"log"
	"os"
	"os/signal"
)

const (
	logFile = "log.txt"
)

var (
	logger = log.Default()
)

func main() {
	logger.SetFlags(log.Ldate | log.Ltime | log.Lmsgprefix)
	logger.SetOutput(os.Stdout)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx, _ = signal.NotifyContext(ctx, os.Interrupt)

	config, err := LoadConfigurationFromYAMLFile("config.yaml")
	if err != nil {
		logger.Fatal(err)
	}

	nodes, err := ParseNodesFromYAML(ctx, config)
	if err != nil {
		logger.Fatal(err)
	}

	links, err := ParseLinksFromYAML(config, nodes)
	if err != nil {
		logger.Fatal(err)
	}

	for _, node := range nodes {
		logger.Println("Find node", node.Name, "address is", node.Address)
	}

	for _, link := range links {
		logger.Println("Find link", link.Input.Name, "triggered on", link.Conditions)

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

	logger.Println("Concorde has lift-off successfully")
	<-ctx.Done()
	cancel()
	logger.Println("Concorde is landing")
}

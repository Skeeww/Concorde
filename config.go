package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Node struct {
	Context  context.Context
	Name     string
	Protocol Protocoler
	Address  string
}

type NodesCollection map[string]*Node

type Action struct {
	Output *Node
	Action []any
}

type Condition struct {
	Type  string `yaml:"type"`
	Value string `yaml:"value"`
}

func (cond *Condition) Eval(val any) bool {
	switch cond.Type {
	case "contains":
		if valToCheck, ok := val.(string); ok {
			return strings.Contains(valToCheck, cond.Value)
		}
	}
	return false
}

type Link struct {
	Input      *Node
	Conditions []Condition
	Actions    []*Action
}

type Config struct {
	Nodes NodesCollection
	Links []*Link
}

type YamlConfig struct {
	Nodes []struct {
		Name     string `yaml:"name"`
		Protocol string `yaml:"protocol"`
		Address  string `yaml:"address"`
	} `yaml:"nodes"`
	Links []struct {
		Input      string      `yaml:"input"`
		Conditions []Condition `yaml:"conditions"`
		Actions    []struct {
			Output string `yaml:"output"`
			Action []any  `yaml:"action"`
		} `yaml:"actions"`
	} `yaml:"links"`
}

func LoadConfigurationFromYAMLFile(filePath string) (*YamlConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	config := new(YamlConfig)
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

func ParseNodesFromYAML(ctx context.Context, yamlConfig *YamlConfig) (NodesCollection, error) {
	nodes := make(NodesCollection)

	for _, yamlNode := range yamlConfig.Nodes {
		node := &Node{
			Context: ctx,
			Name:    yamlNode.Name,
			Address: yamlNode.Address,
		}

		protocolInstanciateFunction, ok := ProtocolCollections[yamlNode.Protocol]
		if !ok {
			return nil, fmt.Errorf("can't find protocol named \"%s\"", yamlNode.Protocol)
		}
		node.Protocol = protocolInstanciateFunction(node)

		nodes[yamlNode.Name] = node
	}

	return nodes, nil
}

func ParseLinksFromYAML(yamlConfig *YamlConfig, nodes NodesCollection) ([]*Link, error) {
	links := make([]*Link, 0)

	for _, yamlLink := range yamlConfig.Links {
		inputNode, ok := nodes[yamlLink.Input]
		if !ok {
			fmt.Println("a link has an unknown input node", yamlLink.Input)
			continue
		}
		if !inputNode.Protocol.CanInput() {
			fmt.Println("a link has an unsupported input node", yamlLink.Input)
			continue
		}

		link := &Link{
			Input:      inputNode,
			Conditions: yamlLink.Conditions,
		}

		for _, yamlAction := range yamlLink.Actions {
			outputNode, ok := nodes[yamlAction.Output]
			if !ok {
				fmt.Println("a link has an unknown output node")
				continue
			}
			if !outputNode.Protocol.CanOutput() {
				fmt.Println("a link has an unsupported output node", yamlAction.Output)
				continue
			}

			link.Actions = append(link.Actions, &Action{
				Output: outputNode,
				Action: yamlAction.Action,
			})
		}

		links = append(links, link)
	}

	return links, nil
}

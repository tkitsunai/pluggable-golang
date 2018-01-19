package main

import (
	"github.com/dullgiulio/pingo"
)

type HelloPlugin struct{}

func (p *HelloPlugin) Say(arg []string, response *string) error {
	*response = "Hello, " + arg[0]
	return nil
}

func main() {
	plugin := &HelloPlugin{}
	pingo.Register(plugin)
	pingo.Run()
}

package main

import (
	"github.com/dullgiulio/pingo"
)

type HelloPlugin struct{}

func (p *HelloPlugin) Say(name string, msg *string) error {
	*msg = "Hello, " + name
	return nil
}

func main() {
	plugin := &HelloPlugin{}
	pingo.Register(plugin)
	pingo.Run()
}

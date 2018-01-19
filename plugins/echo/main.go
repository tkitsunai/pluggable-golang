package main

import (
	"github.com/dullgiulio/pingo"
	"strings"
)

type EchoPlugin struct{}

func (p *EchoPlugin) Echo(args []string, msg *string) error {
	*msg = strings.Join(args, " ")
	return nil
}

func main() {
	plugin := &EchoPlugin{}
	pingo.Register(plugin)
	pingo.Run()
}

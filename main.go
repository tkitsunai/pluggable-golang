package main

import (
	"fmt"
	"github.com/dullgiulio/pingo"
	"github.com/tkitsunai/pacman/protocol"
)

func main() {
	p := pingo.NewPlugin(protocol.TCP.String(), "plugins/hello-world/hello-world")
	p.Start()
	defer p.Stop()

	var res string

	err := p.Call("HelloPlugin.Say", "tkitsunai", &res)

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
package main

import (
	"flag"
	"fmt"
	"github.com/dullgiulio/pingo"
	"github.com/tkitsunai/pluggable-golang/conf"
	"github.com/tkitsunai/pluggable-golang/global"
	"github.com/tkitsunai/pluggable-golang/protocol"
)

func init() {
	confPath := flag.String("f", "", "configuration file for plugins")
	flag.Parse()
	conf.LoadConfig(*confPath, &global.ConfigPlugins)
}

func main() {
	RunPlugins()
}

func RunPlugins() {
	for _, plugin := range global.ConfigPlugins.UsePlugins {
		run(plugin)
	}
}

func run(plugin conf.Plugin) {
	path := "plugins/" + plugin.Path + "/" + plugin.BinName
	p := pingo.NewPlugin(protocol.TCP.String(), path)
	p.Start()
	defer p.Stop()

	var response string
	fmt.Println("params:", plugin.Params)
	err := p.Call(plugin.GetMethod(), plugin.Params, &response)

	if err != nil {
		panic(err)
	}

	fmt.Println("RPC Response", response)
}

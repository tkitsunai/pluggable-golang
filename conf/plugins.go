package conf

type Plugins struct {
	UsePlugins []Plugin `yaml:"plugins"`
}

type Plugin struct {
	BinName    string   `yaml:"bin_name"`
	Path       string   `yaml:"path"`
	PluginName string   `yaml:"plugin_name"`
	Method     string   `yaml:"method"`
	Params     []string `yaml:"params"`
}

func (p *Plugin) GetMethod() string {
	return p.PluginName + "." + p.Method
}

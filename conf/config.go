package conf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func LoadConfig(filePath string, plugins *Plugins) error {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	fmt.Printf("buf: %+v\n", string(buf))
	err = yaml.Unmarshal(buf, plugins)
	if err != nil {
		return err
	}
	return nil
}

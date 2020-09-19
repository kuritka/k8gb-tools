package config

import (
	"fmt"
)

type ConfigFactory struct {

}

func NewConfigFactory(configPath, gslb string) (f *ConfigFactory,err  error) {
	f = new(ConfigFactory)
	if validateConfigPath(configPath) == nil {
		var config *yamlConfig
		if config,err = newConfig(configPath); err != nil {
			return
		}
		if err = config.validate(); err != nil {
			return
		}
		for _, path := range config.K8gbTools.ConfigPaths {
			fmt.Println(path)
		}
		return
	}
	//config yaml is not provided
	//fmt.Println(gslb)
	return
}

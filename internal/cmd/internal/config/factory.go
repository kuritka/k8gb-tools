package config

import "k8s.io/cli-runtime/pkg/genericclioptions"

//GetConfig provides valid configuration
func GetConfig(configPath, gslb string) (config Config,err  error) {
	if configPath == "" {
		configFlags := genericclioptions.NewConfigFlags(true)
		path := *configFlags.KubeConfig
		config = Config{}
		config.K8gbTools.Name = gslb
		config.K8gbTools.ConfigPaths = []string{path}
		return
	}
	if validateConfigPath(configPath) == nil {
		if config,err = newConfig(configPath); err != nil {
			return
		}
		if err = config.validate(); err != nil {
			return
		}
		return
	}
	return
}

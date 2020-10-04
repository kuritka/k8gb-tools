package config

import (
	"fmt"
	"os"

	"github.com/kuritka/k8gb-tools/pkg/common/guard"

	"gopkg.in/yaml.v2"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

//GetConfig provides valid configuration or returns error
func GetConfig(configPath, gslb string) (config Config, err error) {
	if configPath == "" {
		configFlags := genericclioptions.NewConfigFlags(true)
		path := *configFlags.KubeConfig
		config = Config{}
		config.K8gbTools.Name = gslb
		config.K8gbTools.ConfigPaths = []string{path}
		return
	}
	if validateConfigPath(configPath) == nil {
		if config, err = newConfig(configPath); err != nil {
			return
		}
		if err = config.validate(); err != nil {
			return
		}
		return
	}
	return
}

// newConfig returns a new decoded Config struct
func newConfig(configPath string) (config Config, err error) {
	// Create config structure
	config = Config{}
	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return
	}
	defer close(file)

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err = d.Decode(&config); err != nil {
		return
	}
	return
}

func close(file *os.File) {
	err := file.Close()
	if err != nil {
		guard.FailOnError(err, "unable to close file %s", file.Name())
	}
}

// validateConfigPath just makes sure, that the path provided is a file,
// that can be read
func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

func (c *Config) validate() (err error) {
	if c.K8gbTools.Name == "" {
		return fmt.Errorf("yaml configuration: missing gslb-name")
	}
	if len(c.K8gbTools.ConfigPaths) == 0 {
		return fmt.Errorf("yaml configuration: missing ConfigPaths")
	}
	for _, path := range c.K8gbTools.ConfigPaths {
		if err = validateConfigPath(path); err != nil {
			return fmt.Errorf("yaml configuration: %s", err.Error())
		}
	}
	return
}

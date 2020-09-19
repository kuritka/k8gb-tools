package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

//#config.yaml
//
//k8gb-tools:
//	gslb-name: test-gslb
//	config:
//	- test-gslb1
//	- test-gslb2

//Config struct for k8gb-tools
type yamlConfig struct {
	//K8gbTools root structure
	K8gbTools struct {
		//Name of Gslb
		Name string `yaml:"gslb-name"`
		//ConfigPaths array of paths to configuration; if empty it uses default configuration
		ConfigPaths[] string `yaml:"config"`
	}`yaml:"k8gb-tools"`
}

// newConfig returns a new decoded Config struct
func newConfig(configPath string) (config *yamlConfig,err error) {
	// Create config structure
	config = new(yamlConfig)
	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err = d.Decode(&config); err != nil {
		return
	}
	return
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


func (c *yamlConfig) validate() (err error) {
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
package config

//Config struct for k8gb-tools
//#config.yaml
//
//k8gb-tools:
//	gslb-name: test-gslb
//	config:
//	- test-gslb1
//	- test-gslb2
type Config struct {
	//K8gbTools root structure
	K8gbTools struct {
		//Name of Gslb
		Name string `yaml:"gslb-name"`
		//ConfigPaths array of paths to configuration; if empty it uses default configuration
		ConfigPaths []string `yaml:"config"`
	} `yaml:"k8gb-tools"`
}

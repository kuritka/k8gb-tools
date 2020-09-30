package k8s

import (
	"fmt"
	"io/ioutil"

	"github.com/kuritka/k8gb-tools/internal/cmd/internal/k8s/internal/config"
	"k8s.io/client-go/dynamic"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

//KubeConfig
type KubeConfig struct {
	//RestConfig reads common resources
	RestConfig *restclient.Config
	//RawConfig
	RawConfig clientcmdapi.Config
	//DynamicConfig reads custom resources
	DynamicConfig dynamic.Interface
	//ClientConfig, keeps, raw info about cluster search for ToRawKubeConfigLoader()
	ClientConfig clientcmd.ClientConfig
	//Source
	Source string
}

type KubeConfigFactory struct {
	yaml config.Config
}

//NewKubeConfigFactory receives yaml path or gslb name and returns error if invalid
func NewKubeConfigFactory(yaml, gslb string) (factory *KubeConfigFactory, err error) {
	factory = new(KubeConfigFactory)
	factory.yaml, err = config.GetConfig(yaml, gslb)
	return
}

//GetConfig instantiate all possible configurations
func (f *KubeConfigFactory) InitializeConfigs() (configs []*KubeConfig, err error) {
	configs = make([]*KubeConfig, 0)
	for _, path := range f.yaml.K8gbTools.ConfigPaths {
		rawConfig, err := getRawConfigs(path)
		if err != nil {
			return configs, err
		}
		for ctx := range rawConfig.Contexts {
			c, err := switchContextAndReadConfigs(rawConfig, ctx)
			if err != nil {
				return configs, err
			}
			c.Source = path
			configs = append(configs, c)
		}
	}
	return
}

func getRawConfigs(kubeConfigPath string) (rawConfig clientcmdapi.Config, err error) {
	b, err := ioutil.ReadFile(kubeConfigPath)
	if err != nil {
		return
	}
	clientConfig, err := clientcmd.NewClientConfigFromBytes(b)
	if err != nil {
		return rawConfig, fmt.Errorf("reading ClientConfig from %s %s", kubeConfigPath, err)
	}
	rawConfig, err = clientConfig.RawConfig()
	if err != nil {
		return rawConfig, fmt.Errorf("create RawConfig %s", err)
	}
	return
}

func switchContextAndReadConfigs(raw clientcmdapi.Config, ctx string) (config *KubeConfig, err error) {
	config = new(KubeConfig)
	if raw.Contexts[ctx] == nil {
		return config, fmt.Errorf("context %s doesn't exists", ctx)
	}
	override := &clientcmd.ConfigOverrides{CurrentContext: ctx, Context: *raw.Contexts[ctx]}
	config.ClientConfig = clientcmd.NewNonInteractiveClientConfig(raw, ctx,
		override, &clientcmd.ClientConfigLoadingRules{})
	config.RawConfig = raw
	config.RawConfig.CurrentContext = ctx
	config.RestConfig, err = config.ClientConfig.ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("create Rest %s", err)
	}
	config.DynamicConfig, err = dynamic.NewForConfig(config.RestConfig)
	if err != nil {
		return nil, fmt.Errorf("create Dynamic %s", err)
	}
	return
}

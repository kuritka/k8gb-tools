package k8s

import (
	"fmt"
	"github.com/kuritka/k8gb-tools/internal/cmd/internal/k8s/internal/config"
	"io/ioutil"
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
		cfg, err := getConfig(path)
		if err != nil {
			return configs, err
		}
		configs = append(configs, cfg)
	}
	return
}

func getConfig(kubeConfigPath string) (config *KubeConfig, err error) {
	config = new(KubeConfig)
	b, err := ioutil.ReadFile(kubeConfigPath)
	if err != nil {
		return
	}
	config.ClientConfig, err = clientcmd.NewClientConfigFromBytes(b)
	if err != nil {
		return nil, fmt.Errorf("reading ClientConfig from %s %s", kubeConfigPath, err)
	}
	config.RawConfig, err = config.ClientConfig.RawConfig()
	if err != nil {
		return nil, fmt.Errorf("create RawConfig %s", err)
	}
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

func switchContext(cfg KubeConfig, ctx string) (err error){
	if cfg.RawConfig.Contexts[ctx] == nil {
		return fmt.Errorf("context %s doesn't exists", ctx)
	}
	override := &clientcmd.ConfigOverrides{CurrentContext: ctx, Context: *cfg.RawConfig.Contexts[ctx]}
	restConfig := clientcmd.NewNonInteractiveClientConfig(cfg.RawConfig, ctx,
		override, &clientcmd.ClientConfigLoadingRules{})
	cfg.RawConfig.CurrentContext = ctx
	cfg.RestConfig, err = restConfig.ClientConfig()
	if err != nil {
		return
	}
	cfg.DynamicConfig, err = dynamic.NewForConfig(cfg.RestConfig)
	return
}

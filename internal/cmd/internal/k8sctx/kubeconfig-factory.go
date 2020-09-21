package k8sctx

import (
	"fmt"
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


//Reads all configuration from config path
func Get(kubeConfigPath string) (config *KubeConfig, err error) {
	config = new(KubeConfig)
	b, err := ioutil.ReadFile(kubeConfigPath)
	if err != nil {
		return
	}
	config.ClientConfig, err = clientcmd.NewClientConfigFromBytes(b)
	if err != nil {
		return nil, fmt.Errorf("reading ClientConfig from %s %s",kubeConfigPath,err)
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
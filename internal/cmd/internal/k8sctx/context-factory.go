package k8sctx

import (
	"github.com/kuritka/k8gb-tools/internal/cmd/internal/k8s"
)

//ContextFactory produces k8s context
type ContextFactory struct {
	configs []*k8s.KubeConfig
}

//NewContextFactory returns context of command
func NewContextFactory(yaml, gslb string) (factory *ContextFactory, err error) {
	factory = new(ContextFactory)
	var k8sf *k8s.KubeConfigFactory
	k8sf, err = k8s.NewKubeConfigFactory(yaml, gslb)
	if err != nil {
		return
	}
	factory.configs, err = k8sf.InitializeConfigs()
	if err != nil {
		return
	}
	return
}

//List returns list of GSLBs within namespaces
func (f *ContextFactory) List() {

}

//List returns list of GSLBs within namespaces
func (f *ContextFactory) Switch() {

}

//List returns list of GSLBs within namespaces
func (f *ContextFactory) GetContext() {

}

////Get returns context
//func (cf *ContextFactory) Get() (*Context, error) {
//	var err error
//	ctx := new(Context)
//	ctx.Command = new(Command)
//	ctx.Command.Context, ctx.Command.Cancel = context.WithCancel(context.Background())
//	ctx.K8s = new(K8s)
//	ctx.K8s.kubeConfig = *configFlags.KubeConfig
//	ctx.K8s.Cluster = *configFlags.ClusterName
//	ctx.K8s.IOStreams = genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
//	ctx.K8s.RawConfig, err = configFlags.ToRawKubeConfigLoader().RawConfig()
//	if err != nil {
//		return nil, fmt.Errorf("create RawConfig %s", err)
//	}
//	ctx.K8s.ClientConfig, err = configFlags.ToRESTConfig()
//	if err != nil {
//		return nil, fmt.Errorf("create Rest %s", err)
//	}
//	ctx.K8s.DynamicConfig, err = dynamic.NewForConfig(ctx.K8s.ClientConfig)
//	if err != nil {
//		return nil, fmt.Errorf("create Dynamic %s", err)
//	}
//	ctx.K8s.ctxBackup = ctx.K8s.RawConfig.CurrentContext
//	return ctx, nil
//}

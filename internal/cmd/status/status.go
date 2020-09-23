package status

import (
	"fmt"
	"github.com/kuritka/k8gb-tools/internal/cmd/internal/config"
	"github.com/kuritka/k8gb-tools/internal/cmd/internal/k8sctx"
	"github.com/kuritka/k8gb-tools/pkg/common/guard"
)

type Status struct {
	config config.Config
}

func New(cfg config.Config) (status *Status) {
	status = new(Status)
	//ctx,err := k8sctx.NewContextFactory(cfg).Get()
	cfg, err := config.GetConfig(`/Users/ab011th/go/src/github.com/kuritka/k8gb-tools/tests/config.yaml`,"")
	guard.FailOnError(err,"reading configuration")
	load := make([]*k8sctx.KubeConfig,len(cfg.K8gbTools.ConfigPaths))
	for _, path := range cfg.K8gbTools.ConfigPaths {
		k8scfg, err := k8sctx.GetConfig(path)
		guard.FailOnError(err,"fail on reading %s",path)
		ctx,err := k8sctx.NewContextFactory(k8scfg).Get()
		guard.FailOnError(err,"fail on reading %s",path)
		load = append(load, c)
	}
//	ctx,_ := k8sctx.NewContextFactory(cfg).Get()
//	fmt.Println(ctx)
	return
}

func (s *Status) String() string{
	return "status"
}

func (s *Status) Run() (err error) {
	fmt.Println("RUN")
	return
}
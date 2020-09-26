package status

import (
	"fmt"
	"github.com/kuritka/k8gb-tools/internal/cmd/internal/config"
	"github.com/kuritka/k8gb-tools/internal/cmd/internal/k8s"
)

type Status struct {
	config config.Config
	yaml string
	gslb string
}

func New(yaml, gslb string) (status *Status) {
	status = new(Status)
	status.yaml = yaml
	status.gslb = gslb
	return
}

func (s *Status) String() string{
	return "status"
}

func (s *Status) Run() error {
	ctx,err := k8s.NewKubeConfigFactory(s.yaml,s.gslb)
	if err != nil {
		return err
	}
	cfgs,err := ctx.InitializeConfigs()
	if err != nil {
		return err
	}
	fmt.Println(cfgs[0].RestConfig.Host)
	return nil
}
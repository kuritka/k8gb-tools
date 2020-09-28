package status

import (
	"fmt"
	"github.com/kuritka/k8gb-tools/internal/cmd/internal/k8sctx"
)

type Status struct {
	yaml string
	gslb string
}

func New(yaml, gslb string) (status *Status) {
	status = new(Status)
	status.yaml = yaml
	status.gslb = gslb
	return
}

func (s *Status) String() string {
	return "status"
}

func (s *Status) Run() error {
	ctx, err := k8sctx.NewContextFactory(s.yaml, s.gslb)
	if err != nil {
		return err
	}
	list, err := ctx.List()
	if err != nil {
		return err
	}
	fmt.Println(list)
	return err
}

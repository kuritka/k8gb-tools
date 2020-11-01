package status

import (
	"github.com/kuritka/k8gb-tools/pkg/view"

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

func (s *Status) Run() (err error) {
	ctx, err := k8sctx.NewContextFactory(s.yaml, s.gslb)
	if err != nil {
		return err
	}
	model, err := ctx.GetStatus()
	if err != nil {
		return err
	}
	for _, m := range model {
		err = view.NewStatusView(m).Print()
		if err != nil {
			return
		}
	}
	return
}

package status

import (
	"fmt"
	"github.com/kuritka/k8gb-tools/internal/cmd/internal/config"
	"github.com/kuritka/k8gb-tools/internal/cmd/internal/k8sctx"
)

type Status struct {
	config config.Config
}

func New(cfg config.Config) (status *Status) {
	status = new(Status)
	ctx,err := k8sctx.NewContextFactory(cfg).Get()

	return
}

func (s *Status) String() string{
	return "status"
}

func (s *Status) Run() (err error) {
	fmt.Println("RUN")
	return
}
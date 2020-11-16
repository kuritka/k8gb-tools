package list

import (
	"errors"
	"fmt"

	"github.com/enescakir/emoji"
	"github.com/kuritka/k8gb-tools/internal/cmd/internal/k8sctx"
	"github.com/logrusorgru/aurora"
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
	return "list"
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
	if len(list) == 0 {
		return errors.New(fmt.Sprintf( `"%s" not found`, ctx.GetGSLBName()))
	}
	fmt.Printf("\n%s %s\n", emoji.FourLeafClover, aurora.BrightBlue("k8gb"))
	for _, l := range list {
		//TODO: would be in view. what is in formatting, what is here ?
		fmt.Printf("\t%s\n\t      %s\n\n", l, aurora.Cyan(l.Source))
	}
	return err
}

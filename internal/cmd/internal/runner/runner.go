package runner

import (
	"github.com/kuritka/k8gb-tools/pkg/common/guard"
)

//CmdRunner is running all commands
type CmdRunner struct {
	service ICmdRunner
}

//New creates new instance of CmdRunner
func New(command ICmdRunner) *CmdRunner {
	return &CmdRunner{
		command,
	}
}

//MustRun runs service once and panics if service is broken
func (r *CmdRunner) MustRun() {
	err := r.service.Run()
	guard.FailOnError(err, "command %s failed", r.service.String())
}

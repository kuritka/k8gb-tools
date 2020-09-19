package status

import "fmt"

type Status struct {
	opts Options
}

func New(opts Options) (status *Status) {
	status = new(Status)
	status.opts = opts
	return
}

func (s *Status) String() string{
	return "status"
}

func (s *Status) Run() (err error) {
	fmt.Println("RUN")
	return
}
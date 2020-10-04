package view

import (
	"github.com/kuritka/k8gb-tools/pkg/common/guard"
	"github.com/kuritka/k8gb-tools/pkg/model"
)

type StatusView struct {
	printer *PrettyPrinter
	model   model.Status
}

func NewStatusView(model model.Status) *StatusView {
	v := new(StatusView)
	v.printer = DefaultPrettyPrinter()
	v.model = model
	return v
}

func (v *StatusView) Print() error {
	for _, geotag := range v.model.GeoTag.Values {
		err := v.printer.Title(geotag)
		guard.FailOnError(err, "printing geotag: %s", geotag)
	}
	return nil
}

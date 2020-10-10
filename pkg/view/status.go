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
	for i, geotag := range v.model.GeoTag.Values {
		err := v.printer.Title(geotag)
		guard.FailOnError(err, "printing geotag: %s", "Geotag")
		err = v.printer.Subtitle(v.model.Name.Values[i])
		guard.FailOnError(err, "printing name: %s", "Name")
		v.printer.NewLine()
		err = v.printer.Paragraph(v.model.Type.Values[i])
		guard.FailOnError(err, "printing type: %s", "Type")
		v.printer.NewLine()
		v.printer.NewLine()
	}
	return nil
}

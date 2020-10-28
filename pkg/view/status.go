package view

import (
	"fmt"
	"github.com/kuritka/k8gb-tools/pkg/common/guard"
	"github.com/kuritka/k8gb-tools/pkg/model"
)

type StatusView struct {
	printer *PrettyPrinter
	model   model.Status
}

//NewStatusView retreives status view
func NewStatusView(model model.Status) *StatusView {
	v := new(StatusView)
	v.printer = DefaultPrettyPrinter()
	v.model = model
	return v
}


//Print prints view
func (v *StatusView) Print() error {
	for i, geotag := range v.model.GeoTag.Values {
		guard.FailOnError(v.printer.Title(geotag, v.model.Name.Values[i]), "printing geotag or name")
		v.printer.NewLine()
		guard.FailOnError(v.printer.Paragraph(v.model.Type.Property, v.model.Type.Values[i], v.model.Type.Error), "printing type")
		for _, rule := range v.model.Ingress.Rules {
			guard.FailOnError(v.printer.Paragraph("Ingress",fmt.Sprintf("%s %s %s",rule.Host, rule.IpAddress,rule.Node), nil),"printing ingress")
		}
		v.printer.NewLine()
		v.printer.NewLine()
	}
	return nil
}

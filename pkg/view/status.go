package view

import (
	"github.com/kuritka/k8gb-tools/pkg/common/guard"
	"github.com/kuritka/k8gb-tools/pkg/model"
)

type StatusView struct {
	printer *PrettyPrinter
	status  model.Status
}

//NewStatusView retrieves status view
func NewStatusView(status model.Status) *StatusView {
	v := new(StatusView)
	v.printer = DefaultPrettyPrinter()
	v.status = status
	return v
}

//Print prints view
func (v *StatusView) Print() error {
	guard.FailOnError(v.printer.Title(v.status.GeoTag, v.status.Name), "printing geotag or name")
	guard.FailOnError(v.printer.Paragraph("Type", v.status.Type, nil), "printing type")
	for _, ingress := range v.status.Ingresses {
		guard.FailOnError(v.printer.Paragraph("Ingress", ingress.Name, nil), "ingress")
		for _, rule := range ingress.Rules {
			guard.FailOnError(v.printer.NoParagraph(rule.Host, nil), "ingress")
		}
	}
	v.printer.NewLine()
	v.printer.NewLine()
	return nil
}

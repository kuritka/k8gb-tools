package view

import (
	"fmt"
	"github.com/kuritka/k8gb-tools/pkg/common/guard"
	"github.com/kuritka/k8gb-tools/pkg/model"
)

type StatusView struct {
	printer *PrettyPrinter
	status  model.Status
}

//NewStatusView retreives status view
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
			for _, backend := range rule.Backends {
				guard.FailOnError(v.printer.NoParagraph(fmt.Sprintf("%s:%v%s",backend.Service,backend.Port,backend.Path), nil), "ingress")
			}
		}
	}
	v.printer.NewLine()
	v.printer.NewLine()
	//for i, geotag := range v.model.GeoTag.Values {
	//	guard.FailOnError(v.printer.Title(geotag, v.model.Name.Values[i]), "printing geotag or name")
	//	v.printer.NewLine()
	//	guard.FailOnError(v.printer.Paragraph(v.model.Type.Property, v.model.Type.Values[i], v.model.Type.Error), "printing type")
	//	for _, rule := range v.model.Ingresses.Rules {
	//		guard.FailOnError(v.printer.Paragraph("Ingresses",fmt.Sprintf("%s %s %s",rule.Host, rule.IPAddress,rule.Node), nil),"printing ingress")
	//	}
	//	v.printer.NewLine()
	//	v.printer.NewLine()
	//}
	return nil
}

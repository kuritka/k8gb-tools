package view

import (
	"github.com/hashicorp/go-multierror"
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
	var err error
	var e *multierror.Error
	err = v.printer.Title(v.status.GeoTag, v.status.Name)
	e = multierror.Append(e, err)
	err = v.printer.Paragraph("Type", v.status.Type, nil)
	e = multierror.Append(e, err)
	for _, ingress := range v.status.Ingresses {
		err = v.printer.Paragraph("Ingress", ingress.Name, nil)
		e = multierror.Append(e, err)
		for _, rule := range ingress.Rules {
			err = v.printer.NoParagraph(rule.Host, nil)
			e = multierror.Append(e, err)
		}
	}
	v.printer.NewLine()
	v.printer.NewLine()
	return e.ErrorOrNil()
}

package k8sctx

import (
	"github.com/kuritka/k8gb-tools/pkg/model"
)

type raw struct {
	Gslb []GslbRaw
}

//NewRaw
func NewRaw() *raw {
	raw := new(raw)
	raw.Gslb = make([]GslbRaw, 0)
	return raw
}

//ValidateGeoTag
func (r *raw) ValidateGeoTag() *model.Stringr {
	stringr := model.InitStringr()
	for _, gslbRaw := range r.Gslb {
		stringr.Append(gslbRaw.Status.GeoTag)
	}
	return stringr.ValuesAreUnique().ValuesAreNotEmpty()
}

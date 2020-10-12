package k8sctx

import (
	"github.com/kuritka/k8gb-tools/pkg/common"
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
	stringr := model.InitStringr("GeoTag")
	for _, gslbRaw := range r.Gslb {
		stringr.Append(gslbRaw.Status.GeoTag)
	}
	return stringr.ValuesAreUnique().ValuesAreNotEmpty()
}

//ValidateName
func (r *raw) ValidateName() *model.Stringr {
	stringr := model.InitStringr("Name")
	for _, gslbRaw := range r.Gslb {
		stringr.Append(gslbRaw.Name)
	}
	return stringr.ValuesAreEqual().ValuesAreNotEmpty()
}

//ValidateType
func (r *raw) ValidateType() *model.Stringr {
	stringr := model.InitStringr("Type")
	for _, gslbRaw := range r.Gslb {
		stringr.Append(gslbRaw.Type)
	}
	return stringr.ValuesAreEqual().ValuesAreIn(common.Strategy[:]...)
}

////ValidateType
//func (r *raw) ValidateHost() *model.Stringr {
//	stringr := model.InitStringr()
//	for _, gslbRaw := range r.Gslb {
//		stringr.Append(gslbRaw.Rules)
//	}
//	return stringr.ValuesAreEqual().ValuesAreIn(common.Strategy[:]...)
//}

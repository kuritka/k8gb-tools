package k8sctx

import (
	"github.com/kuritka/k8gb-tools/pkg/common"
	"github.com/kuritka/k8gb-tools/pkg/model"
)

type Raw struct {
	Gslb []GslbRaw
}

//NewRaw
func NewRaw() *Raw {
	raw := new(Raw)
	raw.Gslb = make([]GslbRaw, 0)
	return raw
}

//ValidateGeoTag
func (r *Raw) ValidateGeoTag() *model.Stringr {
	stringr := model.InitStringr("GeoTag")
	for _, gslbRaw := range r.Gslb {
		stringr.Append(gslbRaw.Status.GeoTag)
	}
	return stringr.ValuesAreUnique().ValuesAreNotEmpty()
}

//ValidateName
func (r *Raw) ValidateName() *model.Stringr {
	stringr := model.InitStringr("Name")
	for _, gslbRaw := range r.Gslb {
		stringr.Append(gslbRaw.Name)
	}
	return stringr.ValuesAreEqual().ValuesAreNotEmpty()
}

//ValidateType
func (r *Raw) ValidateType() *model.Stringr {
	stringr := model.InitStringr("Type")
	for _, gslbRaw := range r.Gslb {
		stringr.Append(gslbRaw.Type)
	}
	return stringr.ValuesAreEqual().ValuesAreIn(common.Strategy[:]...)
}

func (r *Raw) ValidateIngress() *model.Ingress {
	return nil
}

////ValidateType
//func (r *Raw) ValidateHost() *model.Stringr {
//	stringr := model.InitStringr()
//	for _, gslbRaw := range r.Gslb {
//		stringr.Append(gslbRaw.Rules)
//	}
//	return stringr.ValuesAreEqual().ValuesAreIn(common.Strategy[:]...)
//}

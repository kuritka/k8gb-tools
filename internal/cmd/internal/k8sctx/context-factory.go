package k8sctx

import (
	"fmt"
	"github.com/kuritka/k8gb-tools/internal/cmd/internal/k8s"
	"github.com/kuritka/k8gb-tools/pkg/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var runtimeClassGVR = schema.GroupVersionResource{
	Group:    "k8gb.absa.oss",
	Version:  "v1beta1",
	Resource: "gslbs",
}

var emptyGslb = GslbRaw{
	Source: "",
	Status: Status{GeoTag: ""},
	Metadata: Metadata{Name: "",Namespace: ""},
	CurrentContext: "",
	Error: fmt.Errorf("no gslb in configuration"),
}


//ContextFactory produces k8s context
type ContextFactory struct {
	configs []*k8s.KubeConfig
}

//NewContextFactory returns context of command
func NewContextFactory(yaml, gslb string) (factory *ContextFactory, err error) {
	factory = new(ContextFactory)
	var k8sf *k8s.KubeConfigFactory
	k8sf, err = k8s.NewKubeConfigFactory(yaml, gslb)
	if err != nil {
		return
	}
	factory.configs, err = k8sf.InitializeConfigs()
	if err != nil {
		return
	}
	return
}

//List returns list of GSLBs within namespaces
func (f *ContextFactory) List() (m []model.ListItem, err error) {
	m = make([]model.ListItem, 0)
	raws, err := readRaw(f.configs)
	if err != nil {
		return m, err
	}
	for _, raw := range raws.Gslb {
		item := model.ListItem{
			raw.Namespace,
			raw.Name,
			raw.GeoTag,
			raw.CurrentContext,
			raw.Source,
			raw.Error,
		}
		m = append(m, item)
	}
	return m, nil
}

//List returns list of GSLBs within namespaces
func (f *ContextFactory) GetStatus() (m model.Status, err error) {
	raw, err := readRaw(f.configs)
	if err != nil {
		return m, err
	}
	m.Name = *raw.ValidateName()
	m.GeoTag = *raw.ValidateGeoTag()
	m.Type = *raw.ValidateType()
	//m.Host = *raw.ValidateHost()
	return
}

func readRaw(configs []*k8s.KubeConfig) (raw *raw, err error) {
	raw = NewRaw()
	for _, config := range configs {
		unstructuredList, err := config.DynamicConfig.Resource(runtimeClassGVR).List(metav1.ListOptions{})
		if err != nil {
			return raw, err
		}
		raw.Gslb = append(raw.Gslb, getUnstructured(unstructuredList, config)...)
	}
	return
}

//maps unstructured data into GslbRaw structure. Any CRD change has to be reflected
//in GslbRaw or underlying structures
func getUnstructured(u *unstructured.UnstructuredList, config *k8s.KubeConfig) (gslbRaws []GslbRaw) {
	gslbRaws = make([]GslbRaw, len(u.Items))
	if len(u.Items) == 0 {
		e := emptyGslb
		e.Source = config.Source
		return []GslbRaw{ e }
	}
	for i, o := range u.Items {
		d := GslbRaw{}
		d.Error = runtime.DefaultUnstructuredConverter.FromUnstructured(o.Object, &d)
		d.CurrentContext = config.RawConfig.CurrentContext
		d.Source = config.Source
		gslbRaws[i] = d
	}
	return
}

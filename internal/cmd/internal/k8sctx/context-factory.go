package k8sctx

import (
	"fmt"
	"github.com/kuritka/k8gb-tools/internal/cmd/internal/k8s"
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
func (f *ContextFactory) List() (li []ListItem, err error) {
	li = make([]ListItem, 0)
	raws, err := readRaw(f.configs)
	if err != nil {
		return li, err
	}
	for _, raw := range raws {
		item := ListItem{
			raw.Namespace,
			raw.Name,
			raw.GeoTag,
			raw.CurrentContext,
			raw.Source,
		}
		li = append(li, item)
	}
	return li, nil
}

//List returns list of GSLBs within namespaces
func (f *ContextFactory) GetContext() (err error) {
	for _, config := range f.configs {
		unstructuredList, err := config.DynamicConfig.Resource(runtimeClassGVR).List(metav1.ListOptions{})
		if err != nil {
			return err
		}
		raws := getUnstructured(unstructuredList, config)
		for _, raw := range raws {
			fmt.Println(raw)
		}
	}
	return
}

func readRaw(configs []*k8s.KubeConfig) (gslbRaws []GslbRaw, err error) {
	gslbRaws = make([]GslbRaw,0)
	for _, config := range configs {
		unstructuredList, err := config.DynamicConfig.Resource(runtimeClassGVR).List(metav1.ListOptions{})
		if err != nil {
			return gslbRaws,err
		}
		gslbRaws = append(gslbRaws, getUnstructured(unstructuredList, config)...)
	}
	return
}

//maps unstructured data into GslbRaw structure. Any CRD change has to be reflected
//in GslbRaw or underlying structures
func getUnstructured(u *unstructured.UnstructuredList, config *k8s.KubeConfig) (gslbRaws []GslbRaw) {
	gslbRaws = make([]GslbRaw, len(u.Items))
	for i, o := range u.Items {
		d := GslbRaw{}
		d.Error = runtime.DefaultUnstructuredConverter.FromUnstructured(o.Object, &d)
		d.CurrentContext = config.RawConfig.CurrentContext
		d.Source = config.Source
		gslbRaws[i] = d
	}
	return
}

package k8sctx

import (
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
func (f *ContextFactory) List() ([]ListItem,error){
	li := make([]ListItem,0)
	for _, config := range f.configs {
		unstructuredList,err := config.DynamicConfig.Resource(runtimeClassGVR).List(metav1.ListOptions{})
		if err != nil {
			return li,err
		}
		raws := getUnstructured(unstructuredList)

		for _, raw := range raws {
			item := ListItem{
				raw.Namespace,
				raw.Name,
				raw.Cluster,
				raw.GeoTag,
				config.RawConfig.CurrentContext,
			}
			li =append(li,item)
		}
	}
	return li,nil
}

//List returns list of GSLBs within namespaces
func (f *ContextFactory) Switch() {

}

//List returns list of GSLBs within namespaces
func (f *ContextFactory) GetContext() {

}

//maps unstructured data into GslbRaw structure. Any CRD change has to be reflected
//in GslbRaw or underlying structures
func getUnstructured(u *unstructured.UnstructuredList) (desc []GslbRaw) {
	desc = make([]GslbRaw, len(u.Items))
	for i, o := range u.Items {
		d := GslbRaw{}
		d.Error = runtime.DefaultUnstructuredConverter.FromUnstructured(o.Object, &d)
		desc[i] = d
	}
	return
}

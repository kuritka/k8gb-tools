package k8sctx

import (
	"fmt"

	"github.com/kuritka/k8gb-tools/internal/cmd/internal/k8s"
	"github.com/kuritka/k8gb-tools/pkg/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var runtimeClassGVR = schema.GroupVersionResource{
	Group:    "k8gb.absa.oss",
	Version:  "v1beta1",
	Resource: "gslbs",
}

var emptyGslb = GslbRaw{
	Source:         "",
	Status:         Status{GeoTag: ""},
	Metadata:       Metadata{Name: "", Namespace: ""},
	CurrentContext: "",
	Error:          fmt.Errorf("no gslb in configuration"),
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
			Namespace: raw.Namespace,
			Name:      raw.Name,
			GeoTag:    raw.GeoTag,
			Context:   raw.CurrentContext,
			Source:    raw.Source,
			Error:     raw.Error,
		}
		m = append(m, item)
	}
	return m, nil
}

//GetStatus returns gslb status across all configured contexts
func (f *ContextFactory) GetStatus() (m []model.Status, err error) {
	//Do validations and transitions here!
	m = make([]model.Status, 0)
	r, err := readRaw(f.configs)
	for _, rg := range r.Gslb {
		s := model.Status{}
		s.Host = rg.Cluster
		s.Name = rg.Name
		s.GeoTag = rg.GeoTag
		s.Type = rg.Type
		s.Namespace = rg.Namespace
		for _, ri := range rg.Ingress {
			si := model.Ingress{}
			si.Name = ri.Name
			si.Annotations = ri.Annotations

			for _, rr := range ri.Rules {
				r := model.Rule{}
				r.Host = rr.Host
				for _, rb := range rr.Backends {
					r.Backends = append(r.Backends,
						model.Backend{
						Service: rb.Service,
						Port:    rb.Port,
						Path:    rb.Path})
				}
				si.Rules = append(si.Rules, r)
			}
			s.Ingresses = append(s.Ingresses, si)
		}
		m = append(m, s)
	}
	//m.Name = *Raw.ValidateName()
	//m.GeoTag = *Raw.ValidateGeoTag()
	//m.Type = *Raw.ValidateType()
	//m.Ingresses = *Raw.ValidateIngress()
	//for _, gslb := range Raw.Gslb {
	//	for _, ingress := range gslb.Ingresses {
	//		for _, rule := range ingress.Rules {
	//
	//			m.Ingresses.Rules = append(m.Ingresses.Rules, )
	//		}
	//	}
	//m.Host = *Raw.ValidateHost()
	return m, err
}

func readRaw(configs []*k8s.KubeConfig) (raw *Raw, err error) {
	raw = NewRaw()
	for _, config := range configs {
		unstructuredList, err := config.DynamicConfig.Resource(runtimeClassGVR).List(metav1.ListOptions{})
		if err != nil {
			return raw, err
		}
		gslbs, err := getUnstructured(unstructuredList, config)
		if err != nil {
			return raw, err
		}
		raw.Gslb = append(raw.Gslb, gslbs...)
	}
	return
}

//maps unstructured data into GslbRaw structure. Any CRD change has to be reflected
//in GslbRaw or underlying structures
func getUnstructured(u *unstructured.UnstructuredList, config *k8s.KubeConfig) (gslbRaws []GslbRaw, err error) {
	gslbRaws = make([]GslbRaw, len(u.Items))
	if len(u.Items) == 0 {
		e := emptyGslb
		e.Source = config.Source
		return []GslbRaw{e}, nil
	}
	for i, o := range u.Items {
		d := GslbRaw{}
		d.Error = runtime.DefaultUnstructuredConverter.FromUnstructured(o.Object, &d)
		d.CurrentContext = config.RawConfig.CurrentContext
		d.Source = config.Source
		ns := o.GetNamespace()
		d.Ingress, err = getIngressRaw(config.RestConfig, ns)
		if err != nil {
			panic(err)
		}
		gslbRaws[i] = d
	}
	return
}

func getIngressRaw(cfg *rest.Config, namespace string) (is []IngressRaw, err error) {
	cs, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return
	}
	ings, err := cs.NetworkingV1beta1().Ingresses(namespace).List(metav1.ListOptions{})
	if err != nil {
		return
	}
	for _, ingress := range ings.Items {
		ing := new(IngressRaw)
		ing.Name = ingress.Name
		ing.Namespace = ingress.Namespace
		ing.Annotations = ingress.Annotations
		ing.Labels = ingress.Labels
		for _, lbi := range ingress.Status.LoadBalancer.Ingress {
			ing.LoadBalancers = append(ing.LoadBalancers, EndpointRaw{lbi.IP, lbi.Hostname})
		}
		for _, rule := range ingress.Spec.Rules {
			r := new(RuleRaw)
			r.Host = rule.Host
			for _, p := range rule.HTTP.Paths {
				b := BackendRaw{p.Backend.ServiceName, p.Backend.ServicePort.IntVal, p.Path}
				r.Backends = append(r.Backends, b)
			}
			ing.Rules = append(ing.Rules, *r)
		}
		is = append(is, *ing)
	}
	return
}

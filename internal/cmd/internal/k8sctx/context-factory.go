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

// ContextFactory produces k8s context
type ContextFactory struct {
	configs  []*k8s.KubeConfig
	gslbName string
}

// NewContextFactory returns context of command
func NewContextFactory(yaml, gslb string) (factory *ContextFactory, err error) {
	factory = new(ContextFactory)
	var k8sf *k8s.KubeConfigFactory
	k8sf, err = k8s.NewKubeConfigFactory(yaml, gslb)
	if err != nil {
		return
	}
	factory.configs, err = k8sf.InitializeConfigs()
	factory.gslbName = k8sf.GetYamlName()
	if err != nil {
		return
	}
	return
}

func (f *ContextFactory) GetGSLBName() string {
	return f.gslbName
}

// List returns list of GSLBs within namespaces
func (f *ContextFactory) List() (m []model.ListItem, err error) {
	m = make([]model.ListItem, 0)
	raws, err := readRaw(f.configs)
	if err != nil {
		return m, err
	}
	for _, raw := range raws.Gslb {
		if raw.Name != f.GetGSLBName() {
			continue
		}
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

// GetStatus returns gslb status across all configured contexts
func (f *ContextFactory) GetStatus() (m []model.Status, err error) {
	//Do validations and transitions here!
	m = make([]model.Status, 0)
	raws, err := readRaw(f.configs)
	for _, raw := range raws.Gslb {
		if raw.Name != f.gslbName {
			continue
		}
		status := model.Status{}
		status.Host = raw.Cluster
		status.Name = raw.Name
		status.GeoTag = raw.GeoTag
		status.Type = raw.Type
		status.Namespace = raw.Namespace
		for _, ingRaw := range raw.Ingress {
			ing := model.Ingress{}
			ing.Name = ingRaw.Name
			ing.Annotations = ingRaw.Annotations

			for _, rawRule := range ingRaw.Rules {
				rul := model.Rule{}
				rul.Host = rawRule.Host
				for _, rb := range rawRule.Backends {
					rul.Backends = append(rul.Backends,
						model.Backend{
							Service: rb.Service,
							Port:    rb.Port,
							Path:    rb.Path})
				}
				ing.Rules = append(ing.Rules, rul)
			}
			status.Ingresses = append(status.Ingresses, ing)
		}
		m = append(m, status)
	}
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

// maps unstructured data into GslbRaw structure. Any CRD change has to be reflected
// in GslbRaw or underlying structures
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

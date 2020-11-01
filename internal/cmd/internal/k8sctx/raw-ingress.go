package k8sctx

type IngressRaw struct {
	Name          string
	Namespace     string
	Address       []string
	Annotations   map[string]string
	Rules         []RuleRaw
	Labels        map[string]string
	LoadBalancers []EndpointRaw
}

type EndpointRaw struct {
	IP   string
	Host string
}

type RuleRaw struct {
	Host     string
	Backends []BackendRaw
}

type BackendRaw struct {
	Service string
	Port    int32
	Path    string
}

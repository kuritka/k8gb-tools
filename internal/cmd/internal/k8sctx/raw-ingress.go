package k8sctx

type IngressRaw struct {
	Name        string
	Namespace   string
	Address     []string
	Annotations map[string]string
	Rules       []RuleRaw
	Labels      map[string]string
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

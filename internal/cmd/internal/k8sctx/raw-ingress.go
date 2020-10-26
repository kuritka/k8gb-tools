package k8sctx

type IngressRaw struct {
	Name        string
	Namespace   string
	Address     []string
	Annotations map[string]string
	Rules  []RuleRaw
}

type RuleRaw struct {
	Host string
}
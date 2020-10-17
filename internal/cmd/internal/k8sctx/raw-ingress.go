package k8sctx

type IngressRaw struct {
	Name string
	Namespace string
	Address []string
	Annotations map[string]string
	Rule[] struct{
		Host string
	}
}
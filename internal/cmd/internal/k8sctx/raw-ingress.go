package k8sctx

type IngressRaw struct {
	Name string
	Namespace string
	Address []string
	Annotations []string
	Rule[] struct{
		Host string
	}
}
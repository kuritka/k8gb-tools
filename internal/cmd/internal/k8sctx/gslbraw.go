package k8sctx

//GslbRaw Raw configuration of GSLB
type GslbRaw struct {
	Kind           string
	Cluster        string
	APIVersion     string
	Error          error
	Source         string
	CurrentContext string
	Ingress        []IngressRaw
	Metadata
	Status
	Spec
}

//Metadata Raw metadata of GSLB
type Metadata struct {
	Name        string
	Namespace   string
	Annotations map[string]string
}

//Status Raw status of GSLB
type Status struct {
	GeoTag         string
	HealthyRecords map[string][]string
	ServiceHealth  map[string]string
}

//Spec Raw spec of GSLB
type Spec struct {
	Ingress
	Strategy
}

//Strategy Raw strategy of GSLB
type Strategy struct {
	DNSTtlSeconds              int64 `json:"dnsTtlSeconds,string"`
	SplitBrainThresholdSeconds int64 `json:"splitBrainThresholdSeconds,string"`
	Type                       string
}

//Ingress Raw ingress of GSLB
type Ingress struct {
	Rules []Rule
}

//Rule Raw rule of GSLB
type Rule struct {
	//Host
	Host string
	//HTTP
	HTTP
}

//HTTP Raw http of GSLB
type HTTP struct {
	//Paths
	Paths []Path
}

//Path Raw path of GSLB
type Path struct {
	//Backend
	Backend
	//Path
	Path string
}

//Backend Raw backend of GSLB
type Backend struct {
	ServiceName string
	//ServicePort TODO:local gslb setup contains value http.thats why stored as string now
	ServicePort int64 `json:"servicePort,string"`
}

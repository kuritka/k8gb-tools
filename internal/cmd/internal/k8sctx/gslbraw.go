package k8sctx

type GslbRaw struct {
	Kind       string
	Cluster    string
	ApiVersion string
	Error      error
	Metadata
	Status
	Spec
}

type Metadata struct {
	Name        string
	Namespace   string
	Annotations map[string]string
}

type Status struct {
	GeoTag         string
	HealthyRecords map[string][]string
	ServiceHealth  map[string]string
}

type Spec struct {
	Ingress
	Strategy
}

type Strategy struct {
	DnsTtlSeconds              int64 `json:"dnsTtlSeconds,string"`
	SplitBrainThresholdSeconds int64 `json:"splitBrainThresholdSeconds,string"`
	Type                       string
}

type Ingress struct {
	Rules []Rule
}

type Rule struct {
	Host string
	Http
}

type Http struct {
	Paths []Path
}

type Path struct {
	Backend
	Path string
}

type Backend struct {
	ServiceName string
	//ServicePort TODO:local gslb setup contains value http.thats why stored as string now
	ServicePort string `json:"servicePort,string"`
}

package k8sctx

//GslbRaw raw configuration of GSLB
type GslbRaw struct {
	Kind           string
	Cluster        string
	ApiVersion     string
	Error          error
	Source         string
	CurrentContext string
	Metadata
	Status
	Spec
}

//Metadata raw metadata of GSLB
type Metadata struct {
	Name        string
	Namespace   string
	Annotations map[string]string
}

//Status raw status of GSLB
type Status struct {
	GeoTag         string
	HealthyRecords map[string][]string
	ServiceHealth  map[string]string
}

//Spec raw spec of GSLB
type Spec struct {
	Ingress
	Strategy
}

//Strategy raw strategy of GSLB
type Strategy struct {
	DnsTtlSeconds              int64 `json:"dnsTtlSeconds,string"`
	SplitBrainThresholdSeconds int64 `json:"splitBrainThresholdSeconds,string"`
	Type                       string
}

//Ingress raw ingress of GSLB
type Ingress struct {
	Rules []Rule
}

//Rule raw rule of GSLB
type Rule struct {
	//Host
	Host string
	//Http
	Http
}

//Http raw http of GSLB
type Http struct {
	//Paths
	Paths []Path
}

//Path raw path of GSLB
type Path struct {
	//Backend
	Backend
	//Path
	Path string
}

//Backend raw backend of GSLB
type Backend struct {
	ServiceName string
	//ServicePort TODO:local gslb setup contains value http.thats why stored as string now
	ServicePort int64 `json:"servicePort,string"`
}

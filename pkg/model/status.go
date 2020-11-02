package model

//Status
type Status struct {
	// GeoTag
	GeoTag string
	// Name
	Name string
	// Name
	Namespace string
	// Type
	Type string
	// Host
	Host string
	// Ingresses
	Ingresses []Ingress
	// Dig
	Dig []string
}

//Ingress
type Ingress struct {
	Name          string
	Address       []string
	Rules         []Rule
	LoadBalancers []Endpoint
	Annotations   map[string]string
}

type Endpoint struct {
	IP   string
	Host string
}

type Rule struct {
	Host     string
	Backends []Backend
}

type Backend struct {
	Service string
	Port    int32
	Path    string
}

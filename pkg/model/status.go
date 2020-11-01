package model

//Status
type Status struct {
	//GeoTag
	GeoTag string
	//Name
	Name string
	//Name
	Namespace string
	//Type
	Type string
	//Host
	Host string
	//Ingresses
	Ingresses []Ingress
}

//Ingress
type Ingress struct {
	Name  string
	Rules []struct {
		Host      string
		IPAddress string
		Node      string
	}
	Annotations map[string]string
}

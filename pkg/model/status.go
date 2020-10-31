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
	Rules []struct {
		Host      string
		IpAddress string
		Node      string
	}
	Annotations string
	Name string
}




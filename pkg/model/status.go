package model

//Status
type Status struct {
	//GeoTag
	GeoTag Stringr
	//Name
	Name Stringr
	//Type
	Type Stringr
	//Host
	Host Stringr
	//Ingress
	Ingress struct {
		Rule []struct {
			Host      Stringr
			IpAddress Stringr
			Node      Stringr
		}
		Annotations Stringr
	}
}

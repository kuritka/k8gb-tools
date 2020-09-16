package runner

//ICmdRunner is tunning commandline commands
type ICmdRunner interface {
	Run() error
	String() string
}

package model

type ProxyType string

const (
	ProxyTCP  = ProxyType("tcp")
	ProxyHTTP = ProxyType("http")
)

type ProxyDetail struct {
	Type       ProxyType
	IsRunning  bool
	Host       string // example : localhost:5555
	FullScheme string // example : http://localhost:5555/
}

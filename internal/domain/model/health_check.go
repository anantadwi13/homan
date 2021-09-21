package model

type HealthCheckType string

const HealthCheckTCP = HealthCheckType("tcp")
const HealthCheckHTTP = HealthCheckType("http")

type HealthCheck interface {
	Type() HealthCheckType
	Port() int
	// Endpoint would return "/health-check-endpoint" for http
	// or empty "" for tcp
	Endpoint() string

	Validator
}

type healthCheck struct {
	hcType   HealthCheckType
	port     int
	endpoint string
}

//NewHealthCheckHTTP will expect HTTP success code (2xx)
func NewHealthCheckHTTP(containerPort int, endpoint string) HealthCheck {
	return &healthCheck{hcType: HealthCheckHTTP, port: containerPort, endpoint: endpoint}
}

func NewHealthCheckTCP(containerPort int) HealthCheck {
	return &healthCheck{hcType: HealthCheckTCP, port: containerPort}
}

func (h *healthCheck) Type() HealthCheckType {
	return h.hcType
}

func (h *healthCheck) Port() int {
	return h.port
}

func (h *healthCheck) Endpoint() string {
	return h.endpoint
}

func (h *healthCheck) IsValid() bool {
	if h.port < 0 || h.port > 65535 {
		return false
	}
	return true
}

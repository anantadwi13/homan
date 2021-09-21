package homand

import (
	"context"
	"errors"
	"net"
	"net/http"
)

type HealthCheckType string

const (
	HealthCheckHTTP = HealthCheckType("http")
	HealthCheckTCP  = HealthCheckType("tcp")
)

type HealthChecker interface {
	IsAvailable(ctx context.Context, mode HealthCheckType, address string) (bool, error)
}

type healthChecker struct {
	httpClient *http.Client
}

func NewHealthChecker() HealthChecker {
	httpClient := *http.DefaultClient

	return &healthChecker{
		httpClient: &httpClient,
	}
}

func (h healthChecker) IsAvailable(ctx context.Context, mode HealthCheckType, address string) (bool, error) {
	switch mode {
	case HealthCheckHTTP:
		res, err := h.httpClient.Get(address)
		if err != nil {
			return false, nil
		}
		defer res.Body.Close()
		return res.StatusCode >= 200 && res.StatusCode < 300, nil
	case HealthCheckTCP:
		conn, err := net.Dial("tcp", address)
		if err != nil {
			return false, nil
		}
		defer conn.Close()
		return true, nil
	default:
		return false, errors.New("unkown health check type")
	}
}

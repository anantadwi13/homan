package service

import (
	"context"
	"github.com/anantadwi13/homan/internal/homan/domain/model"
)

type ProxyParams struct {
	Type        model.ProxyType
	TCPHostname string // required for tcp proxy, example "service-haproxy"
	TCPPort     int    // required for tcp proxy, example 5555
}

type Proxy interface {
	Start(ctx context.Context, params ProxyParams) (proxy *model.ProxyDetail, stop func() error, err error)
}

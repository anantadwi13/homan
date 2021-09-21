package service

import (
	"context"
	"github.com/anantadwi13/homan/internal/homan/domain/model"
)

type Proxy interface {
	Execute(ctx context.Context, request func(proxy *model.ProxyDetail) error) error
}

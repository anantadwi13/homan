package service

import (
	"context"
	"github.com/anantadwi13/cli-whm/internal/domain/model"
)

type Proxy interface {
	Execute(ctx context.Context, request func(proxy *model.ProxyDetail) error) error
}

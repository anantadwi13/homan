package service

import (
	"context"
	"github.com/anantadwi13/cli-whm/internal/domain/model"
)

type Executor interface {
	Run(ctx context.Context, config model.ServiceConfig)
	Stop(ctx context.Context, config model.ServiceConfig)
	Restart(ctx context.Context, config model.ServiceConfig)
}

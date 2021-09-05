package service

import (
	"context"
	"github.com/anantadwi13/cli-whm/internal/domain/model"
	domainService "github.com/anantadwi13/cli-whm/internal/domain/service"
)

type dockerExecutor struct {
}

func NewDockerExecutor() domainService.Executor {
	return &dockerExecutor{}
}

func (d *dockerExecutor) Run(ctx context.Context, config model.ServiceConfig) {
	panic("implement me")
}

func (d *dockerExecutor) Stop(ctx context.Context, config model.ServiceConfig) {
	panic("implement me")
}

func (d *dockerExecutor) Restart(ctx context.Context, config model.ServiceConfig) {
	panic("implement me")
}

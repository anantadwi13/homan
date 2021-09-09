package usecase

import (
	"context"
	"github.com/anantadwi13/cli-whm/internal/domain/service"
)

type UcDownParams struct {
}

type UcDown interface {
	Execute(ctx context.Context, params *UcDownParams) error
}

type ucDown struct {
	registry service.Registry
	executor service.Executor
}

func NewUcDown(registry service.Registry, executor service.Executor) UcDown {
	return &ucDown{registry, executor}
}

func (u *ucDown) Execute(ctx context.Context, params *UcDownParams) error {
	systemServices, err := u.registry.GetSystemServices(ctx)
	if err != nil {
		return err
	}
	userServices, err := u.registry.GetUserServices(ctx)
	if err != nil {
		return err
	}
	for _, serviceConfig := range systemServices {
		err = u.executor.Stop(ctx, serviceConfig)
		if err != nil && err != service.ErrorExecutorServiceIsNotRunning {
			return err
		}
	}
	for _, serviceConfig := range userServices {
		err = u.executor.Stop(ctx, serviceConfig)
		if err != nil && err != service.ErrorExecutorServiceIsNotRunning {
			return err
		}
	}
	return nil
}

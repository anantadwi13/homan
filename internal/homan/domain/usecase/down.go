package usecase

import (
	"context"
	service2 "github.com/anantadwi13/homan/internal/homan/domain/service"
)

type UcDownParams struct {
}

type UcDown interface {
	Execute(ctx context.Context, params *UcDownParams) Error
}

type ucDown struct {
	registry service2.Registry
	executor service2.Executor
}

func NewUcDown(registry service2.Registry, executor service2.Executor) UcDown {
	return &ucDown{registry, executor}
}

func (u *ucDown) Execute(ctx context.Context, params *UcDownParams) Error {
	systemServices, err := u.registry.GetSystemServices(ctx)
	if err != nil {
		return WrapErrorSystem(err)
	}
	userServices, err := u.registry.GetUserServices(ctx)
	if err != nil {
		return WrapErrorSystem(err)
	}
	for _, serviceConfig := range systemServices {
		err = u.executor.Stop(ctx, serviceConfig)
		if err != nil && err != service2.ErrorExecutorServiceIsNotRunning {
			return WrapErrorSystem(err)
		}
	}
	for _, serviceConfig := range userServices {
		err = u.executor.Stop(ctx, serviceConfig)
		if err != nil && err != service2.ErrorExecutorServiceIsNotRunning {
			return WrapErrorSystem(err)
		}
	}
	return nil
}

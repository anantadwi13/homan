package usecase

import (
	"context"
	"github.com/anantadwi13/cli-whm/internal/domain/service"
)

type UcUpParams struct {
}

type UcUp interface {
	Execute(ctx context.Context, params *UcUpParams) Error
}

type ucUp struct {
	registry service.Registry
	executor service.Executor
}

func NewUcUp(registry service.Registry, executor service.Executor) UcUp {
	return &ucUp{registry, executor}
}

func (u *ucUp) Execute(ctx context.Context, params *UcUpParams) Error {
	systemServices, err := u.registry.GetSystemServices(ctx)
	if err != nil {
		return WrapErrorSystem(err)
	}
	userServices, err := u.registry.GetUserServices(ctx)
	if err != nil {
		return WrapErrorSystem(err)
	}

	for _, systemService := range systemServices {
		err = u.executor.Run(ctx, systemService)
		if err != nil && err != service.ErrorExecutorServiceIsRunning {
			return WrapErrorSystem(err)
		}
	}

	for _, userService := range userServices {
		err = u.executor.Run(ctx, userService)
		if err != nil && err != service.ErrorExecutorServiceIsRunning {
			return WrapErrorSystem(err)
		}
	}

	return nil
}

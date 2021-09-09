package usecase

import (
	"context"
	"github.com/anantadwi13/cli-whm/internal/domain"
	"github.com/anantadwi13/cli-whm/internal/domain/model"
	"github.com/anantadwi13/cli-whm/internal/domain/service"
	"path/filepath"
)

type ServiceType string

type UcAddParams struct {
	Name        string
	Domain      string
	ServiceType ServiceType
}

type UcAdd interface {
	Execute(ctx context.Context, params *UcAddParams) Error
}

var (
	ServiceTypeBlog   = ServiceType("blog")
	ServiceTypeCustom = ServiceType("custom")

	ErrorUcAddSystemNotRunning = NewErrorUser("system service is not running")
	ErrorUcAddParamsNotFound   = NewErrorUser("please specify parameters")
)

type ucAdd struct {
	config   domain.Config
	registry service.Registry
	executor service.Executor
}

func NewUcAdd(config domain.Config, registry service.Registry, executor service.Executor) UcAdd {
	return &ucAdd{config, registry, executor}
}

func (u *ucAdd) Execute(ctx context.Context, params *UcAddParams) Error {
	err := u.preExecute(ctx, params)
	if err != nil {
		return err
	}

	var config model.ServiceConfig
	switch params.ServiceType {
	case ServiceTypeBlog:
		config = model.NewServiceConfig(
			params.Name,
			params.Domain,
			"wordpress:5.8.0-apache",
			[]string{},
			[]model.Port{
				model.NewPort(80),
			},
			[]model.Volume{
				model.NewVolumeBinding(filepath.Join(u.config.DataPath(), params.Name+"/wp-content/"), "/var/www/html/wp-content/"),
			},
			[]string{u.config.ProjectName()},
			model.TagWeb,
		)
	case ServiceTypeCustom:
	}

	if config == nil {
		return NewErrorSystem("unable to create service")
	}

	errAdd := u.registry.Add(ctx, config)
	if errAdd != nil {
		return WrapErrorSystem(errAdd)
	}

	err = u.postExecute(ctx, params, config)
	if err != nil {
		return err
	}

	return nil
}

func (u *ucAdd) preExecute(ctx context.Context, params *UcAddParams) Error {
	if params == nil || params.Name == "" || params.Domain == "" {
		return ErrorUcAddParamsNotFound
	}

	switch params.ServiceType {
	case ServiceTypeBlog:
	case ServiceTypeCustom:
	default:
		return NewErrorUser("unkown service type")
	}

	systemServices, err := u.registry.GetSystemServices(ctx)
	if err != nil {
		return WrapErrorSystem(err)
	}

	for _, systemService := range systemServices {
		isRunning, err := u.executor.IsRunning(ctx, systemService)
		if err != nil {
			return WrapErrorSystem(err)
		}
		if !isRunning {
			return ErrorUcAddSystemNotRunning
		}
	}

	userServices, err := u.registry.GetUserServices(ctx)
	if err != nil {
		return WrapErrorSystem(err)
	}

	for _, userService := range userServices {
		if userService.Name() == params.Name {
			return NewErrorUser("duplicate service name")
		}
		if userService.DomainName() == params.Domain {
			return NewErrorUser("duplicate domain name")
		}
	}

	return nil
}

func (u *ucAdd) postExecute(ctx context.Context, params *UcAddParams, config model.ServiceConfig) Error {
	return nil
}

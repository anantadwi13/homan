package service

import (
	"context"
	"errors"
	"github.com/anantadwi13/cli-whm/internal/homan/domain/model"
)

var (
	ErrorRegistryServiceConfigInvalid  = errors.New("error [registry]: service config is invalid")
	ErrorRegistryServiceConfigExist    = errors.New("error [registry]: service config already exists")
	ErrorRegistryServiceConfigNotFound = errors.New("error [registry]: service config is not found")
)

type Registry interface {
	GetAll(ctx context.Context) ([]model.ServiceConfig, error)
	GetSystemServices(ctx context.Context) ([]model.ServiceConfig, error)
	GetSystemServiceByTag(ctx context.Context, tag model.ServiceTag) ([]model.ServiceConfig, error)
	GetUserServices(ctx context.Context) ([]model.ServiceConfig, error)
	Add(ctx context.Context, config model.ServiceConfig) error
	AddSystem(ctx context.Context, config model.ServiceConfig) error
	Remove(ctx context.Context, config model.ServiceConfig) error
	RemoveSystem(ctx context.Context, config model.ServiceConfig) error
	IsSystem(ctx context.Context, config model.ServiceConfig) (bool, error)
}

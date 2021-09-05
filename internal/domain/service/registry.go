package service

import (
	"context"
	"errors"
	"github.com/anantadwi13/cli-whm/internal/domain/model"
)

var (
	RegistryErrorServiceConfigInvalid = errors.New("error [registry]: service config is invalid")
	RegistryErrorServiceConfigExist   = errors.New("error [registry]: service config already exists")
)

type Registry interface {
	GetAll(ctx context.Context) ([]model.ServiceConfig, error)
	GetSystemServices(ctx context.Context) ([]model.ServiceConfig, error)
	GetUserServices(ctx context.Context) ([]model.ServiceConfig, error)
	Add(ctx context.Context, config model.ServiceConfig) error
	Remove(ctx context.Context, config model.ServiceConfig) error
	IsSystem(ctx context.Context, config model.ServiceConfig) (bool, error)
}

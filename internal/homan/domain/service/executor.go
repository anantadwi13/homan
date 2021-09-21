package service

import (
	"context"
	"errors"
	"github.com/anantadwi13/cli-whm/internal/homan/domain/model"
)

var (
	ErrorExecutorServiceConfigInvalid = errors.New("error [executor]: service config is invalid")
	ErrorExecutorServiceIsRunning     = errors.New("error [executor]: service is running")
	ErrorExecutorServiceIsNotRunning  = errors.New("error [executor]: service is not running")
)

type Executor interface {
	InitVolume(ctx context.Context, configs ...model.ServiceConfig) error
	Run(ctx context.Context, configs ...model.ServiceConfig) error
	Stop(ctx context.Context, configs ...model.ServiceConfig) error
	Restart(ctx context.Context, configs ...model.ServiceConfig) error
	IsRunning(ctx context.Context, config model.ServiceConfig) (bool, error)
}

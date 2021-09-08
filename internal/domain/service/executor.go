package service

import (
	"context"
	"errors"
	"github.com/anantadwi13/cli-whm/internal/domain/model"
)

var (
	ErrorExecutorServiceConfigInvalid = errors.New("error [executor]: service config is invalid")
	ErrorExecutorServiceIsRunning     = errors.New("error [executor]: service is running")
	ErrorExecutorServiceIsNotRunning  = errors.New("error [executor]: service is not running")
)

type Executor interface {
	RunAll(ctx context.Context) error
	Run(ctx context.Context, configs ...model.ServiceConfig) error
	Stop(ctx context.Context, configs ...model.ServiceConfig) error
	Restart(ctx context.Context, configs ...model.ServiceConfig) error
	IsRunning(ctx context.Context, config model.ServiceConfig) (bool, error)
}

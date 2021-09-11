package usecase

import (
	"context"
)

type UcInitParams struct {
}

type UcInit interface {
	Execute(ctx context.Context, params *UcInitParams) Error
}

var (
	ErrorUcInitAlreadyInitialized   = NewErrorUser("project is already initialized")
	ErrorUcInitServiceConfigInvalid = NewErrorUser("service config is invalid")
)

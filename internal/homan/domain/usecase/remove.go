package usecase

import (
	"context"
)

type UcRemoveParams struct {
	Name string
}

type UcRemove interface {
	Execute(ctx context.Context, params *UcRemoveParams) Error
}

var (
	ErrorUcRemoveSystemNotRunning = NewErrorUser("system service is not running")
	ErrorUcRemoveParamsNotFound   = NewErrorUser("please specify parameters")
	ErrorUcRemoveServiceNotFound  = NewErrorUser("service is not found")
	ErrorUcRemovePostExecution    = NewErrorUser("something went wrong while doing post-execution")
)

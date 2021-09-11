package usecase

import (
	"context"
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
	ErrorUcAddPostExecution    = NewErrorUser("something went wrong while doing post-execution")
)

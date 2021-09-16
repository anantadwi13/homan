package usecase

import (
	"context"
	"github.com/anantadwi13/cli-whm/internal/domain/model"
	"github.com/anantadwi13/cli-whm/internal/domain/service"
	"github.com/anantadwi13/cli-whm/internal/external/api/haproxy"
	"github.com/anantadwi13/cli-whm/internal/external/api/haproxy/client/configuration"
	"github.com/anantadwi13/cli-whm/internal/external/api/haproxy/client/transactions"
	"github.com/anantadwi13/cli-whm/internal/util"
)

type UcUpParams struct {
}

type UcUp interface {
	Execute(ctx context.Context, params *UcUpParams) Error
}

var (
	ErrorUcUpPostExecution = NewErrorUser("error post execution")
)

type ucUp struct {
	registry service.Registry
	executor service.Executor
	proxy    service.Proxy
}

func NewUcUp(registry service.Registry, executor service.Executor, proxy service.Proxy) UcUp {
	return &ucUp{registry, executor, proxy}
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

	services, err := u.registry.GetSystemServiceByTag(ctx, model.TagGateway)
	if err != nil {
		return WrapErrorSystem(err)
	}
	if len(services) != 1 {
		return ErrorUcUpPostExecution
	}
	haproxyService := services[0]

	err = u.proxy.Execute(ctx, func(proxy *model.ProxyDetail) error {

		// Force Reload HAProxy

		haproxyClient, auth := haproxy.NewHaproxyClient(proxy.Host, haproxyService.Name()+":5555")

		version, err := haproxyClient.Configuration.GetConfigurationVersion(configuration.NewGetConfigurationVersionParams(), auth)
		if err != nil {
			return err
		}

		transaction, err := haproxyClient.Transactions.StartTransaction(transactions.NewStartTransactionParams().WithVersion(version.Payload), auth)
		if err != nil {
			return err
		}

		transactionId := &transaction.Payload.ID

		_, _, err = haproxyClient.Transactions.CommitTransaction(
			transactions.NewCommitTransactionParams().WithID(*transactionId).WithForceReload(util.Bool(true)),
			auth,
		)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return WrapErrorSystem(err)
	}

	return nil
}

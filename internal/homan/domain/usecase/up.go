package usecase

import (
	"context"
	domainModel "github.com/anantadwi13/homan/internal/homan/domain/model"
	domainService "github.com/anantadwi13/homan/internal/homan/domain/service"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/configuration"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/transactions"
	"github.com/anantadwi13/homan/internal/util"
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
	registry domainService.Registry
	executor domainService.Executor
	proxy    domainService.Proxy
}

func NewUcUp(registry domainService.Registry, executor domainService.Executor, proxy domainService.Proxy) UcUp {
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

	err = u.executor.RunWait(ctx, 10, u.registry.GetCoreDaemon(ctx))
	if err != nil && err != domainService.ErrorExecutorServiceIsRunning {
		return WrapErrorSystem(err)
	}

	for _, systemService := range systemServices {
		err = u.executor.RunWait(ctx, 60, systemService)
		if err != nil && err != domainService.ErrorExecutorServiceIsRunning {
			return WrapErrorSystem(err)
		}
	}

	for _, userService := range userServices {
		err = u.executor.RunWait(ctx, 60, userService)
		if err != nil && err != domainService.ErrorExecutorServiceIsRunning {
			return WrapErrorSystem(err)
		}
	}

	services, err := u.registry.GetSystemServiceByTag(ctx, domainModel.TagGateway)
	if err != nil {
		return WrapErrorSystem(err)
	}
	if len(services) != 1 {
		return ErrorUcUpPostExecution
	}
	haproxyService := services[0]

	proxy, stop, err := u.proxy.Start(ctx, domainService.ProxyParams{Type: domainModel.ProxyHTTP})
	if err != nil {
		return WrapErrorSystem(err)
	}
	defer stop()

	// Force Reload HAProxy

	haproxyClient, auth := haproxy.NewHaproxyClient(proxy.Host, haproxyService.Name()+":5555")

	version, err := haproxyClient.Configuration.GetConfigurationVersion(configuration.NewGetConfigurationVersionParams(), auth)
	if err != nil {
		return WrapErrorSystem(err)
	}

	transaction, err := haproxyClient.Transactions.StartTransaction(transactions.NewStartTransactionParams().WithVersion(version.Payload), auth)
	if err != nil {
		return WrapErrorSystem(err)
	}

	transactionId := &transaction.Payload.ID

	_, _, err = haproxyClient.Transactions.CommitTransaction(
		transactions.NewCommitTransactionParams().WithID(*transactionId).WithForceReload(util.Bool(true)),
		auth,
	)
	if err != nil {
		return WrapErrorSystem(err)
	}

	return nil
}

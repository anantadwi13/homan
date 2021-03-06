package usecase

import (
	"context"
	"errors"
	"github.com/anantadwi13/homan/internal/homan/domain"
	"github.com/anantadwi13/homan/internal/homan/domain/model"
	"github.com/anantadwi13/homan/internal/homan/domain/service"
	"github.com/anantadwi13/homan/internal/homan/domain/usecase"
	"github.com/anantadwi13/homan/internal/homan/external/api/dns"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/backend"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/backend_switching_rule"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/configuration"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/http_request_rule"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/storage"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/transactions"
	"github.com/anantadwi13/homan/internal/util"
	"strings"
)

type ucRemove struct {
	config   domain.Config
	registry service.Registry
	executor service.Executor
	proxy    service.Proxy
}

func NewUcRemove(
	config domain.Config, registry service.Registry, executor service.Executor, proxy service.Proxy,
) usecase.UcRemove {
	return &ucRemove{config: config, registry: registry, executor: executor, proxy: proxy}
}

func (u *ucRemove) Execute(ctx context.Context, params *usecase.UcRemoveParams) usecase.Error {
	domainErr := u.preExecute(ctx, params)
	if domainErr != nil {
		return domainErr
	}

	services, err := u.registry.GetUserServices(ctx)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}

	var serviceConfig model.ServiceConfig = nil

	for _, config := range services {
		if config.Name() == params.Name {
			serviceConfig = config
			break
		}
	}

	if serviceConfig == nil {
		return usecase.ErrorUcRemoveServiceNotFound
	}

	err = u.executor.Stop(ctx, serviceConfig)
	if err != nil && err != service.ErrorExecutorServiceIsNotRunning {
		return usecase.WrapErrorSystem(err)
	}

	err = u.registry.Remove(ctx, serviceConfig)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}

	domainErr = u.postExecute(ctx, serviceConfig)
	if domainErr != nil {
		return domainErr
	}

	return nil
}

func (u *ucRemove) preExecute(ctx context.Context, params *usecase.UcRemoveParams) usecase.Error {
	if params == nil || params.Name == "" {
		return usecase.ErrorUcRemoveParamsNotFound
	}

	systemServices, err := u.registry.GetSystemServices(ctx)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}

	for _, systemService := range systemServices {
		isRunning, err := u.executor.IsRunning(ctx, systemService, false)
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}
		if !isRunning {
			return usecase.ErrorUcRemoveSystemNotRunning
		}
	}

	return nil
}

func (u *ucRemove) postExecute(ctx context.Context, config model.ServiceConfig) usecase.Error {

	services, err := u.registry.GetSystemServiceByTag(ctx, model.TagGateway)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}
	if len(services) != 1 {
		return usecase.ErrorUcRemovePostExecution
	}
	haproxyService := services[0]

	services, err = u.registry.GetSystemServiceByTag(ctx, model.TagDNS)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}
	if len(services) != 1 {
		return usecase.ErrorUcRemovePostExecution
	}
	dnsService := services[0]

	if config.DomainName() != "" {
		proxy, stop, err := u.proxy.Start(ctx, service.ProxyParams{Type: model.ProxyHTTP})
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}
		defer stop()

		// Setup HAProxy

		haproxyClient, auth := haproxy.NewHaproxyClient(proxy.Host, haproxyService.Name()+":5555")

		version, err := haproxyClient.Configuration.GetConfigurationVersion(configuration.NewGetConfigurationVersionParams(), auth)
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}

		transaction, err := haproxyClient.Transactions.StartTransaction(transactions.NewStartTransactionParams().WithVersion(version.Payload), auth)
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}

		transactionId := &transaction.Payload.ID
		mainFrontendName := u.config.ProjectName()

		// Delete Backend Switching Rules

		rules, err := haproxyClient.BackendSwitchingRule.GetBackendSwitchingRules(
			backend_switching_rule.NewGetBackendSwitchingRulesParams().WithTransactionID(transactionId).WithFrontend(mainFrontendName),
			auth,
		)
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}

		for i := len(rules.Payload.Data) - 1; i >= 0; i-- {
			// Reverse deletion
			rule := rules.Payload.Data[i]
			if rule.Name == config.Name() {
				_, _, err = haproxyClient.BackendSwitchingRule.DeleteBackendSwitchingRule(
					backend_switching_rule.NewDeleteBackendSwitchingRuleParams().WithTransactionID(transactionId).WithFrontend(mainFrontendName).WithIndex(*rule.Index),
					auth,
				)
				if err != nil {
					return usecase.WrapErrorSystem(err)
				}
			}
		}

		// Delete Http Request Rules

		requestRules, err := haproxyClient.HTTPRequestRule.GetHTTPRequestRules(
			http_request_rule.NewGetHTTPRequestRulesParams().WithParentType("frontend").WithParentName(mainFrontendName).WithTransactionID(transactionId),
			auth,
		)
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}

		for i := len(requestRules.Payload.Data) - 1; i >= 0; i-- {
			// Reverse deletion
			rule := requestRules.Payload.Data[i]
			if strings.Contains(rule.CondTest, config.DomainName()) {
				_, _, err = haproxyClient.HTTPRequestRule.DeleteHTTPRequestRule(
					http_request_rule.NewDeleteHTTPRequestRuleParams().WithParentType("frontend").WithParentName(mainFrontendName).WithTransactionID(transactionId).WithIndex(*rule.Index),
					auth,
				)
				if err != nil {
					return usecase.WrapErrorSystem(err)
				}
			}
		}

		_, _, err = haproxyClient.Backend.DeleteBackend(
			backend.NewDeleteBackendParams().WithTransactionID(transactionId).WithName(config.Name()),
			auth,
		)
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}

		_, _, err = haproxyClient.Transactions.CommitTransaction(transactions.NewCommitTransactionParams().WithID(*transactionId).WithForceReload(util.Bool(true)), auth)
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}

		// Delete Domain Name
		dnsClient, err := dns.NewDnsClient(proxy.FullScheme, dnsService.Name()+":5555")
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}

		deleteZone, err := dnsClient.DeleteZoneWithResponse(ctx, config.DomainName())
		if err != nil || deleteZone.JSON200 == nil {
			if err == nil {
				err = errors.New("domain : unable to delete zone")
			}
			return usecase.WrapErrorSystem(err)
		}

		// Delete Certificate in HAProxy
		_, err = haproxyClient.Storage.GetOneStorageSSLCertificate(storage.NewGetOneStorageSSLCertificateParams().WithName(config.DomainName()), auth)
		if err == nil {
			_, _, err = haproxyClient.Storage.DeleteStorageSSLCertificate(
				storage.NewDeleteStorageSSLCertificateParams().WithName(config.DomainName()).WithForceReload(util.Bool(true)),
				auth,
			)
			if err != nil {
				return usecase.WrapErrorSystem(err)
			}
		}
	}
	return nil
}

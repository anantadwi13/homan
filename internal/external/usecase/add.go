package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/anantadwi13/cli-whm/internal/domain"
	"github.com/anantadwi13/cli-whm/internal/domain/model"
	"github.com/anantadwi13/cli-whm/internal/domain/service"
	domainUsecase "github.com/anantadwi13/cli-whm/internal/domain/usecase"
	"github.com/anantadwi13/cli-whm/internal/external/api/certman"
	"github.com/anantadwi13/cli-whm/internal/external/api/dns"
	"github.com/anantadwi13/cli-whm/internal/external/api/haproxy"
	"github.com/anantadwi13/cli-whm/internal/external/api/haproxy/client/backend"
	"github.com/anantadwi13/cli-whm/internal/external/api/haproxy/client/backend_switching_rule"
	"github.com/anantadwi13/cli-whm/internal/external/api/haproxy/client/configuration"
	"github.com/anantadwi13/cli-whm/internal/external/api/haproxy/client/server"
	"github.com/anantadwi13/cli-whm/internal/external/api/haproxy/client/transactions"
	"github.com/anantadwi13/cli-whm/internal/external/api/haproxy/models"
	"github.com/anantadwi13/cli-whm/internal/util"
	"path/filepath"
)

type ucAdd struct {
	config   domain.Config
	registry service.Registry
	executor service.Executor
	proxy    service.Proxy
}

func NewUcAdd(
	config domain.Config, registry service.Registry, executor service.Executor, proxy service.Proxy,
) domainUsecase.UcAdd {
	return &ucAdd{config, registry, executor, proxy}
}

func (u *ucAdd) Execute(ctx context.Context, params *domainUsecase.UcAddParams) domainUsecase.Error {
	err := u.preExecute(ctx, params)
	if err != nil {
		return err
	}

	var config model.ServiceConfig
	switch params.ServiceType {
	case domainUsecase.ServiceTypeBlog:
		config = model.NewServiceConfig(
			params.Name,
			params.Domain,
			"wordpress:5.8.0-apache",
			[]string{},
			[]model.Port{
				model.NewPort(80),
			},
			[]model.Volume{
				model.NewVolumeBinding(filepath.Join(u.config.DataPath(), params.Name+"/wp-content/"), "/var/www/html/wp-content/"),
			},
			[]string{u.config.ProjectName()},
			model.TagWeb,
		)
	case domainUsecase.ServiceTypeCustom:
	}

	if config == nil {
		return domainUsecase.NewErrorSystem("unable to create service")
	}

	errAdd := u.registry.Add(ctx, config)
	if errAdd != nil {
		return domainUsecase.WrapErrorSystem(errAdd)
	}

	err = u.postExecute(ctx, params, config)
	if err != nil {
		return err
	}

	return nil
}

func (u *ucAdd) preExecute(ctx context.Context, params *domainUsecase.UcAddParams) domainUsecase.Error {
	if params == nil || params.Name == "" || params.Domain == "" {
		return domainUsecase.ErrorUcAddParamsNotFound
	}

	switch params.ServiceType {
	case domainUsecase.ServiceTypeBlog:
	case domainUsecase.ServiceTypeCustom:
	default:
		return domainUsecase.NewErrorUser("unkown service type")
	}

	systemServices, err := u.registry.GetSystemServices(ctx)
	if err != nil {
		return domainUsecase.WrapErrorSystem(err)
	}

	for _, systemService := range systemServices {
		isRunning, err := u.executor.IsRunning(ctx, systemService)
		if err != nil {
			return domainUsecase.WrapErrorSystem(err)
		}
		if !isRunning {
			return domainUsecase.ErrorUcAddSystemNotRunning
		}
	}

	userServices, err := u.registry.GetUserServices(ctx)
	if err != nil {
		return domainUsecase.WrapErrorSystem(err)
	}

	for _, userService := range userServices {
		if userService.Name() == params.Name {
			return domainUsecase.NewErrorUser("duplicate service name")
		}
		if userService.DomainName() == params.Domain {
			return domainUsecase.NewErrorUser("duplicate domain name")
		}
	}

	return nil
}

func (u *ucAdd) postExecute(
	ctx context.Context, params *domainUsecase.UcAddParams, config model.ServiceConfig,
) domainUsecase.Error {
	// Run service
	err := u.executor.Run(ctx, config)
	if err != nil {
		return nil
	}

	services, err := u.registry.GetSystemServiceByTag(ctx, model.TagGateway)
	if err != nil {
		return domainUsecase.WrapErrorSystem(err)
	}
	if len(services) != 1 {
		return domainUsecase.ErrorUcAddPostExecution
	}
	haproxyService := services[0]

	services, err = u.registry.GetSystemServiceByTag(ctx, model.TagDNS)
	if err != nil {
		return domainUsecase.WrapErrorSystem(err)
	}
	if len(services) != 1 {
		return domainUsecase.ErrorUcAddPostExecution
	}
	dnsService := services[0]

	services, err = u.registry.GetSystemServiceByTag(ctx, model.TagCertMan)
	if err != nil {
		return domainUsecase.WrapErrorSystem(err)
	}
	if len(services) != 1 {
		return domainUsecase.ErrorUcAddPostExecution
	}
	certmanService := services[0]

	if config.DomainName() != "" {
		err = u.proxy.Execute(ctx, func(proxy *model.ProxyDetail) error {

			// Add Haproxy Backend

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
			mainFrontendName := u.config.ProjectName()

			rules, err := haproxyClient.BackendSwitchingRule.GetBackendSwitchingRules(
				backend_switching_rule.NewGetBackendSwitchingRulesParams().WithTransactionID(transactionId).WithFrontend(mainFrontendName),
				auth,
			)
			if err != nil {
				return err
			}

			backendRules := rules.Payload.Data

			configBackend := &models.Backend{
				Name:       config.Name(),
				Mode:       "http",
				Forwardfor: &models.Forwardfor{Enabled: util.String("enabled")},
				Balance:    &models.Balance{Algorithm: util.String("roundrobin")},
			}

			configServer := &models.Server{
				Name:    "server1",
				Address: config.Name(),
				Port:    util.Int64(80),
				Check:   "enabled",
			}

			configBackendRule := &models.BackendSwitchingRule{
				Index:    util.Int64(int64(len(backendRules))),
				Name:     configBackend.Name,
				Cond:     "if",
				CondTest: fmt.Sprintf("{ hdr(host) -i %v }", config.DomainName()),
			}

			_, _, err = haproxyClient.Backend.CreateBackend(backend.NewCreateBackendParams().WithTransactionID(transactionId).WithData(configBackend), auth)
			if err != nil {
				return err
			}

			_, _, err = haproxyClient.Server.CreateServer(server.NewCreateServerParams().WithTransactionID(transactionId).WithBackend(configBackend.Name).WithData(configServer), auth)
			if err != nil {
				return err
			}

			_, _, err = haproxyClient.BackendSwitchingRule.CreateBackendSwitchingRule(
				backend_switching_rule.NewCreateBackendSwitchingRuleParams().WithTransactionID(transactionId).WithFrontend(mainFrontendName).WithData(configBackendRule),
				auth,
			)
			if err != nil {
				return err
			}

			_, _, err = haproxyClient.Transactions.CommitTransaction(transactions.NewCommitTransactionParams().WithID(*transactionId).WithForceReload(util.Bool(true)), auth)
			if err != nil {
				return err
			}

			// Add Domain Name
			dnsClient, err := dns.NewDnsClient(proxy.FullPath, dnsService.Name()+":5555")
			if err != nil {
				return err
			}

			createZoneRes, err := dnsClient.CreateZoneWithResponse(ctx, dns.CreateZoneJSONRequestBody{
				Domain:    config.DomainName(),
				MailAddr:  fmt.Sprintf("root.%v", config.DomainName()),
				PrimaryNs: fmt.Sprintf("ns1.%v", config.DomainName()),
			})
			if err != nil || createZoneRes.JSON201 == nil {
				if err == nil {
					err = errors.New("domain : unable to create zone")
				}
				return err
			}

			createRecordRes, err := dnsClient.CreateRecordWithResponse(ctx, config.DomainName(), dns.CreateRecordJSONRequestBody{
				Name:  "@",
				Type:  "A",
				Value: u.config.PublicIP(),
			})
			if err != nil || createRecordRes.JSON201 == nil {
				if err == nil {
					err = errors.New("domain : unable to create record")
				}
				return err
			}

			createRecordRes, err = dnsClient.CreateRecordWithResponse(ctx, config.DomainName(), dns.CreateRecordJSONRequestBody{
				Name:  "www",
				Type:  "A",
				Value: u.config.PublicIP(),
			})
			if err != nil || createRecordRes.JSON201 == nil {
				if err == nil {
					err = errors.New("domain : unable to create record")
				}
				return err
			}

			// Add Certificate
			certmanClient, err := certman.NewCertmanClient(proxy.FullPath, certmanService.Name()+":5555")
			if err != nil {
				return err
			}

			createCertificateRes, err := certmanClient.CreateCertificateWithResponse(ctx, certman.CreateCertificateJSONRequestBody{
				Domain:     config.DomainName(),
				Email:      fmt.Sprintf("admin@%v", config.DomainName()),
				AltDomains: &[]string{fmt.Sprintf("www.%v", config.DomainName())},
			})
			if err != nil || createCertificateRes.JSON201 == nil {
				if err == nil {
					err = errors.New("certificate : unable to create ssl certificate")
				}
				return err
			}

			return nil
		})
		if err != nil {
			return domainUsecase.WrapErrorSystem(err)
		}
	}

	return nil
}

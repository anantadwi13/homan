package usecase

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/anantadwi13/homan/internal/homan/domain"
	model2 "github.com/anantadwi13/homan/internal/homan/domain/model"
	service2 "github.com/anantadwi13/homan/internal/homan/domain/service"
	"github.com/anantadwi13/homan/internal/homan/domain/usecase"
	certman2 "github.com/anantadwi13/homan/internal/homan/external/api/certman"
	dns2 "github.com/anantadwi13/homan/internal/homan/external/api/dns"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/backend"
	backend_switching_rule2 "github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/backend_switching_rule"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/configuration"
	http_request_rule2 "github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/http_request_rule"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/server"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/storage"
	transactions2 "github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/transactions"
	models2 "github.com/anantadwi13/homan/internal/homan/external/api/haproxy/models"
	"github.com/anantadwi13/homan/internal/util"
	"github.com/go-openapi/runtime"
	"path/filepath"
	"regexp"
)

type ucAdd struct {
	config      domain.Config
	registry    service2.Registry
	executor    service2.Executor
	proxy       service2.Proxy
	regexDomain *regexp.Regexp
}

func NewUcAdd(
	config domain.Config, registry service2.Registry, executor service2.Executor, proxy service2.Proxy,
) usecase.UcAdd {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		// todo handle error
		return nil
	}
	return &ucAdd{config, registry, executor, proxy, reg}
}

func (u *ucAdd) Execute(ctx context.Context, params *usecase.UcAddParams) usecase.Error {
	err := u.preExecute(ctx, params)
	if err != nil {
		return err
	}

	var config model2.ServiceConfig
	switch params.ServiceType {
	case usecase.ServiceTypeBlog:
		dbName, err := u.getDbName(params.Domain)
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}

		config = model2.NewServiceConfig(
			params.Name,
			params.Domain,
			"wordpress:5.8.0-apache",
			[]string{
				"WORDPRESS_DB_HOST=system-mysql",
				"WORDPRESS_DB_USER=root",
				"WORDPRESS_DB_PASSWORD=my-secret-pw",
				"WORDPRESS_DB_NAME=" + dbName,
				"WORDPRESS_TABLE_PREFIX=wp_",
			},
			[]model2.Port{
				model2.NewPort(80),
			},
			[]model2.Volume{
				model2.NewVolumeBinding(filepath.Join(u.config.DataPath(), params.Name+"/wp-content/"), "/var/www/html/wp-content/"),
			},
			[]model2.HealthCheck{
				model2.NewHealthCheckHTTP(80, "/"),
			},
			[]string{u.config.ProjectName()},
			model2.TagWeb,
		)
	case usecase.ServiceTypeCustom:
	}

	if config == nil {
		return usecase.NewErrorSystem("unable to create service")
	}

	errAdd := u.registry.Add(ctx, config)
	if errAdd != nil {
		return usecase.WrapErrorSystem(errAdd)
	}

	err = u.postExecute(ctx, params, config)
	if err != nil {
		return err
	}

	return nil
}

func (u *ucAdd) preExecute(ctx context.Context, params *usecase.UcAddParams) usecase.Error {
	if params == nil || params.Name == "" || params.Domain == "" {
		return usecase.ErrorUcAddParamsNotFound
	}

	switch params.ServiceType {
	case usecase.ServiceTypeBlog:
	case usecase.ServiceTypeCustom:
	default:
		return usecase.NewErrorUser("unkown service type")
	}

	systemServices, err := u.registry.GetSystemServices(ctx)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}

	for _, systemService := range systemServices {
		isRunning, err := u.executor.IsRunning(ctx, systemService)
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}
		if !isRunning {
			return usecase.ErrorUcAddSystemNotRunning
		}
	}

	userServices, err := u.registry.GetUserServices(ctx)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}

	for _, userService := range userServices {
		if userService.Name() == params.Name {
			return usecase.NewErrorUser("duplicate service name")
		}
		if userService.DomainName() == params.Domain {
			return usecase.NewErrorUser("duplicate domain name")
		}
	}

	return nil
}

func (u *ucAdd) postExecute(
	ctx context.Context, params *usecase.UcAddParams, config model2.ServiceConfig,
) usecase.Error {
	// Run service
	err := u.executor.Run(ctx, config)
	if err != nil {
		return nil
	}

	// Add Database
	// todo

	services, err := u.registry.GetSystemServiceByTag(ctx, model2.TagGateway)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}
	if len(services) != 1 {
		return usecase.ErrorUcAddPostExecution
	}
	haproxyService := services[0]

	services, err = u.registry.GetSystemServiceByTag(ctx, model2.TagDNS)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}
	if len(services) != 1 {
		return usecase.ErrorUcAddPostExecution
	}
	dnsService := services[0]

	services, err = u.registry.GetSystemServiceByTag(ctx, model2.TagCertMan)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}
	if len(services) != 1 {
		return usecase.ErrorUcAddPostExecution
	}
	certmanService := services[0]

	if config.DomainName() != "" {
		err = u.proxy.Execute(ctx, func(proxy *model2.ProxyDetail) error {

			// Add Haproxy Backend

			haproxyClient, auth := haproxy.NewHaproxyClient(proxy.Host, haproxyService.Name()+":5555")

			version, err := haproxyClient.Configuration.GetConfigurationVersion(configuration.NewGetConfigurationVersionParams(), auth)
			if err != nil {
				return err
			}

			transaction, err := haproxyClient.Transactions.StartTransaction(transactions2.NewStartTransactionParams().WithVersion(version.Payload), auth)
			if err != nil {
				return err
			}

			transactionId := &transaction.Payload.ID
			mainFrontendName := u.config.ProjectName()

			rules, err := haproxyClient.BackendSwitchingRule.GetBackendSwitchingRules(
				backend_switching_rule2.NewGetBackendSwitchingRulesParams().WithTransactionID(transactionId).WithFrontend(mainFrontendName),
				auth,
			)
			if err != nil {
				return err
			}
			requestRulesRes, err := haproxyClient.HTTPRequestRule.GetHTTPRequestRules(
				http_request_rule2.NewGetHTTPRequestRulesParams().WithTransactionID(transactionId).WithParentType("frontend").WithParentName(mainFrontendName),
				auth,
			)
			if err != nil {
				return err
			}

			backendRules := rules.Payload.Data
			requestRules := requestRulesRes.Payload.Data

			configBackend := &models2.Backend{
				Name:       config.Name(),
				Mode:       "http",
				Forwardfor: &models2.Forwardfor{Enabled: util.String("enabled")},
				Balance:    &models2.Balance{Algorithm: util.String("roundrobin")},
			}

			configServer := &models2.Server{
				Name:    "server1",
				Address: config.Name(),
				Port:    util.Int64(80),
				Check:   "enabled",
			}

			configBackendRule := &models2.BackendSwitchingRule{
				Index:    util.Int64(int64(len(backendRules))),
				Name:     configBackend.Name,
				Cond:     "if",
				CondTest: fmt.Sprintf("{ hdr(host) -i %v www.%v }", config.DomainName(), config.DomainName()),
			}

			configHttpRequestRule := &models2.HTTPRequestRule{
				Index:      util.Int64(int64(len(requestRules))),
				Type:       "redirect",
				RedirType:  "scheme",
				RedirValue: "https",
				RedirCode:  util.Int64(302),
				Cond:       "if",
				CondTest:   fmt.Sprintf("{ hdr(host) -i %v www.%v } !{ ssl_fc }", config.DomainName(), config.DomainName()),
			}

			configHttpRequestHeaderRule := &models2.HTTPRequestRule{
				Index:     util.Int64(int64(len(requestRules) + 1)),
				Type:      "set-header",
				HdrName:   "X-Forwarded-Proto",
				HdrFormat: "https",
				Cond:      "if",
				CondTest:  fmt.Sprintf("{ hdr(host) -i %v www.%v }", config.DomainName(), config.DomainName()),
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
				backend_switching_rule2.NewCreateBackendSwitchingRuleParams().WithTransactionID(transactionId).WithFrontend(mainFrontendName).WithData(configBackendRule),
				auth,
			)
			if err != nil {
				return err
			}

			_, _, err = haproxyClient.HTTPRequestRule.CreateHTTPRequestRule(
				http_request_rule2.NewCreateHTTPRequestRuleParams().WithTransactionID(transactionId).WithParentType("frontend").WithParentName(mainFrontendName).WithData(configHttpRequestRule),
				auth,
			)
			if err != nil {
				return err
			}

			_, _, err = haproxyClient.HTTPRequestRule.CreateHTTPRequestRule(
				http_request_rule2.NewCreateHTTPRequestRuleParams().WithTransactionID(transactionId).WithParentType("frontend").WithParentName(mainFrontendName).WithData(configHttpRequestHeaderRule),
				auth,
			)
			if err != nil {
				return err
			}

			_, _, err = haproxyClient.Transactions.CommitTransaction(transactions2.NewCommitTransactionParams().WithID(*transactionId).WithForceReload(util.Bool(true)), auth)
			if err != nil {
				return err
			}

			// Add Domain Name
			dnsClient, err := dns2.NewDnsClient(proxy.FullPath, dnsService.Name()+":5555")
			if err != nil {
				return err
			}

			createZoneRes, err := dnsClient.CreateZoneWithResponse(ctx, dns2.CreateZoneJSONRequestBody{
				Domain:    config.DomainName(),
				MailAddr:  fmt.Sprintf("root.%v.", config.DomainName()),
				PrimaryNs: fmt.Sprintf("ns1.%v.", config.DomainName()),
			})
			if err != nil || createZoneRes.JSON201 == nil {
				if err == nil {
					err = errors.New("domain : unable to create zone")
				}
				return err
			}

			createRecordRes, err := dnsClient.CreateRecordWithResponse(ctx, config.DomainName(), dns2.CreateRecordJSONRequestBody{
				Name:  "ns1",
				Type:  "A",
				Value: u.config.PublicIP(),
			})
			if err != nil || createRecordRes.JSON201 == nil {
				if err == nil {
					err = errors.New("domain : unable to create record")
				}
				return err
			}

			createRecordRes, err = dnsClient.CreateRecordWithResponse(ctx, config.DomainName(), dns2.CreateRecordJSONRequestBody{
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

			createRecordRes, err = dnsClient.CreateRecordWithResponse(ctx, config.DomainName(), dns2.CreateRecordJSONRequestBody{
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
			certmanClient, err := certman2.NewCertmanClient(proxy.FullPath, certmanService.Name()+":5555")
			if err != nil {
				return err
			}

			createCertificateRes, err := certmanClient.CreateCertificateWithResponse(ctx, certman2.CreateCertificateJSONRequestBody{
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

			getCertificateRes, err := certmanClient.GetCertificateByDomainWithResponse(ctx, config.DomainName())
			if err != nil || getCertificateRes.JSON200 == nil {
				if err == nil {
					err = errors.New("certificate : unable to get ssl certificate")
				}
				return err
			}

			sslRaw := fmt.Sprintf("%v\n%v", getCertificateRes.JSON200.PublicCert, getCertificateRes.JSON200.PrivateCert)

			ssl := bytes.NewReader([]byte(sslRaw))

			_, err = haproxyClient.Storage.CreateStorageSSLCertificate(
				storage.NewCreateStorageSSLCertificateParams().WithFileUpload(runtime.NamedReader(
					config.DomainName(),
					ssl,
				)).WithForceReload(util.Bool(true)),
				auth,
			)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}
	}

	return nil
}

func (u *ucAdd) getDbName(domainName string) (string, error) {
	return u.regexDomain.ReplaceAllString(domainName, "_"), nil
}

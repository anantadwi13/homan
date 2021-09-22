package usecase

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/anantadwi13/homan/internal/homan/domain"
	domainModel "github.com/anantadwi13/homan/internal/homan/domain/model"
	domainService "github.com/anantadwi13/homan/internal/homan/domain/service"
	"github.com/anantadwi13/homan/internal/homan/domain/usecase"
	"github.com/anantadwi13/homan/internal/homan/external/api/certman"
	"github.com/anantadwi13/homan/internal/homan/external/api/dns"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/backend"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/backend_switching_rule"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/configuration"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/http_request_rule"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/server"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/storage"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/transactions"
	haproxyModel "github.com/anantadwi13/homan/internal/homan/external/api/haproxy/models"
	"github.com/anantadwi13/homan/internal/homan/external/service"
	"github.com/anantadwi13/homan/internal/util"
	"github.com/go-openapi/runtime"
	"path/filepath"
	"regexp"
)

type ucAdd struct {
	config      domain.Config
	registry    domainService.Registry
	executor    domainService.Executor
	proxy       domainService.Proxy
	cmd         service.Commander
	regexDomain *regexp.Regexp
}

func NewUcAdd(
	config domain.Config, registry domainService.Registry, executor domainService.Executor, proxy domainService.Proxy,
	commander service.Commander,
) usecase.UcAdd {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		// todo handle error
		return nil
	}
	return &ucAdd{
		config:      config,
		registry:    registry,
		executor:    executor,
		proxy:       proxy,
		cmd:         commander,
		regexDomain: reg,
	}
}

func (u *ucAdd) Execute(ctx context.Context, params *usecase.UcAddParams) usecase.Error {
	err := u.preExecute(ctx, params)
	if err != nil {
		return err
	}

	var config domainModel.ServiceConfig
	switch params.ServiceType {
	case usecase.ServiceTypeBlog:
		dbName := u.getDbName(params.Domain)
		if dbName == "" {
			return usecase.NewErrorSystem("cannot assign database name")
		}

		config = domainModel.NewServiceConfig(
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
			[]domainModel.Port{
				domainModel.NewPort(80),
			},
			[]domainModel.Volume{
				domainModel.NewVolumeBinding(filepath.Join(u.config.DataPath(), params.Name+"/"), "/var/www/html/"),
			},
			[]domainModel.HealthCheck{
				domainModel.NewHealthCheckTCP(80),
			},
			[]string{u.config.ProjectName()},
			domainModel.TagWeb,
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
		isRunning, err := u.executor.IsRunning(ctx, systemService, true)
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
	ctx context.Context, params *usecase.UcAddParams, config domainModel.ServiceConfig,
) usecase.Error {
	// Init service
	err := u.executor.Init(ctx, config)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}

	// Run service
	err = u.executor.RunWait(ctx, 60, config)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}

	// Add Database
	cmdRes, err := u.cmd.RunCommand(ctx, "docker", "run", "--rm", "--network", u.config.ProjectName(),
		"mysql:8", "mysql", "-u", "root", "-pmy-secret-pw", "-h", "system-mysql", "-e",
		fmt.Sprintf("create database if not exists %v", u.getDbName(config.DomainName())))
	if err != nil {
		return usecase.NewErrorSystem(string(cmdRes))
	}

	services, err := u.registry.GetSystemServiceByTag(ctx, domainModel.TagGateway)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}
	if len(services) != 1 {
		return usecase.ErrorUcAddPostExecution
	}
	haproxyService := services[0]

	services, err = u.registry.GetSystemServiceByTag(ctx, domainModel.TagDNS)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}
	if len(services) != 1 {
		return usecase.ErrorUcAddPostExecution
	}
	dnsService := services[0]

	services, err = u.registry.GetSystemServiceByTag(ctx, domainModel.TagCertMan)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}
	if len(services) != 1 {
		return usecase.ErrorUcAddPostExecution
	}
	certmanService := services[0]

	if config.DomainName() != "" {
		proxy, stop, err := u.proxy.Start(ctx, domainService.ProxyParams{Type: domainModel.ProxyHTTP})
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}
		defer stop()

		// Add Haproxy Backend

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

		rules, err := haproxyClient.BackendSwitchingRule.GetBackendSwitchingRules(
			backend_switching_rule.NewGetBackendSwitchingRulesParams().WithTransactionID(transactionId).WithFrontend(mainFrontendName),
			auth,
		)
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}
		requestRulesRes, err := haproxyClient.HTTPRequestRule.GetHTTPRequestRules(
			http_request_rule.NewGetHTTPRequestRulesParams().WithTransactionID(transactionId).WithParentType("frontend").WithParentName(mainFrontendName),
			auth,
		)
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}

		backendRules := rules.Payload.Data
		requestRules := requestRulesRes.Payload.Data

		configBackend := &haproxyModel.Backend{
			Name:       config.Name(),
			Mode:       "http",
			Forwardfor: &haproxyModel.Forwardfor{Enabled: util.String("enabled")},
			Balance:    &haproxyModel.Balance{Algorithm: util.String("roundrobin")},
		}

		configServer := &haproxyModel.Server{
			Name:    "server1",
			Address: config.Name(),
			Port:    util.Int64(80),
			Check:   "enabled",
		}

		configBackendRule := &haproxyModel.BackendSwitchingRule{
			Index:    util.Int64(int64(len(backendRules))),
			Name:     configBackend.Name,
			Cond:     "if",
			CondTest: fmt.Sprintf("{ hdr(host) -i %v www.%v }", config.DomainName(), config.DomainName()),
		}

		configHttpRequestRule := &haproxyModel.HTTPRequestRule{
			Index:      util.Int64(int64(len(requestRules))),
			Type:       "redirect",
			RedirType:  "scheme",
			RedirValue: "https",
			RedirCode:  util.Int64(302),
			Cond:       "if",
			CondTest:   fmt.Sprintf("{ hdr(host) -i %v www.%v } !{ ssl_fc }", config.DomainName(), config.DomainName()),
		}

		configHttpRequestHeaderRule := &haproxyModel.HTTPRequestRule{
			Index:     util.Int64(int64(len(requestRules) + 1)),
			Type:      "set-header",
			HdrName:   "X-Forwarded-Proto",
			HdrFormat: "https",
			Cond:      "if",
			CondTest:  fmt.Sprintf("{ hdr(host) -i %v www.%v }", config.DomainName(), config.DomainName()),
		}

		_, _, err = haproxyClient.Backend.CreateBackend(backend.NewCreateBackendParams().WithTransactionID(transactionId).WithData(configBackend), auth)
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}

		_, _, err = haproxyClient.Server.CreateServer(server.NewCreateServerParams().WithTransactionID(transactionId).WithBackend(configBackend.Name).WithData(configServer), auth)
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}

		_, _, err = haproxyClient.BackendSwitchingRule.CreateBackendSwitchingRule(
			backend_switching_rule.NewCreateBackendSwitchingRuleParams().WithTransactionID(transactionId).WithFrontend(mainFrontendName).WithData(configBackendRule),
			auth,
		)
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}

		_, _, err = haproxyClient.HTTPRequestRule.CreateHTTPRequestRule(
			http_request_rule.NewCreateHTTPRequestRuleParams().WithTransactionID(transactionId).WithParentType("frontend").WithParentName(mainFrontendName).WithData(configHttpRequestRule),
			auth,
		)
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}

		_, _, err = haproxyClient.HTTPRequestRule.CreateHTTPRequestRule(
			http_request_rule.NewCreateHTTPRequestRuleParams().WithTransactionID(transactionId).WithParentType("frontend").WithParentName(mainFrontendName).WithData(configHttpRequestHeaderRule),
			auth,
		)
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}

		_, _, err = haproxyClient.Transactions.CommitTransaction(transactions.NewCommitTransactionParams().WithID(*transactionId).WithForceReload(util.Bool(true)), auth)
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}

		// Add Domain Name
		dnsClient, err := dns.NewDnsClient(proxy.FullScheme, dnsService.Name()+":5555")
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}

		createZoneRes, err := dnsClient.CreateZoneWithResponse(ctx, dns.CreateZoneJSONRequestBody{
			Domain:    config.DomainName(),
			MailAddr:  fmt.Sprintf("root.%v.", config.DomainName()),
			PrimaryNs: fmt.Sprintf("ns1.%v.", config.DomainName()),
		})
		if err != nil || createZoneRes.JSON201 == nil {
			if err == nil {
				err = errors.New("domain : unable to create zone")
			}
			return usecase.WrapErrorSystem(err)
		}

		createRecordRes, err := dnsClient.CreateRecordWithResponse(ctx, config.DomainName(), dns.CreateRecordJSONRequestBody{
			Name:  "ns1",
			Type:  "A",
			Value: u.config.PublicIP(),
		})
		if err != nil || createRecordRes.JSON201 == nil {
			if err == nil {
				err = errors.New("domain : unable to create record")
			}
			return usecase.WrapErrorSystem(err)
		}

		createRecordRes, err = dnsClient.CreateRecordWithResponse(ctx, config.DomainName(), dns.CreateRecordJSONRequestBody{
			Name:  "@",
			Type:  "A",
			Value: u.config.PublicIP(),
		})
		if err != nil || createRecordRes.JSON201 == nil {
			if err == nil {
				err = errors.New("domain : unable to create record")
			}
			return usecase.WrapErrorSystem(err)
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
			return usecase.WrapErrorSystem(err)
		}

		// Add Certificate
		certmanClient, err := certman.NewCertmanClient(proxy.FullScheme, certmanService.Name()+":5555")
		if err != nil {
			return usecase.WrapErrorSystem(err)
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
			return usecase.WrapErrorSystem(err)
		}

		getCertificateRes, err := certmanClient.GetCertificateByDomainWithResponse(ctx, config.DomainName())
		if err != nil || getCertificateRes.JSON200 == nil {
			if err == nil {
				err = errors.New("certificate : unable to get ssl certificate")
			}
			return usecase.WrapErrorSystem(err)
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
			return usecase.WrapErrorSystem(err)
		}

		return nil
	}

	return nil
}

func (u *ucAdd) getDbName(domainName string) string {
	return u.regexDomain.ReplaceAllString(domainName, "_")
}

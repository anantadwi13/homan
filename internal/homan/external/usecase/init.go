package usecase

import (
	"context"
	"embed"
	"github.com/anantadwi13/homan/internal/homan/domain"
	model2 "github.com/anantadwi13/homan/internal/homan/domain/model"
	"github.com/anantadwi13/homan/internal/homan/domain/service"
	"github.com/anantadwi13/homan/internal/homan/domain/usecase"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/backend"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/backend_switching_rule"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/bind"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/configuration"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/frontend"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/server"
	transactions2 "github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/transactions"
	models2 "github.com/anantadwi13/homan/internal/homan/external/api/haproxy/models"
	"github.com/anantadwi13/homan/internal/util"
	"path/filepath"
)

type ucInit struct {
	registry  service.Registry
	executor  service.Executor
	config    domain.Config
	storage   service.Storage
	proxy     service.Proxy
	ucUp      usecase.UcUp
	templates embed.FS
}

func NewUcInit(
	registry service.Registry, executor service.Executor, config domain.Config,
	storage service.Storage, templates embed.FS, ucUp usecase.UcUp, proxy service.Proxy,
) usecase.UcInit {
	return &ucInit{
		registry:  registry,
		executor:  executor,
		config:    config,
		storage:   storage,
		templates: templates,
		ucUp:      ucUp,
		proxy:     proxy,
	}
}

func (u *ucInit) Execute(ctx context.Context, params *usecase.UcInitParams) usecase.Error {
	services, err := u.registry.GetSystemServices(ctx)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}
	if len(services) > 0 {
		return usecase.ErrorUcInitAlreadyInitialized
	}
	systemServices := u.systemServices()
	for _, serviceConfig := range systemServices {
		if !serviceConfig.IsValid() {
			return usecase.ErrorUcInitServiceConfigInvalid
		}
	}
	for _, serviceConfig := range systemServices {
		err := u.registry.AddSystem(ctx, serviceConfig)
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}
	}

	errUc := u.postExecute(ctx, params, systemServices)
	if errUc != nil {
		return errUc
	}
	return nil
}

func (u *ucInit) postExecute(
	ctx context.Context, params *usecase.UcInitParams, services map[string]model2.ServiceConfig,
) usecase.Error {
	// Copy Data
	for _, config := range services {
		if config.Name() == u.systemName("mysql") {
			for _, volume := range config.VolumeBindings() {
				err := u.storage.Mkdir(volume.HostPath())
				if err != nil {
					return usecase.WrapErrorSystem(err)
				}
			}
			continue
		}
		err := u.executor.InitVolume(ctx, config)
		if err != nil {
			return usecase.WrapErrorSystem(err)
		}
	}

	// Edit HAProxy Config
	haproxyCfg, err := u.templates.ReadFile("template/haproxy/haproxy.cfg")
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}
	err = u.storage.WriteFile(u.filePathJoin("/haproxy/haproxy.cfg"), haproxyCfg)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}

	// Copy Example SSL
	haproxySSL, err := u.templates.ReadFile("template/haproxy/example.com")
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}
	err = u.storage.WriteFile(u.filePathJoin("/haproxy/ssl/example.com"), haproxySSL)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}

	// Start All Services
	err2 := u.ucUp.Execute(ctx, nil)
	if err2 != nil {
		return err2
	}

	err = u.proxy.Execute(ctx, func(proxy *model2.ProxyDetail) error {
		haproxyClient, auth := haproxy.NewHaproxyClient(proxy.Host, services["haproxy"].Name()+":5555")

		version, err := haproxyClient.Configuration.GetConfigurationVersion(configuration.NewGetConfigurationVersionParams(), auth)
		if err != nil {
			return err
		}

		transaction, err := haproxyClient.Transactions.StartTransaction(transactions2.NewStartTransactionParams().WithVersion(version.Payload), auth)
		if err != nil {
			return err
		}

		transactionId := &transaction.Payload.ID

		mainBackend := &models2.Backend{
			Name: u.config.ProjectName(),
			Mode: "http",
		}

		certmanBackend := &models2.Backend{
			Name:       services["certman"].Name(),
			Mode:       "http",
			Forwardfor: &models2.Forwardfor{Enabled: util.String("enabled")},
			Balance:    &models2.Balance{Algorithm: util.String("roundrobin")},
		}

		certmanServer := &models2.Server{
			Name:    "server1",
			Address: services["certman"].Name(),
			Port:    util.Int64(80),
			Check:   "enabled",
		}

		certmanBackendRule := &models2.BackendSwitchingRule{
			Index:    util.Int64(0),
			Name:     certmanBackend.Name,
			Cond:     "if",
			CondTest: "{ path_beg /.well-known }",
		}

		mainFrontend := &models2.Frontend{
			Mode:           "http",
			Name:           u.config.ProjectName(),
			DefaultBackend: mainBackend.Name,
		}

		mainBind := &models2.Bind{
			Address: "*",
			Name:    "http",
			Port:    util.Int64(80),
		}

		mainSecureBind := &models2.Bind{
			Address:        "*",
			Name:           "https",
			Port:           util.Int64(443),
			Ssl:            true,
			SslCertificate: "/etc/haproxy/ssl/",
		}

		// Create Main Backend

		_, _, err = haproxyClient.Backend.CreateBackend(backend.NewCreateBackendParams().WithTransactionID(transactionId).WithData(mainBackend), auth)
		if err != nil {
			return err
		}

		// Create Certman Backend

		_, _, err = haproxyClient.Backend.CreateBackend(backend.NewCreateBackendParams().WithTransactionID(transactionId).WithData(certmanBackend), auth)
		if err != nil {
			return err
		}

		_, _, err = haproxyClient.Server.CreateServer(server.NewCreateServerParams().WithTransactionID(transactionId).WithBackend(certmanBackend.Name).WithData(certmanServer), auth)
		if err != nil {
			return err
		}

		// Create Main Frontend

		_, _, err = haproxyClient.Frontend.CreateFrontend(frontend.NewCreateFrontendParams().WithData(mainFrontend).WithTransactionID(transactionId), auth)
		if err != nil {
			return err
		}

		_, _, err = haproxyClient.Bind.CreateBind(bind.NewCreateBindParams().WithTransactionID(transactionId).WithData(mainBind).WithFrontend(mainFrontend.Name), auth)
		if err != nil {
			return err
		}

		_, _, err = haproxyClient.Bind.CreateBind(bind.NewCreateBindParams().WithTransactionID(transactionId).WithData(mainSecureBind).WithFrontend(mainFrontend.Name), auth)
		if err != nil {
			return err
		}

		// Create Certman Backend Rule

		_, _, err = haproxyClient.BackendSwitchingRule.CreateBackendSwitchingRule(
			backend_switching_rule.NewCreateBackendSwitchingRuleParams().WithTransactionID(transactionId).WithFrontend(mainFrontend.Name).WithData(certmanBackendRule),
			auth,
		)
		if err != nil {
			return err
		}

		_, _, err = haproxyClient.Transactions.CommitTransaction(transactions2.NewCommitTransactionParams().WithID(*transactionId), auth)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}
	return nil
}

func (u *ucInit) systemServices() map[string]model2.ServiceConfig {
	return map[string]model2.ServiceConfig{
		"haproxy": model2.NewServiceConfig(
			u.systemName("haproxy"),
			"",
			"haproxytech/haproxy-debian:2.4",
			[]string{},
			[]model2.Port{
				model2.NewPort(5555),
				model2.NewPortBinding(80, 80),
				model2.NewPortBinding(443, 443),
			},
			[]model2.Volume{
				model2.NewVolumeBinding(u.filePathJoin("/haproxy"), "/usr/local/etc/haproxy"),
			},
			[]model2.HealthCheck{model2.NewHealthCheckHTTP(5555, "/v2")},
			[]string{u.config.ProjectName()},
			model2.TagGateway,
		),
		"dns": model2.NewServiceConfig(
			u.systemName("dns"),
			"",
			"anantadwi13/dns-server-manager:0.3.0",
			[]string{},
			[]model2.Port{
				model2.NewPort(5555),
				model2.NewPortBindingTCP(53, 53),
				model2.NewPortBindingUDP(53, 53),
			},
			[]model2.Volume{
				model2.NewVolumeBinding(u.filePathJoin("/dns/data"), "/data"),
			},
			[]model2.HealthCheck{model2.NewHealthCheckTCP(5555)},
			[]string{u.config.ProjectName()},
			model2.TagDNS,
		),
		"certman": model2.NewServiceConfig(
			u.systemName("certman"),
			"",
			"anantadwi13/letsencrypt-manager:0.2.0",
			[]string{},
			[]model2.Port{
				model2.NewPort(5555),
				model2.NewPort(80),
			},
			[]model2.Volume{
				model2.NewVolumeBinding(u.filePathJoin("/certman/etc/letsencrypt"), "/etc/letsencrypt"),
			},
			[]model2.HealthCheck{model2.NewHealthCheckTCP(5555)},
			[]string{u.config.ProjectName()},
			model2.TagCertMan,
		),
		"mysql": model2.NewServiceConfig(
			u.systemName("mysql"),
			"",
			"mysql:8",
			[]string{
				"MYSQL_ROOT_PASSWORD=my-secret-pw",
			},
			[]model2.Port{
				model2.NewPort(3306),
			},
			[]model2.Volume{
				model2.NewVolumeBinding(u.filePathJoin("/mysql"), "/var/lib/mysql"),
			},
			[]model2.HealthCheck{model2.NewHealthCheckTCP(3306)},
			[]string{u.config.ProjectName()},
			model2.TagDB,
		),
	}
}

func (u *ucInit) systemName(name string) string {
	return u.config.SystemNamePrefix() + name
}

func (u *ucInit) filePathJoin(filePath string) string {
	path := filepath.Join(u.config.DataPath(), "/system")
	path = filepath.Join(path, filePath)
	return path
}

package usecase

import (
	"context"
	"embed"
	"github.com/anantadwi13/homan/internal/homan/domain"
	domainModel "github.com/anantadwi13/homan/internal/homan/domain/model"
	"github.com/anantadwi13/homan/internal/homan/domain/service"
	"github.com/anantadwi13/homan/internal/homan/domain/usecase"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/backend"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/backend_switching_rule"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/bind"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/configuration"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/frontend"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/server"
	"github.com/anantadwi13/homan/internal/homan/external/api/haproxy/client/transactions"
	haproxyModel "github.com/anantadwi13/homan/internal/homan/external/api/haproxy/models"
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
	ctx context.Context, params *usecase.UcInitParams, services map[string]domainModel.ServiceConfig,
) usecase.Error {
	// Run Core Daemon
	err := u.executor.RunWait(ctx, 10, u.registry.GetCoreDaemon(ctx))
	if err != nil && err != service.ErrorExecutorServiceIsRunning {
		return usecase.WrapErrorSystem(err)
	}

	// Copy Data
	for _, config := range services {
		switch config.Name() {
		case u.systemName("haproxy"):
			err := u.executor.InitVolume(ctx, config)
			if err != nil {
				return usecase.WrapErrorSystem(err)
			}
		default:
			for _, volume := range config.VolumeBindings() {
				err := u.storage.Mkdir(volume.HostPath())
				if err != nil {
					return usecase.WrapErrorSystem(err)
				}
			}
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

	proxy, stop, err := u.proxy.Start(ctx, service.ProxyParams{Type: domainModel.ProxyHTTP})
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}
	defer stop()

	haproxyClient, auth := haproxy.NewHaproxyClient(proxy.Host, services["haproxy"].Name()+":5555")

	version, err := haproxyClient.Configuration.GetConfigurationVersion(configuration.NewGetConfigurationVersionParams(), auth)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}

	transaction, err := haproxyClient.Transactions.StartTransaction(transactions.NewStartTransactionParams().WithVersion(version.Payload), auth)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}

	transactionId := &transaction.Payload.ID

	mainBackend := &haproxyModel.Backend{
		Name: u.config.ProjectName(),
		Mode: "http",
	}

	certmanBackend := &haproxyModel.Backend{
		Name:       services["certman"].Name(),
		Mode:       "http",
		Forwardfor: &haproxyModel.Forwardfor{Enabled: util.String("enabled")},
		Balance:    &haproxyModel.Balance{Algorithm: util.String("roundrobin")},
	}

	certmanServer := &haproxyModel.Server{
		Name:    "server1",
		Address: services["certman"].Name(),
		Port:    util.Int64(80),
		Check:   "enabled",
	}

	certmanBackendRule := &haproxyModel.BackendSwitchingRule{
		Index:    util.Int64(0),
		Name:     certmanBackend.Name,
		Cond:     "if",
		CondTest: "{ path_beg /.well-known }",
	}

	mainFrontend := &haproxyModel.Frontend{
		Mode:           "http",
		Name:           u.config.ProjectName(),
		DefaultBackend: mainBackend.Name,
	}

	mainBind := &haproxyModel.Bind{
		Address: "*",
		Name:    "http",
		Port:    util.Int64(80),
	}

	mainSecureBind := &haproxyModel.Bind{
		Address:        "*",
		Name:           "https",
		Port:           util.Int64(443),
		Ssl:            true,
		SslCertificate: "/etc/haproxy/ssl/",
	}

	// Create Main Backend

	_, _, err = haproxyClient.Backend.CreateBackend(backend.NewCreateBackendParams().WithTransactionID(transactionId).WithData(mainBackend), auth)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}

	// Create Certman Backend

	_, _, err = haproxyClient.Backend.CreateBackend(backend.NewCreateBackendParams().WithTransactionID(transactionId).WithData(certmanBackend), auth)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}

	_, _, err = haproxyClient.Server.CreateServer(server.NewCreateServerParams().WithTransactionID(transactionId).WithBackend(certmanBackend.Name).WithData(certmanServer), auth)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}

	// Create Main Frontend

	_, _, err = haproxyClient.Frontend.CreateFrontend(frontend.NewCreateFrontendParams().WithData(mainFrontend).WithTransactionID(transactionId), auth)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}

	_, _, err = haproxyClient.Bind.CreateBind(bind.NewCreateBindParams().WithTransactionID(transactionId).WithData(mainBind).WithFrontend(mainFrontend.Name), auth)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}

	_, _, err = haproxyClient.Bind.CreateBind(bind.NewCreateBindParams().WithTransactionID(transactionId).WithData(mainSecureBind).WithFrontend(mainFrontend.Name), auth)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}

	// Create Certman Backend Rule

	_, _, err = haproxyClient.BackendSwitchingRule.CreateBackendSwitchingRule(
		backend_switching_rule.NewCreateBackendSwitchingRuleParams().WithTransactionID(transactionId).WithFrontend(mainFrontend.Name).WithData(certmanBackendRule),
		auth,
	)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}

	_, _, err = haproxyClient.Transactions.CommitTransaction(transactions.NewCommitTransactionParams().WithID(*transactionId), auth)
	if err != nil {
		return usecase.WrapErrorSystem(err)
	}

	return nil
}

func (u *ucInit) systemServices() map[string]domainModel.ServiceConfig {
	return map[string]domainModel.ServiceConfig{
		"haproxy": domainModel.NewServiceConfig(
			u.systemName("haproxy"),
			"",
			"haproxytech/haproxy-debian:2.4",
			[]string{},
			[]domainModel.Port{
				domainModel.NewPort(5555),
				domainModel.NewPortBinding(80, 80),
				domainModel.NewPortBinding(443, 443),
			},
			[]domainModel.Volume{
				domainModel.NewVolumeBinding(u.filePathJoin("/haproxy"), "/usr/local/etc/haproxy"),
			},
			[]domainModel.HealthCheck{domainModel.NewHealthCheckTCP(5555)},
			[]string{u.config.ProjectName()},
			domainModel.TagGateway,
		),
		"dns": domainModel.NewServiceConfig(
			u.systemName("dns"),
			"",
			"anantadwi13/dns-server-manager:0.3.0",
			[]string{},
			[]domainModel.Port{
				domainModel.NewPort(5555),
				domainModel.NewPortBindingTCP(53, 53),
				domainModel.NewPortBindingUDP(53, 53),
			},
			[]domainModel.Volume{
				domainModel.NewVolumeBinding(u.filePathJoin("/dns/data"), "/data"),
			},
			[]domainModel.HealthCheck{domainModel.NewHealthCheckTCP(5555)},
			[]string{u.config.ProjectName()},
			domainModel.TagDNS,
		),
		"certman": domainModel.NewServiceConfig(
			u.systemName("certman"),
			"",
			"anantadwi13/letsencrypt-manager:0.2.0",
			[]string{},
			[]domainModel.Port{
				domainModel.NewPort(5555),
				domainModel.NewPort(80),
			},
			[]domainModel.Volume{
				domainModel.NewVolumeBinding(u.filePathJoin("/certman/etc/letsencrypt"), "/etc/letsencrypt"),
			},
			[]domainModel.HealthCheck{domainModel.NewHealthCheckTCP(5555)},
			[]string{u.config.ProjectName()},
			domainModel.TagCertMan,
		),
		"mysql": domainModel.NewServiceConfig(
			u.systemName("mysql"),
			"",
			"mysql:8",
			[]string{
				"MYSQL_ROOT_PASSWORD=my-secret-pw",
			},
			[]domainModel.Port{
				domainModel.NewPort(3306),
			},
			[]domainModel.Volume{
				domainModel.NewVolumeBinding(u.filePathJoin("/mysql"), "/var/lib/mysql"),
			},
			[]domainModel.HealthCheck{domainModel.NewHealthCheckTCP(3306)},
			[]string{u.config.ProjectName()},
			domainModel.TagDB,
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

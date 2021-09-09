package usecase

import (
	"context"
	"errors"
	"github.com/anantadwi13/cli-whm/internal/domain"
	"github.com/anantadwi13/cli-whm/internal/domain/model"
	"github.com/anantadwi13/cli-whm/internal/domain/service"
	"path/filepath"
)

type UcInitParams struct {
}

type UcInit interface {
	Execute(ctx context.Context, params *UcInitParams) error
}

var (
	ErrorUcInitAlreadyInitialized   = errors.New("error [init]: project is already initialized")
	ErrorUcInitServiceConfigInvalid = errors.New("error [init]: service config is invalid")
)

type ucInit struct {
	registry service.Registry
	config   domain.Config
	storage  service.Storage
}

func NewUcInit(registry service.Registry, config domain.Config, storage service.Storage) UcInit {
	return &ucInit{registry, config, storage}
}

func (u *ucInit) Execute(ctx context.Context, params *UcInitParams) error {
	services, err := u.registry.GetSystemServices(ctx)
	if err != nil {
		return err
	}
	if len(services) > 0 {
		return ErrorUcInitAlreadyInitialized
	}
	services = u.systemServices()
	for _, serviceConfig := range services {
		if !serviceConfig.IsValid() {
			return ErrorUcInitServiceConfigInvalid
		}
	}
	for _, serviceConfig := range services {
		err := u.registry.AddSystem(ctx, serviceConfig)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *ucInit) systemServices() []model.ServiceConfig {
	return []model.ServiceConfig{
		model.NewServiceConfig(
			"haproxy",
			"",
			"haproxytech/haproxy-debian:2.4",
			[]string{},
			[]model.Port{
				model.NewPort(5555),
				model.NewPortBinding(80, 80),
				model.NewPortBinding(443, 443),
			},
			[]model.Volume{
				model.NewVolumeBinding(u.filePathJoin("/haproxy"), "/etc/haproxy"),
			},
			[]string{u.config.ProjectName()},
			model.TypeProxy,
		),
		model.NewServiceConfig(
			"dns",
			"",
			"anantadwi13/dns-server-manager:0.2.0",
			[]string{},
			[]model.Port{
				model.NewPort(5555),
				model.NewPortBindingTCP(53, 53),
				model.NewPortBindingUDP(53, 53),
			},
			[]model.Volume{
				model.NewVolumeBinding(u.filePathJoin("/dns/data"), "/data"),
			},
			[]string{u.config.ProjectName()},
			model.TypeDNS,
		),
		model.NewServiceConfig(
			"certman",
			"",
			"anantadwi13/letsencrypt-manager:0.1.1",
			[]string{},
			[]model.Port{
				model.NewPort(5555),
				model.NewPort(80),
			},
			[]model.Volume{
				model.NewVolumeBinding(u.filePathJoin("/certman/etc/letsencrypt"), "/etc/letsencrypt"),
			},
			[]string{u.config.ProjectName()},
			model.TypeCertMan,
		),
	}
}

func (u *ucInit) filePathJoin(filePath string) string {
	path := filepath.Join(u.config.DataPath(), "/system")
	path = filepath.Join(path, filePath)
	return path
}

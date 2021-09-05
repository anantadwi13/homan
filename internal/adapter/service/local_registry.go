package service

import (
	"context"
	"encoding/json"
	"github.com/anantadwi13/cli-whm/internal/adapter/service/dto"
	"github.com/anantadwi13/cli-whm/internal/domain/model"
	domainService "github.com/anantadwi13/cli-whm/internal/domain/service"
)

type registry struct {
	config  domainService.Config
	storage domainService.Storage
}

func NewRegistry(
	config domainService.Config, storage domainService.Storage,
) domainService.Registry {
	return &registry{
		config:  config,
		storage: storage,
	}
}

func (i *registry) GetAll(ctx context.Context) ([]model.ServiceConfig, error) {
	system, user, err := i.readFromFile(ctx)
	if err != nil {
		return nil, err
	}
	var services []model.ServiceConfig
	for _, serviceConfig := range system {
		services = append(services, serviceConfig)
	}
	for _, serviceConfig := range user {
		services = append(services, serviceConfig)
	}
	return services, nil
}

func (i *registry) GetSystemServices(ctx context.Context) ([]model.ServiceConfig, error) {
	system, _, err := i.readFromFile(ctx)
	if err != nil {
		return nil, err
	}
	var services []model.ServiceConfig
	for _, serviceConfig := range system {
		services = append(services, serviceConfig)
	}
	return services, nil
}

func (i *registry) GetUserServices(ctx context.Context) ([]model.ServiceConfig, error) {
	_, user, err := i.readFromFile(ctx)
	if err != nil {
		return nil, err
	}
	var services []model.ServiceConfig
	for _, serviceConfig := range user {
		services = append(services, serviceConfig)
	}
	return services, nil
}

func (i *registry) Add(ctx context.Context, sc model.ServiceConfig) error {
	if sc == nil || !sc.IsValid() {
		return domainService.RegistryErrorServiceConfigInvalid
	}
	system, user, err := i.readFromFile(ctx)
	if err != nil {
		return err
	}
	if s, ok := user[sc.Name()]; ok && s != nil {
		return domainService.RegistryErrorServiceConfigExist
	}
	user[sc.Name()] = sc
	err = i.writeToFile(ctx, system, user)
	if err != nil {
		return err
	}
	return nil
}

func (i *registry) Remove(ctx context.Context, sc model.ServiceConfig) error {
	if sc == nil || !sc.IsValid() {
		return domainService.RegistryErrorServiceConfigInvalid
	}
	system, user, err := i.readFromFile(ctx)
	if err != nil {
		return err
	}
	if s, ok := user[sc.Name()]; ok && s != nil {
		return domainService.RegistryErrorServiceConfigExist
	}
	user[sc.Name()] = sc
	err = i.writeToFile(ctx, system, user)
	if err != nil {
		return err
	}
	return nil
}

func (i *registry) IsSystem(ctx context.Context, sc model.ServiceConfig) (bool, error) {
	if sc == nil || !sc.IsValid() {
		return false, domainService.RegistryErrorServiceConfigInvalid
	}
	system, _, err := i.readFromFile(ctx)
	if err != nil {
		return false, err
	}
	if s, ok := system[sc.Name()]; ok && s != nil {
		return true, nil
	}
	return false, nil
}

func (i *registry) writeToFile(ctx context.Context, systemServices, userServices map[string]model.ServiceConfig) error {
	system := make(map[string]*dto.Service)
	user := make(map[string]*dto.Service)
	for name, service := range systemServices {
		s, err := dto.MapServiceConfigToExternal(service)
		if err != nil {
			return err
		}
		system[name] = s
	}
	for name, service := range userServices {
		s, err := dto.MapServiceConfigToExternal(service)
		if err != nil {
			return err
		}
		user[name] = s
	}
	data, err := json.Marshal(dto.RegistryData{
		SystemServices: system,
		UserServices:   user,
	})
	if err != nil {
		return err
	}
	err = i.storage.WriteFile(i.config.ServiceRegistryConfPath(), data)
	if err != nil {
		return err
	}
	return nil
}

func (i *registry) readFromFile(ctx context.Context) (
	systemServices, userServices map[string]model.ServiceConfig, err error,
) {
	data, err := i.storage.ReadFile(i.config.ServiceRegistryConfPath())
	if err != nil {
		return
	}
	rData := &dto.RegistryData{}
	err = json.Unmarshal(data, rData)
	if err != nil {
		return
	}
	systemServices = make(map[string]model.ServiceConfig)
	userServices = make(map[string]model.ServiceConfig)
	for name, service := range rData.SystemServices {
		serviceConfig, err := dto.MapExternalToServiceConfig(name, service)
		if err != nil {
			return nil, nil, err
		}
		systemServices[name] = serviceConfig
	}
	for name, service := range rData.UserServices {
		serviceConfig, err := dto.MapExternalToServiceConfig(name, service)
		if err != nil {
			return nil, nil, err
		}
		systemServices[name] = serviceConfig
	}
	return
}

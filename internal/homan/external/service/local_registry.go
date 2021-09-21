package service

import (
	"context"
	"encoding/json"
	"github.com/anantadwi13/homan/internal/homan/domain"
	"github.com/anantadwi13/homan/internal/homan/domain/model"
	"github.com/anantadwi13/homan/internal/homan/domain/service"
	dto2 "github.com/anantadwi13/homan/internal/homan/external/service/dto"
	"io/fs"
)

type localRegistry struct {
	config  domain.Config
	storage service.Storage
}

func NewLocalRegistry(
	config domain.Config, storage service.Storage,
) service.Registry {
	return &localRegistry{
		config:  config,
		storage: storage,
	}
}

func (i *localRegistry) GetAll(ctx context.Context) ([]model.ServiceConfig, error) {
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

func (i *localRegistry) GetSystemServices(ctx context.Context) ([]model.ServiceConfig, error) {
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

func (i *localRegistry) GetSystemServiceByTag(ctx context.Context, tag model.ServiceTag) (
	[]model.ServiceConfig, error,
) {
	system, _, err := i.readFromFile(ctx)
	if err != nil {
		return nil, err
	}
	var services []model.ServiceConfig
	for _, serviceConfig := range system {
		if serviceConfig.Tag() == tag {
			services = append(services, serviceConfig)
		}
	}
	return services, nil
}

func (i *localRegistry) GetUserServices(ctx context.Context) ([]model.ServiceConfig, error) {
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

func (i *localRegistry) Add(ctx context.Context, sc model.ServiceConfig) error {
	if sc == nil || !sc.IsValid() {
		return service.ErrorRegistryServiceConfigInvalid
	}
	system, user, err := i.readFromFile(ctx)
	if err != nil {
		return err
	}
	if s, ok := user[sc.Name()]; ok && s != nil {
		return service.ErrorRegistryServiceConfigExist
	}
	user[sc.Name()] = sc
	err = i.writeToFile(ctx, system, user)
	if err != nil {
		return err
	}
	return nil
}

func (i *localRegistry) AddSystem(ctx context.Context, sc model.ServiceConfig) error {
	if sc == nil || !sc.IsValid() {
		return service.ErrorRegistryServiceConfigInvalid
	}
	system, user, err := i.readFromFile(ctx)
	if err != nil {
		return err
	}
	if s, ok := system[sc.Name()]; ok && s != nil {
		return service.ErrorRegistryServiceConfigExist
	}
	system[sc.Name()] = sc
	err = i.writeToFile(ctx, system, user)
	if err != nil {
		return err
	}
	return nil
}

func (i *localRegistry) Remove(ctx context.Context, sc model.ServiceConfig) error {
	if sc == nil || !sc.IsValid() {
		return service.ErrorRegistryServiceConfigInvalid
	}
	system, user, err := i.readFromFile(ctx)
	if err != nil {
		return err
	}
	if s, ok := user[sc.Name()]; !ok || s == nil {
		return service.ErrorRegistryServiceConfigNotFound
	}
	delete(user, sc.Name())
	err = i.writeToFile(ctx, system, user)
	if err != nil {
		return err
	}
	return nil
}

func (i *localRegistry) RemoveSystem(ctx context.Context, sc model.ServiceConfig) error {
	if sc == nil || !sc.IsValid() {
		return service.ErrorRegistryServiceConfigInvalid
	}
	system, user, err := i.readFromFile(ctx)
	if err != nil {
		return err
	}
	if s, ok := system[sc.Name()]; !ok || s == nil {
		return service.ErrorRegistryServiceConfigNotFound
	}
	delete(system, sc.Name())
	err = i.writeToFile(ctx, system, user)
	if err != nil {
		return err
	}
	return nil
}

func (i *localRegistry) IsSystem(ctx context.Context, sc model.ServiceConfig) (bool, error) {
	if sc == nil || !sc.IsValid() {
		return false, service.ErrorRegistryServiceConfigInvalid
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

func (i *localRegistry) writeToFile(
	ctx context.Context, systemServices, userServices map[string]model.ServiceConfig,
) error {
	system := make(map[string]*dto2.Service)
	user := make(map[string]*dto2.Service)
	for name, service := range systemServices {
		s, err := dto2.MapServiceConfigToExternal(service)
		if err != nil {
			return err
		}
		system[name] = s
	}
	for name, service := range userServices {
		s, err := dto2.MapServiceConfigToExternal(service)
		if err != nil {
			return err
		}
		user[name] = s
	}
	data, err := json.Marshal(dto2.RegistryData{
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

func (i *localRegistry) readFromFile(ctx context.Context) (
	systemServices, userServices map[string]model.ServiceConfig, err error,
) {
	data, err := i.storage.ReadFile(i.config.ServiceRegistryConfPath())
	if err != nil {
		switch err.(type) {
		case *fs.PathError:
			systemServices = make(map[string]model.ServiceConfig)
			userServices = make(map[string]model.ServiceConfig)
			return systemServices, userServices, nil
		}
		return
	}
	rData := &dto2.RegistryData{}
	err = json.Unmarshal(data, rData)
	if err != nil {
		return
	}
	systemServices = make(map[string]model.ServiceConfig)
	userServices = make(map[string]model.ServiceConfig)
	for name, service := range rData.SystemServices {
		serviceConfig, err := dto2.MapExternalToServiceConfig(name, service)
		if err != nil {
			return nil, nil, err
		}
		systemServices[name] = serviceConfig
	}
	for name, service := range rData.UserServices {
		serviceConfig, err := dto2.MapExternalToServiceConfig(name, service)
		if err != nil {
			return nil, nil, err
		}
		userServices[name] = serviceConfig
	}
	return
}

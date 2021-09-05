package service

import (
	"context"
	"github.com/anantadwi13/cli-whm/internal/adapter/service/dto"
	"github.com/anantadwi13/cli-whm/internal/domain/model"
	domainService "github.com/anantadwi13/cli-whm/internal/domain/service"
	"gopkg.in/yaml.v3"
)

type scp struct {
	registry domainService.Registry
}

func NewServiceConfigParser(
	registry domainService.Registry,
) *scp {
	return &scp{registry: registry}
}

func (s *scp) Marshal(ctx context.Context, configs ...model.ServiceConfig) ([]byte, error) {
	compose := &dto.DockerCompose{
		Version: "3.9",
	}

	isSystem := false
	networkNames := make(map[string]interface{})
	for _, config := range configs {
		if config.IsCustom() {
			continue
		}

		if checkSystem, err := s.registry.IsSystem(ctx, config); checkSystem && err != nil {
			isSystem = true
		}
		for _, networkName := range config.Networks() {
			networkNames[networkName] = nil
		}

		service, err := dto.MapServiceConfigToExternal(config)
		if err != nil {
			return nil, err
		}
		compose.Services[config.Name()] = service
	}
	for networkName := range networkNames {
		compose.Networks[networkName] = &dto.Network{
			Name:     networkName,
			External: !isSystem,
		}
	}

	return yaml.Marshal(compose)
}

func (s *scp) Unmarshal(_ context.Context, data []byte) ([]model.ServiceConfig, error) {
	compose := &dto.DockerCompose{}
	err := yaml.Unmarshal(data, compose)
	if err != nil {
		return nil, err
	}
	var configs []model.ServiceConfig
	for name, svc := range compose.Services {
		if svc.Build != nil {
			continue
		}

		config, err := dto.MapExternalToServiceConfig(name, svc)
		if err != nil {
			return nil, err
		}
		configs = append(configs, config)
	}
	return configs, nil
}

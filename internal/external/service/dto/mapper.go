package dto

import (
	"github.com/anantadwi13/cli-whm/internal/domain/model"
	"strconv"
	"strings"
)

func MapServiceConfigToExternal(config model.ServiceConfig) (*Service, error) {
	var (
		ports   []string
		volumes []string
	)
	for _, port := range config.PortBindings() {
		ports = append(ports, port.String())
	}
	for _, volume := range config.VolumeBindings() {
		volumes = append(volumes, volume.String())
	}

	return &Service{
		FilePath:    config.FilePath(),
		Image:       config.Image(),
		DomainName:  config.DomainName(),
		Environment: config.Environments(),
		Ports:       ports,
		Networks:    config.Networks(),
		Volumes:     volumes,
		Type:        string(config.Type()),
	}, nil
}

func MapExternalToServiceConfig(name string, svc *Service) (model.ServiceConfig, error) {
	var (
		ports          []model.Port
		volumeBindings []model.Volume
	)

	for _, port := range svc.Ports {
		p := strings.SplitN(port, ":", 2)
		var pInt []int
		for _, v := range p {
			vInt, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			pInt = append(pInt, vInt)
		}
		switch len(pInt) {
		case 2:
			ports = append(ports, model.NewPortBinding(pInt[0], pInt[1]))
		case 1:
			ports = append(ports, model.NewPort(pInt[0]))
		}
	}
	for _, volume := range svc.Volumes {
		v := strings.SplitN(volume, ":", 2)
		if len(v) == 2 {
			volumeBindings = append(volumeBindings)
		}
	}

	if svc.FilePath != "" {
		config := model.NewCustomServiceConfig(name, svc.DomainName, svc.FilePath, ports)
		return config, nil
	}

	return model.NewServiceConfig(
		name,
		svc.DomainName,
		svc.Image,
		svc.Environment,
		ports,
		volumeBindings,
		svc.Networks,
		model.ServiceType(svc.Type),
	), nil
}